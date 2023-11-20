package session

import "fmt"

//中间件
//用户选择使用版本

var (
	sessionManager SessionManager
)

func Init(provider string, addr string, options ...string) (sessionManager SessionManager, err error) {
	switch provider {
	case "memory":
		sessionManager = NewMemorySessionManager()
	case "redis":
		sessionManager = NewRedisSessionManager()
	default:
		fmt.Errorf("不支持")
		return
	}
	err = sessionManager.Init(addr, options...)
	return sessionManager, err
}
