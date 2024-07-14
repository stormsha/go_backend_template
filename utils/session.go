package utils

import (
	"bytes"
	uuid "github.com/satori/go.uuid"
	"runtime"
	"strconv"
	"sync"
)

type Session struct {
	gid  uint64
	uuid string
	user *User
}

var GlobalSessionMap = make(map[uint64]*Session)
var mapLock = sync.RWMutex{}

func (s *Session) GetSessionUUID() string {
	return s.uuid
}

func (s *Session) GetSessionGID() uint64 {
	return s.gid
}

func NewSession() (r *Session) {
	s := Session{}
	s.uuid = uuid.NewV4().String()
	s.gid = GetGoroutineID()
	return &s
}

func SetSessionUUid(uuid string) {
	s := GetCurrentSession()
	s.uuid = uuid
}

func SetSessionUser(user *User) {
	s := GetCurrentSession()
	s.user = user
}

func CacheSession() {
	s := NewSession()
	mapLock.Lock()
	GlobalSessionMap[s.gid] = s
	mapLock.Unlock()
}

func GetCurrentSession() *Session {
	gid := GetGoroutineID()
	mapLock.RLock()
	session := GlobalSessionMap[gid]
	mapLock.RUnlock()
	return session
}

func ClearCurrentSession() {
	gid := GetGoroutineID()
	mapLock.Lock()
	delete(GlobalSessionMap, gid)
	mapLock.Unlock()
}

func GetGoroutineID() uint64 {
	b := make([]byte, 64)
	runtime.Stack(b, false)
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func GetCurrentUUID() string {
	s := GetCurrentSession()
	if s == nil {
		return ""
	}
	return s.uuid
}

func GetCurrentUser() *User {
	s := GetCurrentSession()
	if s == nil || s.user == nil {
		return &User{}
	}
	return s.user
}

func GetCurrentUserName() string {
	s := GetCurrentSession()
	if s == nil || s.user == nil {
		return ""
	}
	return s.user.UserName
}
