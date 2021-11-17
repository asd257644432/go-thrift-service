package main

import (
	"context"
	"github.com/go/go-thrift-service/kitex_gen/github/go/go_thrift_service"
)

// HelloServiceImpl implements the last service interface defined in the IDL.
type HelloServiceImpl struct{}

// Hello implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) Hello(ctx context.Context, request *go_thrift_service.ReqOfHello) (resp *go_thrift_service.RespOfHello, err error) {
	// TODO: Your code here...
	return
}
