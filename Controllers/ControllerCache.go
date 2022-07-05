package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"redis-api/Controllers/Dto/request"
	"redis-api/Controllers/Dto/response"
	"redis-api/Library"
	"redis-api/Repository"
	"strconv"
	"time"
)

func RedisSetter(w http.ResponseWriter, r *http.Request) {
	var redisData request.ReqRedis
	err := json.NewDecoder(r.Body).Decode(&redisData)
	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	repository := Repository.AllRepository.Public

	ttl := r.URL.Query().Get("ttl_minutes")
	intMinutes, err := strconv.Atoi(ttl)
	minutes := time.Duration(intMinutes)
	err = repository.RedisSetCache(redisData.Key, redisData.Value, time.Minute*minutes)

	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	Library.HttpResponseSuccess(w, r, http.StatusOK)
}

func RedisGetter(w http.ResponseWriter, r *http.Request) {
	var redisData response.ResRedis
	repository := Repository.AllRepository.Public
	keys := chi.URLParam(r, "key")
	redisData.Key = keys
	err := repository.RedisGetCache(keys, &redisData.Value)
	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	Library.HttpResponseSuccess(w, r, redisData)
}

func GetCompanyTask(w http.ResponseWriter, r *http.Request) {
	companyName := r.URL.Query().Get("companyName")
	staffName := chi.URLParam(r, "username")
	keys := fmt.Sprintf("staff::%s", staffName)
	var companyTask response.TaskDetailInfo
	companyTask.MainTask = companyName
	repository := Repository.AllRepository.Public
	err := repository.RedisGetHMCache(keys, companyName, &companyTask.SubTask)
	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	Library.HttpResponseSuccess(w, r, companyTask)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	field := chi.URLParam(r, "username")
	keys := fmt.Sprintf("staff::%s", field)
	repository := Repository.AllRepository.Public
	taskData, err := repository.RedisHGetAllCache(keys)
	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	var userTaskData []response.TaskDetailInfo
	for keys, values := range taskData {
		var userTaskDatum response.TaskDetailInfo
		userTaskDatum.MainTask = keys
		json.Unmarshal([]byte(values), &userTaskDatum.SubTask)
		userTaskData = append(userTaskData, userTaskDatum)
	}
	Library.HttpResponseSuccess(w, r, userTaskData)
}

func DeleteChoosenCompanyTaskFromUser(w http.ResponseWriter, r *http.Request) {
	staffName := chi.URLParam(r, "username")
	companyName := r.URL.Query().Get("companyName")
	keys := fmt.Sprintf("staff::%s", staffName)
	repository := Repository.AllRepository.Public
	err := repository.RedisHDelCache(keys, companyName)
	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	Library.HttpResponseSuccess(w, r, http.StatusOK)
}

func DeleteAllUserTask(w http.ResponseWriter, r *http.Request) {
	staffName := chi.URLParam(r, "username")
	keys := fmt.Sprintf("staff::%s", staffName)
	repository := Repository.AllRepository.Public
	err := repository.RedisDelCache(keys)
	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	Library.HttpResponseSuccess(w, r, http.StatusOK)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var reqStaffTask request.ReqStaffTask
	err := json.NewDecoder(r.Body).Decode(&reqStaffTask)
	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	staffName := r.URL.Query().Get("staffName")
	ttl := r.URL.Query().Get("ttl_minutes")
	intMinutes, err := strconv.Atoi(ttl)
	minutes := time.Duration(intMinutes)
	keys := fmt.Sprintf("staff::%s", staffName)

	mapTask := make(map[string]interface{})
	for _, val := range reqStaffTask.TaskInfo {
		dataBytes, _ := json.Marshal(val.SubTask)
		mapTask[val.MainTask] = string(dataBytes)
	}

	repository := Repository.AllRepository.Public
	err = repository.RedisSetHMCache(keys, mapTask, time.Minute*minutes)
	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	Library.HttpResponseSuccess(w, r, http.StatusOK)
}
