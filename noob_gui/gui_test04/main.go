package main

import (
	"fmt"
	"ioutil"
    "github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyMainWindow struct {
    *walk.MainWindow
    te *walk.TextEdit
}

func main() {
	defer func() {
        if err := recover(); err != nil {
            errMsg := fmt.Sprintf("%#v", err)
            ioutil.WriteFile("fuck.log", []byte(errMsg), 0644)
        }
    }()

	myWindow = &MyMainWindow{model: NewEnvModel()}

}