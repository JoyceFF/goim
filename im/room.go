package im

import (
	"github.com/gin-gonic/gin"
	log "goim/Ilog"
	"goim/db"
	"strconv"
)

type Room struct {
	UId  string `json:"uid"`
	Name string `json:"name"`
	Ts   string `json:"ts"`
}

type MyRoom struct {
	UId    string `json:"uid" xorm:"uid"`
	Name   string `json:"name"`
	Id     int    `json:"id"`
	UCount int64  `json:"uCount"`
	Ts     string `json:"ts" xorm:"ts"`
}

func CreateRoom(c *gin.Context) {
	room := &Room{}
	c.ShouldBind(room)

	result, err := db.Mysql().Exec("insert into im_room(uid,name) values(?,?)", room.UId, room.Name)
	if err != nil {
		log.Error(err)
		c.JSON(500, NewError(1, err.Error()))
	}
	id, _ := result.LastInsertId()

	_, err = db.Mysql().Exec("insert into im_room_users(uid,rid) values(?,?)", room.UId, strconv.FormatInt(id, 10))

	if err != nil {
		c.JSON(500, NewError(1, err.Error()))
	}

	c.JSON(200, room)
}

func MyRooms(c *gin.Context) {
	uid := c.Param("uid")
	myRooms := GetMyRooms(uid)
	c.JSON(200, myRooms)
}

func GetMyRooms(uid string) []MyRoom {
	myRooms := make([]MyRoom,0)
	db.Mysql().SQL("select im_room.id, im_room.name,im_room.uid from im_room left join im_room_users on im_room.id = im_room_users.rid where im_room_users.uid=? group by im_room.id", uid).Find(&myRooms)

	for i, v := range myRooms {
		v.UCount = GetRoomUserCount(v.Id)
		myRooms[i] = v;
	}
	return myRooms
}

func GetRoomUserCount(rid int) int64 {
	var count int64
	db.Mysql().SQL("select count(1) as count from im_room_users where rid=?", rid).Get(&count)
	return count
}

func JoinRoom(c *gin.Context) {
	uid := c.Param("uid")
	rid := c.Param("rid")

	_, err := db.Mysql().Exec("insert into im_room_users(uid,rid) values(?,?)", uid, rid)

	if err != nil {
		c.JSON(500, NewError(1, err.Error()))
	}

	c.JSON(200, gin.H{"join": "success"})
}
