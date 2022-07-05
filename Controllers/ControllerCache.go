package Controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"redis-api/Controllers/Dto/request"
	"redis-api/Library"
	"redis-api/Repository"
	"time"
)

//func GetStaffList(w http.ResponseWriter, r *http.Request) {
//	var staffList response.StaffList
//	repository := Repository.AllRepository.Public
//	err := repository.GetCache("staff_list", &staffInfo)
//	if err != nil {
//		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	Library.HttpResponseSuccess(w, r, dataUser)
//}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var reqStaffTask request.ReqStaffTask
	err := json.NewDecoder(r.Body).Decode(&reqStaffTask)
	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	staffName := r.URL.Query().Get("staffName")
	keys := fmt.Sprintf("staff::%s", staffName)

	mapTask := make(map[string]interface{})
	for _, val := range reqStaffTask.TaskInfo {
		dataBytes, _ := json.Marshal(val.SubTask)
		mapTask[val.MainTask] = string(dataBytes)
	}

	repository := Repository.AllRepository.Public
	err = repository.SetHMCache(keys, mapTask, time.Minute*10)
	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	Library.HttpResponseSuccess(w, r, http.StatusOK)
}
