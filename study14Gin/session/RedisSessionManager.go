package session

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

//定义RedisSessionManager对象（字段 ：存放所有session的map，读写锁）
//构造函数，用于获取对象

type RedisSessionManager struct {
	addr       string //redis 地址
	password   string //密码
	pool       *redis.Pool
	rwlock     sync.RWMutex //d读写锁
	sessionMap map[string]Session
}

// 构造函数
func NewRedisSessionManager() *RedisSessionManager {
	sm := &RedisSessionManager{
		sessionMap: make(map[string]Session, 32),
	}
	return sm
}
func (sm *RedisSessionManager) Init(addr string, options ...string) (err error) {
	//若有其它参数
	if len(options) > 0 {
		sm.password = options[0]
	}
	//创建连接池
	sm.pool = myPool(addr, sm.password)
	sm.addr = addr
	return
}
func myPool(addr, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: time.Second * 20,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			//若有密码，判断
			if _, err := conn.Do("AUTH", password); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, err
		},
		//连接测试，开发时用，上线注掉
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
func (sm *RedisSessionManager) CreateSession() (session Session, err error) {
	sm.rwlock.Lock()
	defer sm.rwlock.Unlock()
	//uuid
	id := uuid.NewV4()
	sessionId := id.String() //转string
	//创建session
	session = NewRedisSession(sessionId, sm.pool)
	//加到大map
	sm.sessionMap[sessionId] = session
	return
}
func (sm *RedisSessionManager) Get(sessionId string) (session Session, err error) {
	sm.rwlock.Lock()
	defer sm.rwlock.Unlock()
	session, ok := sm.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists")
		return
	}
	return
}
