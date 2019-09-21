package main

import (
	"time"
	"packages/taskRunner"
)

func main() {

	app := taskRunner.Start{}
	app.Watch(500 * time.Millisecond)
}