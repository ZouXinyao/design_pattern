package proxy

type IServer interface {
	Request(UserId int) error
}

type Server struct {
}

func (s *Server) Request(Id int) error {
	// 这中间有请求的业务代码
	return nil
}