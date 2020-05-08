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

	/* 
	 *	KILL ALL PREVIOUS TASK EXECUTED AND CLEAR ARRAY
	 *
	 */
	for _, item := range taskExecution {

	  err := item.Process.Kill()

	  if err != nil {
	  	fmt.Println(err)
	  }
	}
	taskExecution = nil


	/* 
	 *	CLEAR TERMINAL
	 *
	 */
	
	//Unix example
	cmd := exec.Command("clear") 
	
	if (runtime.GOOS == "windows") {

	  //Windows example, its tested 
	  cmd = exec.Command("cmd", "/c", "cls") 
	}
	cmd.Stdout = os.Stdout
  	cmd.Run()
  	cmd.Process.Kill()

	/* 
	 *	EXEC ALL TASKS LIST IN config.json
	 *
	 */
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

		/* 
		 *	Insert pointer into taskExecution. We'll have to kill the processes
		 *
		 */
		taskExecution = append(taskExecution, cmd)
	}

	fmt.Println("\n\n TASK END at "+time.Now().String()+"\n\n")	

	return
}