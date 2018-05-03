/***********************************************************************
DvClient
Copyright 2018 - 2019 by Volodymyr Dobryvechir (dobrivecher@yahoo.com vdobryvechir@gmail.com)
************************************************************************/

package dvtask

import (
	"github.com/Dobryvechir/dvserver/src/dvevaluation"
	"errors"
        "log" 
	"strings"
)

type DvTask struct {
	Name     string            `json:"name"`
	SubTasks []string          `json:"subtasks"`
	Params   map[string]string `json:"params"`
}

const (
	BLOCK_INITIAL = iota
	BLOCK_INFO    = iota
	BLOCK_EXECUTE = iota
	BLOCK_VERIFY  = iota
	BLOCK_FIX     = iota
)

const (
	PHASE_INFO    = (1 << BLOCK_INITIAL) + (1 << BLOCK_INFO)
	PHASE_VERIFY  = (1 << BLOCK_INITIAL) + (1 << BLOCK_INFO) + (1 << BLOCK_VERIFY)
	PHASE_EXECUTE = (1 << BLOCK_INITIAL) + (1 << BLOCK_INFO) + (1 << BLOCK_EXECUTE) + (1 << BLOCK_VERIFY)
	PHASE_FIX     = (1 << BLOCK_FIX) + PHASE_EXECUTE
)

type DvBlock struct {
	Name    string   `json:"name"`
	Initial []string `json:"initial"`
	Info    []string `json:"info"`
	Verify  []string `json:"verify"`
	Execute []string `json:"execute"`
}

type DvExtendedBlock struct {
	Routines [][]string
}

var LogTask bool = true
var engine dvevaluation.DvScript
var engineBlocks map[string]DvExtendedBlock

func ExecuteTasks(tasks []DvTask, phase string) error {
	var phaseScope int
	switch strings.TrimSpace(strings.ToLower(phase)) {
	case "info":
		phaseScope = PHASE_INFO
	case "verify":
		phaseScope = PHASE_VERIFY
	case "execute":
		phaseScope = PHASE_EXECUTE
	case "fix":
		phaseScope = PHASE_FIX
	default:
		return errors.New("Only the following phases are supported: info, verify, execute, fix, but not \"" + phase + "\"")
	}
	if LogTask {
		LogTasks(tasks)
                log.Print(phaseScope)
	}
	context := DvExecutionContext{}
	if err := context.initContext(tasks); err != nil {
		return err
	}
	if err := context.executeTasks(); err != nil {
		return err
	}
	return nil
}

func InitTasks(scripts []string, routines []dvevaluation.DvRoutine, blocks []DvBlock) error {
	var err error
	engine, err = dvevaluation.ParseScripts(scripts)
	if err != nil {
		return err
	}
	err = engine.AddRoutines(routines)
	if err != nil {
		return err
	}
	engineBlocks = make(map[string]DvExtendedBlock)
	if len(blocks) > 0 {
		for _, block := range blocks {
			if _, ok := engineBlocks[block.Name]; ok {
				return errors.New("Block name " + block.Name + " is duplicated")
			}
			extendedBlock := DvExtendedBlock{Routines: make([][]string, BLOCK_FIX)}
			if err = extendedBlock.addPhase(BLOCK_INITIAL, block.Initial); err != nil {
				return err
			}
			if err = extendedBlock.addPhase(BLOCK_INFO, block.Info); err != nil {
				return err
			}
			if err = extendedBlock.addPhase(BLOCK_EXECUTE, block.Execute); err != nil {
				return err
			}
			if err = extendedBlock.addPhase(BLOCK_VERIFY, block.Verify); err != nil {
				return err
			}
		}
	}
	return nil
}

func (block DvExtendedBlock) addPhase(index int, routines []string) error {
	block.Routines[index] = routines
	err := engine.VerifyRoutines(routines)
	return err
}
