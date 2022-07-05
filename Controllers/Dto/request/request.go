package request

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
	Started string `json:"started"`
	End     string `json:"end"`
	About   string `json:"about"`
	Status  int    `json:"status"`
}
