package request

import "time"

type ReqStaffTask struct {
	TaskInfo []taskDetailInfo `json:"list_task"`
}

type taskDetailInfo struct {
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

type ReqRedis struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
