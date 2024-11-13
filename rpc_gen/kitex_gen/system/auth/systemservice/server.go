// Code generated by Kitex v0.9.1. DO NOT EDIT.
package systemservice

import (
	server "github.com/cloudwego/kitex/server"
	auth "rpc_gen/kitex_gen/system/auth"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler auth.SystemService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)
	options = append(options, server.WithCompatibleMiddlewareForUnary())

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler auth.SystemService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
