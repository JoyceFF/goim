package im

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"goim/db"
	"gopkg.in/mgo.v2/bson"
	log "goim/Ilog"
	"bytes"
	"github.com/gin-gonic/gin"
	"fmt"
)

type User struct {
	UId        string                 `json:"uid" xorm:"uid"`
	Attributes map[string]interface{} `json:"attributes"`
	Status     bool                   `json:"status"`
	Type       int                    `json:"type" xorm:"type"`
	Ts         string                 `json:"ts" xorm:"ts"`
}

type UserAttributes struct {
	UId        string                 `json:"uid"`
	Attributes map[string]interface{} `json:"attributes"`
}

type Friends struct {
	Uid string `json:"uid"`
	Fid string `json:"fid"`
	Ts  string `json:"ts"`
}

func NewUserAttributes(uid string, attributes map[string]interface{}) *UserAttributes {
	userAttributes := &UserAttributes{}
	userAttributes.UId = uid
	userAttributes.Attributes = attributes
	return userAttributes
}

func CreateUser(c *gin.Context) {
	user := &User{}
	c.ShouldBind(user)

	if user.UId == "" {
		c.JSON(500, NewError(1, "uid not cant empty"))
		return
	}

	session := db.Mysql().NewSession()
	defer session.Close()

	// 启动事务
	if err := session.Begin(); err != nil {
		log.Error(err)
		c.JSON(500, NewError(1, err.Error()))
		return
	}

	exist, err := session.SQL("select * from im_user where uid=?", user.UId).Exist()

	if err != nil {
		log.Error(err)
		c.JSON(500, NewError(1, err.Error()))
		return
	}

	if exist {
		c.JSON(500, NewError(1, "uid already existed"))
		return
	}

	_, err = session.Exec("insert into im_user(uid,type) values(?,?)", user.UId, user.Type)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, NewError(1, err.Error()))
	} else {
		merr := SetUserAttributes(NewUserAttributes(user.UId, user.Attributes))
		if merr != nil {
			session.Rollback()
			c.JSON(http.StatusInternalServerError, merr)
			return
		}
		session.Commit()
		c.JSON(http.StatusOK, user)
	}
}

func FindUser(c *gin.Context) {

	skip := DefaultValue(c.Query("skip"), "0")
	limit := DefaultValue(c.Query("limit"), "20")
	where := DefaultValue(c.Query("where"), "{}")

	attributes := make([]UserAttributes,0)
	db.Mongo().Database.C("im_user_attributes").Find(ParseBson(where)).Skip(ParseInt(skip)).Limit(ParseInt(limit)).All(&attributes)
	count, _ := db.Mongo().Database.C("im_user_attributes").Find(ParseBson(where)).Count()
	var in bytes.Buffer
	var len = len(attributes)

	if len == 0 {
		c.JSON(200, gin.H{"count": count, "results": []User{}})
		return
	}

	for i, v := range attributes {
		if i != (len - 1) {
			in.WriteString("'" + v.UId + "',")
		} else {
			in.WriteString("'" + v.UId + "'")
		}
	}

	users := make([]User,1)

	err := db.Mysql().SQL(fmt.Sprintf("select * from im_user where uid in (%s)", in.String())).Find(&users)

	for i, user := range users {
		for _, attributes := range attributes {
			if attributes.UId == user.UId {
				user.Attributes = attributes.Attributes
				break
			}
		}
		user.Status = getStatus(user.UId)
		users[i] = user
	}

	if err != nil {
		c.Error(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"count": count, "results": users})
	}
}

func GetUser(c *gin.Context) {
	uid := c.Param("uid")
	user := &User{}
	exists, err := db.Mysql().SQL("select * from im_user where uid=?", uid).Get(user)
	if err != nil {
		c.JSON(500, NewError(1, err.Error()))
		return
	} else {
		if !exists {
			c.JSON(500, NewError(1, "user non existent"))
			return
		}

		attributes, _ := GetUserAttributes(uid)
		user.Attributes = attributes[0].Attributes
		c.JSON(http.StatusOK, user)
	}
}

