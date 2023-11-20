package session

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"sync"
)

//定义MemorySessionManager对象（字段 ：存放所有session的map，读写锁）
//构造函数，用于获取对象

type MemorySessionManager struct {
	//大map
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

// 构造函数
func NewMemorySessionManager() *MemorySessionManager {
	sm := &MemorySessionManager{
		sessionMap: make(map[string]Session, 1024),
	}
	return sm
}
func (sm *MemorySessionManager) Init(addr string, options ...string) (err error) {
	return
}
func (sm *MemorySessionManager) CreateSession() (session Session, err error) {
	sm.rwlock.Lock()
	defer sm.rwlock.Unlock()
	//uuid
	id := uuid.NewV4()
	sessionId := id.String() //转string
	//创建session
	session = NewMemorySession(sessionId)
	//加到大map
	sm.sessionMap[sessionId] = session
	return

}
func (sm *MemorySessionManager) Get(sessionId string) (session Session, err error) {
	sm.rwlock.Lock()
	defer sm.rwlock.Unlock()
	session, ok := sm.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists")
		return
	}
	return
}
