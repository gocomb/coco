package api

type Task struct {
	TaskID      string     `json:"_id,omitempty"`
	TaskName    string     `json:"task_name,omitempty"`
	UserID      string     `json:"user_id,omitempty"`
	UserName    string     `json:"user_name,omitempty"`
	Description string     `json:"description,omitempty"`
	Repository  Repository `json:"repository,omitempty"`
	Hooks       []Hook     `json:"hook,omitempty"`
}

type Hook struct {
	CallBack string `json:"callback,omitempty"`
	Token    string `json:"token,omitempty"`
}

type Repository struct {
	Url      string `json:"url,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
	Webhook  string `json:"webhook,omitempty"`
}



type TaskResponse struct {
	TaskId string `json:"TaskId,omitempty"`
	ErrorMessage string `json:"ErrorMessage,omitempty"`
}