func Login(uid string) *User {
	user := &User{}
	exists, err := db.Mysql().SQL("select * from im_user where uid=?", uid).Get(user)
	if err != nil {
		log.Error(err)
		return nil
	} else {
		if !exists {
			log.Error("user non existent")
			return nil
		}

		attributes, _ := GetUserAttributes(uid)
		user.Attributes = attributes[0].Attributes
		return user
	}
}

func UpdateUserAttributes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	attr := &UserAttributes{}
	GetBody(r, attr)

	err := SetUserAttributes(attr)
	if err != nil {
		SendError(w, err)
	} else {
		Send(w, "{\"msg\":\"success\"}")
	}
}

func SetUserAttributes(attributes *UserAttributes) *Error {
	userAttributes, err := GetUserAttributes(attributes.UId)
	if len(userAttributes) > 0 {
		err = db.Mongo().Database.C("im_user_attributes").Update(bson.M{"uid": attributes.UId}, attributes)
	} else {
		err = db.Mongo().Database.C("im_user_attributes").Insert(attributes)
	}
	var merr *Error
	if err != nil {
		merr = NewError(1, err.Error())
	}
	return merr
}

func formatUserAttributes(users ...User) {
	var uids []string
	for _, v := range users {
		uids = append(uids, v.UId)
	}

	userAttributes, _ := GetUserAttributes(uids...)
	for i, user := range users {

		for _, attributes := range userAttributes {
			if attributes.UId == user.UId {
				users[i].Attributes = attributes.Attributes
			}
		}
	}
}

func GetUserAttributes(uid ...string) ([]UserAttributes, error) {
	attributes := &[]UserAttributes{}
	err := db.Mongo().Database.C("im_user_attributes").Find(bson.M{"uid": bson.M{"$in": uid}}).All(attributes)
	return *attributes, err
}

func AddFriends(c *gin.Context) {

	uid := c.Param("uid")
	fid := c.Param("fid")

	if uid == "" || fid == "" {
		c.JSON(http.StatusInternalServerError, NewError(1, "uid and fid not cant empty"))
		return
	}

	if uid == fid {
		c.JSON(http.StatusInternalServerError, NewError(1, "not can add yourself"))
		return
	}

	exists := checkFriends(uid, fid)

	if exists {
		c.JSON(http.StatusInternalServerError, NewError(1, "already is friend"))
		return
	}
	_, err := db.Mysql().Exec("insert into im_friends(uid,fid) values(?,?)", uid, fid)
	if err != nil {
		log.Error(err)
		c.JSON(500, NewError(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"add": "success"})
}

func checkFriends(uid string, fid string) bool {
	exists, _ := db.Mysql().SQL("select * from im_friends where uid=? and fid=?", uid, fid).Exist()
	return exists
}

func MyFriends(c *gin.Context) {
	uid := c.Param("uid")
	users := GetMyFriends(uid)
	c.JSON(http.StatusOK, users)
}

func GetMyFriends(uid string) []User {
	friends := make([]Friends,0)
	db.Mysql().SQL("select * from im_friends where uid=?", uid).Find(&friends)

	users := make([]User,0)
	for _, fd := range friends {
		user := User{}
		user.UId = fd.Fid
		user.Status = getStatus(fd.Fid)
		users = append(users, user)
	}
	formatUserAttributes(users...)
	return users
}

func GetMyFriendIds(uid string) []string {
	friends := make([]Friends,0)
	db.Mysql().SQL("select fid from im_friends where uid=?", uid).Find(&friends)

	fids := make([]string, 0, 10)

	for _, fd := range friends {
		fids = append(fids, fd.Fid)
	}

	return fids
}

func deleteFriends(c *gin.Context) {
	uid := c.Param("uid")
	fid := c.Param("fid")
	db.Mysql().Exec("delete from im_friends where uid=? and fid=?", uid, fid)
	c.JSON(200, gin.H{"delete": "success"})
}
