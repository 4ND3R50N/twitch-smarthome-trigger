package project

import (
	"flag"
	"fmt"
)

var Path string

func InitializeProjectPath() {
	flag.StringVar(&Path, "PROJECT_PATH", "/", "Provide project path as an absolute path")
	flag.Parse()
	fmt.Println("Project path is: " + Path)
}
