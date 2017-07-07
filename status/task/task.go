package task

import "github.com/koko990/coco/api"

type taskStatus struct {
	taskUpdate <-chan taskUpdate
}

//0删除 1创建 2更新
type TaskOperarion int

type taskUpdate struct {
	Operation TaskOperarion
	task api.Task
}


//从etcd更新状态到channel，触发controller的构建
//每10秒钟去获取一次，或者当有etcd事件发生时更新channel
func (ts *taskStatus)SyncTask(){
}