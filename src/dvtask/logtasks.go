/***********************************************************************
DvClient
Copyright 2018 - 2019 by Volodymyr Dobryvechir (dobrivecher@yahoo.com vdobryvechir@gmail.com)
************************************************************************/

package dvtask

import (
	"log"
)

func LogTasks(tasks []DvTask) {
	if len(tasks) == 0 {
		log.Print("No tasks")
	} else {
		for _, task := range tasks {
			p := task.Name + " ["
			comma := ""
			for k, v := range task.Params {
				p += comma + `"` + k + `": "` + v + `"`
				comma = ","
			}
			p += "] " 
			log.Print(p)
		}
	}
}
