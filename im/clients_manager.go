package im

import (
	"sync"
	"github.com/googollee/go-socket.io"
)

type socketsManage struct {
	sockets map[string]socketio.Socket  //uid:socket
	mapping map[string]string //sid:uid
	locker  sync.RWMutex
}

func NewSocketsManage() *socketsManage {
	return &socketsManage{
		sockets: make(map[string]socketio.Socket),
		mapping: make(map[string]string),
	}
}

func (s *socketsManage) Get(uid string) socketio.Socket {
	ret, ok := s.sockets[uid]
	if !ok {
		return nil
	}
	return ret
}

func (s *socketsManage) GetUidBySid(sid string) string {
	ret, ok := s.mapping[sid]
	if !ok {
		return ""
	}
	return ret
}

func (s *socketsManage) GetSid(uid string) string {
	ret, ok := s.sockets[uid]
	if !ok {
		return ""
	}

	sid,ok := s.mapping[ret.Id()]
	if !ok {
		return ""
	}
	return sid
}

func (s *socketsManage) Set(uid string, socket socketio.Socket) {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.sockets[uid] = socket
	s.mapping[socket.Id()] = uid
}

func (s *socketsManage) RemoveByUid(uid string) {
	s.locker.Lock()
	defer s.locker.Unlock()

	delete(s.sockets, uid)
	delete(s.mapping, s.GetSid(uid))
}


func (s *socketsManage) RemoveBySid(sid string) {
	s.locker.Lock()
	defer s.locker.Unlock()

	delete(s.sockets, s.mapping[sid])
	delete(s.mapping, sid)
}

