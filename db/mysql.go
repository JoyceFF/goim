package db

import (
	log "goim/Ilog"
	_ "github.com/go-sql-driver/mysql"
	"goim/config"
	"fmt"
	"strings"
	"bytes"
	"github.com/go-xorm/xorm"
)

type myMysql *xorm.Engine


var mysql *xorm.Engine

func Mysql() *xorm.Engine {
	if mysql == nil {
		initMysql()
	}
	return mysql
}

func initMysql() {
	if mysql == nil {
		mysqlConfig := config.GetConfig().Mysql
		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)
		log.Infof("mysql:%s", dataSourceName)
		engine, err := xorm.NewEngine("mysql", dataSourceName)
		if err !=nil{
			log.Error(err)
			panic(err)
		}
		engine.SetMaxIdleConns(10)
		engine.SetMaxOpenConns(50)
		//engine.ShowSQL(true)
		mysql = engine
	}
}

//首字母大写  a_b_c aBC
func upperCase(str string) string {
	temp := strings.Split(str, "_")
	var upperStr bytes.Buffer
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		if y != 0 {
			for i := 0; i < len(vv); i++ {
				if i == 0 {
					vv[i] -= 32
					upperStr.WriteString(string(vv[i])) // + string(vv[i+1])
				} else {
					upperStr.WriteString(string(vv[i]))
				}
			}
		}
	}
	return temp[0] + upperStr.String()
}
