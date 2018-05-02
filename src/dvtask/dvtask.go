/***********************************************************************
DvClient
Copyright 2018 - 2019 by Volodymyr Dobryvechir (dobrivecher@yahoo.com vdobryvechir@gmail.com)
************************************************************************/

package dvtask

type DvTask struct {
	Name   string   `json:"name"`
	Params []string `json:"params"`
}

var LogTask bool = true

func ExecuteTasks(tasks []DvTask) {
	if LogTask {
		LogTasks(tasks)
	}
}

func InitTasks() {

}
