/***********************************************************************
DvClient
Copyright 2018 - 2019 by Volodymyr Dobryvechir (dobrivecher@yahoo.com vdobryvechir@gmail.com)
************************************************************************/

package dvclient

import (
	"github.com/Dobryvechir/dvclient/src/dvtask"

	"github.com/Dobryvechir/dvserver/src/dvcom"
	"github.com/Dobryvechir/dvserver/src/dvconfig"
	"github.com/Dobryvechir/dvserver/src/dvevaluation"
	"github.com/Dobryvechir/dvserver/src/dvjson"
	"github.com/Dobryvechir/dvserver/src/dvlog"
	"github.com/Dobryvechir/dvserver/src/dvparser"

	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

type DvConfig struct {
	Namespace  string                   `json:"namespace"`
	Listen     string                   `json:"listen"`
	ServerPath string                   `json:"serverPath"`
	RootFolder string                   `json:"rootFolder"`
	LogLevel   string                   `json:"logLevel"`
	LogModules string                   `json:"logModules"`
	Tasks      []dvtask.DvTask          `json:"tasks"`
	Phase      string                   `json:"phase"`
	Scripts    []string                 `json:"scripts"`
	Routines   []dvevaluation.DvRoutine `json:"routines"`
	Blocks     []dvtask.DvBlock         `json:"blocks"`
}

const DV_CLIENT_CONFIG = "DvClient.conf"
const DV_CLIENT_PROPERTIES = "DvClient.properties"

var LogConfig bool = false

func ReadConfig() DvConfig {
	filename := dvconfig.FindAndReadConfigs(DV_CLIENT_CONFIG, DV_CLIENT_PROPERTIES)
	cf := DvConfig{}
	if filename == "" {
		cf.Namespace = dvlog.CurrentNamespace
		cf.Listen = ":80"
		cf.ServerPath = "."
	} else {
		data, err := dvparser.SmartReadTemplate(filename, 3, byte(' '))
		if err == nil {
			dvlog.CleanEOL(data)
			if saveConfig, okSave := dvparser.GlobalProperties[dvconfig.DV_CONFIG_DEBUG_WRITE]; okSave {
				err2 := ioutil.WriteFile(saveConfig, data, 0644)
				if err2 != nil {
					log.Print("Cannot write resulted config to " + saveConfig + ": " + err2.Error())
				}
			}
			err = json.Unmarshal(data, &cf)
		}
		if err != nil {
			err2 := ioutil.WriteFile(dvconfig.CurrentDir+"/debug_dvclient_conf.txt", data, 0644)
			if err2 != nil {
				log.Print("Cannot write ./debug_dvclient_conf.txt: " + err2.Error())
			}
			panic("\nError: (see debug_dvclient_conf.txt)Incorrect json in " + filename + ": " + err.Error())
		}
	}
	dvparser.DvParserLog = false
	dvlog.SetCurrentNamespace(cf.Namespace)
	dvconfig.ResetNamespaceFolder()
	if cf.RootFolder != "" {
		dvlog.CurrentRootFolder = cf.RootFolder
	}
	if cf.LogLevel != "" {
		dvlog.SetLogLevel(cf.LogLevel)
	}
	logModules := strings.TrimSpace(cf.LogModules)
	if logModules != "" {
		logMods := dvparser.ConvertToList(logModules)
		for _, logModule := range logMods {
			if logModule == "" {
				continue
			}
			switch logModule {
			case "config":
				LogConfig = true
				if dvlog.CurrentLogLevel >= dvlog.LOG_LEVEL_ERROR {
					dvparser.DvParserLog = true
				}
			case "crud":
				dvjson.LogCrud = true
			case "json":
				dvjson.LogJson = true
			case "server":
				dvcom.LogServer = true
			case "hosts":
				dvcom.LogHosts = true
			default:
				log.Print("Available modules for logging are only server, config, crud and json, not " + logModule)
			}
		}
	}
	if LogConfig && dvlog.CurrentLogLevel >= dvlog.LOG_LEVEL_INFO {
		dvparser.LogVariables("GLOBAL VARIABLES", dvparser.GlobalProperties)
	}
	return cf
}
