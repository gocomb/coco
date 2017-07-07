package main

import (
	"github.com/spf13/pflag"
	"fmt"
	"os"
	"github.com/emicklei/go-restful"
	"runtime"
	"math/rand"
	"net/http"
	"github.com/koko990/coco/api/rest"
	"github.com/koko990/coco/util/logger"

	"time"
	"github.com/golang/glog"
)



func main(){

	ops:=NewApiOptions()
	ops.AddFlags(pflag.CommandLine)
	pflag.Parse()

	logger.InitLogs()
	s:=NewApiServer(ops)

	if err:=s.Run();err!=nil{
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}


}


func(s *ApiServer) Run() error{


	runtime.GOMAXPROCS(runtime.NumCPU())

	rand.Seed(time.Now().UTC().UnixNano())

	httpServer:=s.ApiServerOptions.NewHttpServer()

	//TODO 添加三个协程，controller、statusManager、monitor


	glog.V(0).Infoln("now start api server ......")
	if err:=httpServer.ListenAndServe();err!=nil{
		return err
	}
	return nil
}




func(s *ApiServerOptions) NewHttpServer() *http.Server{
	wsContainer:=restful.NewContainer()
	rest.InitApiserver(wsContainer)
	ListenningAddress:=s.HostIP+":"+s.Port
	server:=&http.Server{Addr:ListenningAddress,Handler:wsContainer}
	return server
}



func(s *ApiServerOptions) InitEtcd(){
//TODO 初始化etcd


}


func(s *ApiServerOptions) InitStorage(){
	//TODO 初始化storage


}
