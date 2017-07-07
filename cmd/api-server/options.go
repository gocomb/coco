package main

import (
	"github.com/spf13/pflag"
)

type ApiServerOptions struct {
	HostIP string `json:"host_ip,omitempty"`
	Port   string `json:"port,omitempty"`
}

type ApiServer struct {
	ApiServerOptions *ApiServerOptions
}


func NewApiServer(opts *ApiServerOptions) *ApiServer {
	return &ApiServer{
		ApiServerOptions:opts,
	}
}

func NewApiOptions() *ApiServerOptions{
	return &ApiServerOptions{}
}


func (a *ApiServerOptions) AddFlags(fs *pflag.FlagSet){
	fs.StringVar(&a.HostIP,"host","127.0.0.1","set api-server host-ip")
	fs.StringVar(&a.Port,"port","8001","set api-server port")
	return
}







