package im

import (
	log "goim/Ilog"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"fmt"
	"strconv"
	"gopkg.in/mgo.v2/bson"
)

type Error struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
}

func NewError(code int, msg string) *Error {
	err := &Error{}
	err.Code = code
	err.Msg = msg
	return err
}

func ParamsConvert(ps httprouter.Params, v interface{}) error {
	b, err := json.Marshal(v)
	if err!=nil {
		return err
	}
	json.Unmarshal(b, v)
	return nil
}

func GetBody(r *http.Request, v interface{}) error{
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err!=nil {
		return err
	}
	json.Unmarshal(body, v)
	return nil
}

func Send(w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(http.StatusOK)
	switch result.(type) {
	case string:
		_,err := fmt.Fprint(w,result)
		if err !=nil {
			log.Error(err.Error())
		}
	default:
		buff,_:=json.Marshal(result)
		w.Write(buff)
	}
}

func SendError(w http.ResponseWriter, err *Error){
	log.Error(err.Msg)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	buff,_:=json.Marshal(err)
    w.WriteHeader(500)
    w.Write(buff)
}

func DefaultValue(value string,defaultValue string)string{
	if value == ""{
		return defaultValue
	}
	return value
}

func ParseInt(value string)int{
	i,_:=strconv.Atoi(value)
	return i
}

func ParseBson(value string)*bson.M{
	b :=&bson.M{}
	json.Unmarshal([]byte(value),b)
	return b
}