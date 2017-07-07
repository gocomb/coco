package rest

import (
	"github.com/emicklei/go-restful"
	"github.com/koko990/coco/api"
)


func InitApiserver(contaner *restful.Container){
	ws:=new(restful.WebService)
	RegisterTaskAPI(ws)
	contaner.Add(ws)
}


func RegisterTaskAPI(ws *restful.WebService){
	ws.Path("/task").
		Doc("for given users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	ws.Route(ws.POST("/").
		Doc("create tasks for given user").
		To(CreateTask).Reads(api.Task{}).
		Writes(api.TaskResponse{}))
}




