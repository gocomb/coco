###初步想法

	
- step1:根据命令行传入的参数值，启动restapi(nerverStop)，响应构建请求
    - 后端每次得到请求后，检查权限、资源是否足够等等，然后将构建对象存入etcd、mongoDB中，返回新建任务成功
- step2:启动"操作"协程，处理构建队列中的任务，处理队列中的构建任务（创建、删除、更新），该协程只会生成一个taskBuild的job，无需对job进行监控
    - 这样做的目的一方面是为了：
	    - 在有多个job时从容处理，以后可以做成多个子任务串行进行，满足用户多种触发策略
        - 当用户需要定时触发时，可以选择cron job
	    - 无需监控一个是为了快速处理构建请求，另一个原因是k8s目前job的状态细粒度不高，很多用户关心的状态很难获取到
    - 在taskBuild中我们会让它完成以下任务：
	     生成一个构建pod,在
	    - 第一个init container 中添加git clone 容器，将代码下载到pod共享volume中的代码目录下
	    - 第二个init container 中进行任务的构建，并将构建产物存放在构建目录下，并将用户的运行脚本放在该目录下
		- 最终为用户启动运行时的container，挂载构建volume，运行"运行时"脚本，启动应用
	- taskBuilder全程会观察构建pod的状态，当有状态变更时，调用api-server响应的rest-api，更新状态(etcd,mongoDB),需要添加认证，类似于serviceaccount
- step3:启动监控协程，监控taskBuilder的心跳信息，更新状态到etcd
- step4:启动状态同步协程，将etcd的变动更新到构建队列