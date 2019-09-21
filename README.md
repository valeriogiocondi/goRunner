# goRunner

A task runner for everything

## Introduction
This is a task runner written in Golang.
You can use goRunner alongside your project, no matter what program language you use, you just have to set config.json properly.

In this example goRunner is set to execute a simple python script everytime you change and save example.py, or some files into src directory.

## Config.json
This is the main, and only, configuration file.

You can set:

  *. list of tasks you want to execute
  *. types of file you don't want to watch
  *. directories you don't want to watch
  
  
 Example:
   
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
   
 
   *. the only task we want to execute is helloWorld.py
   *. we don't want to watch .DS_Sore and .txt files
   *. we don't want to watch ./goRunner/ and /files/dir-2/dir-2-1/ directories
   
   
