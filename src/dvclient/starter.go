/***********************************************************************
DvClient
Copyright 2018 - 2019 by Volodymyr Dobryvechir (dobrivecher@yahoo.com vdobryvechir@gmail.com)
************************************************************************/

package dvclient

import (
        "log"
	"github.com/Dobryvechir/dvclient/src/dvtask"
	"github.com/Dobryvechir/dvserver/src/dvlog"
)

func ClientStart() {
	cf := ReadConfig()
	dvlog.StartingLogFile()
	if err:=dvtask.InitTasks(cf.Scripts, cf.Routines, cf.Blocks);err!=nil {
            log.Print(err.Error())
        }
	if err:=dvtask.ExecuteTasks(cf.Tasks, cf.Phase);err!=nil {
            log.Print(err.Error())
        }
}
