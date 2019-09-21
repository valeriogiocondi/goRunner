package taskRunner

import (
	"time"

	"packages/vars"
	"packages/config"
	"packages/filesListPkg"
)


type Start struct {
}

func(self *Start) Watch(ms time.Duration) {

	config.Load();

	for {

		filesListPkg.GetFilesList(vars.InitPath)
		time.Sleep(ms)
	}
}
