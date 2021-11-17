// Code generated by Kitex v0.0.8. DO NOT EDIT.

package helloservice

import (
	"github.com/cloudwego/kitex/server"
	"github.com/go/go-thrift-service/kitex_gen/github/go/go_thrift_service"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler go_thrift_service.HelloService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
