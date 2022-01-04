// Code generated by Kitex v0.1.3. DO NOT EDIT.
package externalstorage

import (
	"github.com/cloudwego/kitex/server"
	"github.com/pingcap/kvproto/proto/kitex_gen/brpb"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler brpb.ExternalStorage, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
