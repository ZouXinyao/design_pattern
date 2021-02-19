package proxy

import (
	"log"
	"time"
)

type ServerProxy struct {
	server *Server
}

func (p *ServerProxy) Request(Id int) error {
	start := time.Now()

	// 这其中也可以添加一些校验逻辑。

	// 原有的业务逻辑
	if err := p.server.Request(Id); err != nil {
		return err
	}

	// 监控统计等
	log.Printf("user login cost time: %s", time.Now().Sub(start))

	return nil
}

func NewServerProxy(s *Server) *ServerProxy {
	return &ServerProxy{server: s}
}