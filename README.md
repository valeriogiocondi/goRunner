# goRunner

A task runner for everything


## Introduction

This is a task runner developed in [Golang](https://golang.org/).
You don't need to know Golang to use goRunner.

You can use goRunner alongside your project, no matter what program language you use, you just have to set config.json properly.

In this example goRunner is set to execute a simple python script everytime you edit and save example.py, or some other files in *src* directory.


## Config.json

This is the main, and only, configuration file.

You can set:

  * list of tasks you want to execute
  * types of file you don't want to watch
  * directories you don't want to watch
  
   
```
   {
    "tasklist": [
        "python helloWorld.py"
      ],
      "exceptFile": [

        "DS_Store",
        "txt"
      ],
      "exceptDirectories": [
        "./goRunner/",
        "./files/dir-2/dir-2-1/"
      ]
    }
   
```
   
 
   * the only task we want to execute is helloWorld.py
   * we don't want to watch .DS_Sore and .txt files
   * we don't want to watch ./goRunner/ and /files/dir-2/dir-2-1/ directories
   
   
   ## Custom packages
   
   Remeber to move packages directory *src/packages/* into your Golang package path $GOPATH/src/ [(More infos)](https://github.com/golang/go/wiki/SettingGOPATH).
   
   
   ## License
   
   This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
