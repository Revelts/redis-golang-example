package response

import "time"

type TaskDetailInfo struct {
	MainTask string        `json:"main_task"`
	SubTask  []subTaskInfo `json:"sub_task"`
}

type subTaskInfo struct {
	TaskTitle  string          `json:"task_title"`
	TaskDetail []subTaskDetail `json:"task_detail"`
}

type subTaskDetail struct {
	Started time.Time `json:"started"`
	End     time.Time `json:"end"`
	About   string    `json:"about"`
	Status  string    `json:"status"`
}

type ResRedis struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
