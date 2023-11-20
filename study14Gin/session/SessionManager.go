package session

// 定义管理者 ，管理所有session
type SessionManager interface {
	//初始化
	Init(addr string, options ...string) (err error)
	CreateSession() (session Session, err error)
	Get(sessionId string) (session Session, err error)
}
