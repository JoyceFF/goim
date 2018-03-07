package db

import (
	"gopkg.in/mgo.v2"
	"goim/config"
	log "goim/Ilog"
	"fmt"
)

type myMongo struct {
	Session  *mgo.Session
	Database *mgo.Database
}

var mongo *myMongo

func Mongo() *myMongo {
	if mongo == nil {
		initMongo()
	}
	return mongo
}

func initMongo() {
	if mongo == nil {
		mongoConfig := config.GetConfig().Mongo
		url := fmt.Sprintf("%s:%s", mongoConfig.Host, mongoConfig.Port)
		log.Info("mongo:%s",url)
		session, err := mgo.Dial(url)
		if err != nil {
			log.Errorf("mongo Dial error:%s", err.Error())
			panic(err)
		}
		database := session.DB(mongoConfig.Database)
		mongo = &myMongo{}
		mongo.Session = session
		mongo.Database = database
	}
}
