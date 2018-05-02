/***********************************************************************
DvClient
Copyright 2018 - 2019 by Volodymyr Dobryvechir (dobrivecher@yahoo.com vdobryvechir@gmail.com)
************************************************************************/

package dvtask

import (
	"log"
	"strings"
)

func LogTasks(tasks []DvTask) {
	if len(tasks) == 0 {
		log.Print("No tasks")
	} else {
		for _, task := range tasks {
			p := task.Name + " [" + strings.Join(task.Params, ",") + "]"
			log.Print(p)
		}
	}
}
