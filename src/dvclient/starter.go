/***********************************************************************
DvClient
Copyright 2018 - 2019 by Volodymyr Dobryvechir (dobrivecher@yahoo.com vdobryvechir@gmail.com)
************************************************************************/

package dvclient

import (
	"github.com/Dobryvechir/dvclient/src/dvtask"
	"github.com/Dobryvechir/dvserver/src/dvlog"
)

func ClientStart() {
	cf := ReadConfig()
	dvlog.StartingLogFile()
	dvtask.InitTasks()
	dvtask.ExecuteTasks(cf.Tasks)
}
