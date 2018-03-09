package im

import (
	"github.com/googollee/go-socket.io"
	log "goim/Ilog"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-contrib/cors"
	"time"
	"encoding/json"
	"errors"
	"goim/config"
)

func Init() {
	log.InitConfig();
}

var socket *socketio.Server
var socketManage = NewSocketsManage()

func Start() {
	Init()

	var err error
	socket, err = socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	//跨域中间件
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "DELETE", "POST", "PUT", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	//路由rest api
	routeApi(router)
	//路由socket.io
	routeChat()

	router.GET("/socket.io/", socketHandler)

	port := config.GetConfig().Port

	server := &http.Server{
		Addr:           ":"+port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	httpErr := server.ListenAndServe()

	if httpErr != nil {
		log.Error(httpErr)
		panic(httpErr)
	}

	log.Infof("Serving at localhost:%s...",port)
}

func routeApi(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.POST("", CreateUser)
		user.GET("/:uid", GetUser)
		user.GET("", FindUser)
	}

	friends := router.Group("/friends")
	{
		friends.POST(":uid/add/:fid", AddFriends)
		friends.GET("/:uid", MyFriends)
		friends.DELETE(":uid/delete/:fid", deleteFriends)
	}

	room := router.Group("/room")
    {
    	room.POST("",CreateRoom)
    	room.GET("/myRooms/:uid",MyRooms)
    	room.POST("/:uid/join/:rid",JoinRoom)
	}
}

type Message struct {
	Uid     string `json:"uid"`
	To      string `json:"to"`
	ToType  int    `json:"toType"`
	Msg     string `json:"msg"`
	MsgType int    `json:"type"`
	Ts      string `json:"ts"`
}

type ToMessage struct {
	Fid     string `json:"fid"`
	FType   int    `json:"fType"`
	Msg     string `json:"msg"`
	MsgType int    `json:"type"`
	Ts      string `json:"ts"`
}

func routeChat() {

	socket.SetAllowRequest(func(request *http.Request) error {
		auth := request.URL.Query().Get("auth")
		if auth != "go_im" {
          return errors.New("connection denied")
		}
		return nil;
	})

	socket.On("connection", func(so socketio.Socket) {
		log.Println("on connection")

		so.On("login", func(uid string) {
			log.Println("login:" + uid)
			socketManage.Set(uid, so)
			user := Login(uid)
			result := make(map[string]interface{})
			if user == nil {
				result["status"] = 4
			} else {
				result["status"] = 1
				result["user"] = user
			}

			str, _ := json.Marshal(result)
			so.Emit("login", string(str))
			if user == nil {
				so.Disconnect()
			} else {
				online(uid)
			}
		})

		so.On("disconnection", func() {
			log.Println("on disconnect")
			uid := socketManage.GetUidBySid(so.Id())
			if uid != "" {
				offline(uid)
			}
			socketManage.RemoveBySid(so.Id())
		})

		so.On("chat message", func(msg string) {
			log.Info("emit:", msg)

			message := &Message{}
			json.Unmarshal([]byte(msg), message)
			fSocket := socketManage.Get(message.To)
			if fSocket != nil {
				toMessage := &ToMessage{}

				toMessage.Msg = message.Msg
				toMessage.MsgType = message.MsgType
				toMessage.Fid = message.Uid
				toMessage.FType = message.ToType

				byte, _ := json.Marshal(toMessage)

				fSocket.Emit("chat message", string(byte))
			} else {

			}
		})
	})

	socket.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})
}

func socketHandler(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")
	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	socket.ServeHTTP(c.Writer, c.Request)
}

func joinMy(uid string){
	so := socketManage.Get(uid)
	so.Join(uid)
}

func roomExists(room string,rooms []string) bool {
	for _, v := range rooms {
		if v == room {
			return true
		}
	}

	return false
}

func online(uid string) {
	fids := GetMyFriendIds(uid)
	for _, fid := range fids {
		so := socketManage.Get(fid)
		if so != nil {
			so.Emit("online", uid)
		}
	}
}

func offline(uid string) {
	fids := GetMyFriendIds(uid)
	for _, fid := range fids {
		so := socketManage.Get(fid)
		if so != nil {
			so.Emit("offline", uid)
		}
	}
}

func getStatus(uid string) bool {
	if socketManage.Get(uid) != nil {
		return true
	}
	return false
}
