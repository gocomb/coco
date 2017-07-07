package rest

import (
	"github.com/emicklei/go-restful"
	"github.com/koko990/coco/api"
)

func CreateTask(request *restful.Request,response *restful.Response){

	response.WriteEntity(&api.TaskResponse{
		TaskId:"",
	})
}