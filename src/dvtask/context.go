package dvtask

import (
	"errors"
	"github.com/Dobryvechir/dvserver/src/dvevaluation"
	"log"
	"strconv"
)

type DvExecutionContext struct {
	Tasks   []DvTask
	Context *dvevaluation.DvContext
}

func (context *DvExecutionContext) initContext(tasks []DvTask) error {
	context.Tasks = tasks
	context.Context = engine.GetNewContext()
	return nil
}

func (context *DvExecutionContext) executeTasks(phaseScope int, blocks map[string]*DvExtendedBlock) error {
	n := len(context.Tasks)
	for i := 0; i < BLOCK_FIX; i++ {
		if (phaseScope & (1 << uint(i))) != 0 {
			for j := 0; j < n; j++ {
				task := &context.Tasks[j]
				block, ok := blocks[task.Name]
				if !ok {
					return errors.New("Task " + task.Name + " is unknown")
				}
				err := context.Context.SetGeneralArguments(task.Params)
				if err != nil {
					return dvevaluation.EnrichErrorStr(err, "At phase "+strconv.Itoa(i)+" setting parameters for routines["+strconv.Itoa(j)+"] in task "+task.Name)
				}
				routines := block.Routines[i]
				err = context.Context.ExecuteRoutines(routines, task.SubTasks)
				if err != nil {
					return dvevaluation.EnrichErrorStr(err, "At phase "+strconv.Itoa(i)+" executing routines["+strconv.Itoa(j)+"] in task "+task.Name)
				}
			}
		}
	}
	if LogTask {
		LogTasks(context.Tasks)
		log.Print(phaseScope)
		context.Context.DumpContextMemory()
	}
	return nil
}
