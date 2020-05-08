package util

import (
	"strings"	
	"packages/vars"
	"packages/execTask"
)


func IsDirToCheck(path string) bool {

	for _, item := range vars.ExceptionDirectories {

		if (strings.Compare(path, item) == 0) {

			return false
			break
		}
	}

	return true
}

func IsFileToCheck(path string) bool {

	extArr := strings.Split(path, ".")
	ext := extArr[len(extArr)-1]

	for _, item := range vars.ExceptionFileExtension {

		if (strings.Compare(ext, item) == 0) {

			return false
			break
		}
	}

	return true
}

func CheckFileUpdated(path string, timestampLastModification string) bool {

	// 	Make and firstTimeHandlFile, to load every files in 
	//  mapFiles without check if exist or not into array

	if (strings.Compare(vars.MapFiles[path], "") == 0) {

		// It doesn't exist, insert into mapFiles
		vars.MapFiles[path] = timestampLastModification
		execTask.ExecTask()

		// is unuseful fetch other files
		return true
	
	} else {

		// CHECK DATE MODIFIED
		if (strings.Compare(vars.MapFiles[path], timestampLastModification) == -1) {

			vars.MapFiles[path] = timestampLastModification
			execTask.ExecTask()

			// is unuseful fetch other files
			return true
		}
	}

	return false
}
