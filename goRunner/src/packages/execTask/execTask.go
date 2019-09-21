package execTask

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"packages/vars"
)

var taskExecution []*exec.Cmd

func ExecTask() {	

	// KILL ALL PREVIOUS TASK EXECUTED AND CLEAR ARRAY
	for _, item := range taskExecution {

	  err := item.Process.Kill()

	  if err != nil {
	  	fmt.Println(err)
	  }
	}
	taskExecution = nil


	// CLEAR TERMINAL
	cmd := exec.Command("clear") //Unix example
	
	if (runtime.GOOS == "windows") {

	  cmd = exec.Command("cmd", "/c", "cls") //Windows example, its tested 
	}
	cmd.Stdout = os.Stdout
  cmd.Run()
  cmd.Process.Kill()

	// EXEC ALL TASKS LIST IN config.json
	for _, item := range vars.TaskList {



		var stderr bytes.Buffer
		var arrayTask []string = strings.Split(item, " ")
		var task string = arrayTask[0]
		var parameters []string = arrayTask[1:]

		cmd := exec.Command(task, parameters...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = &stderr

	  err := cmd.Run()

		if err != nil {

			fmt.Println(err.Error())
			fmt.Println("\n\n"+stderr.String())
		}

		// Need pointers to the task in order to kill at next save
		taskExecution = append(taskExecution, cmd)
	}

	fmt.Println("\n\n TASK END at "+time.Now().String()+"\n\n")	

	return
}