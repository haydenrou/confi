package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

const CONFIG_PATH = "/.ssh/config"

var HOME_PATH = os.Getenv("HOME")
var SSH_CONFIG_PATH = HOME_PATH + CONFIG_PATH

func main() {
    _, err := ioutil.ReadFile(SSH_CONFIG_PATH)

    errMessage := checkValidity(err)

    if errMessage != "" {
        fmt.Println(errMessage)
        return
    } else if len(os.Args) >= 3 {
        processStatement(os.Args)
    }
}

func sshConfig() string {
    data, _ := ioutil.ReadFile(SSH_CONFIG_PATH)

    return string(data)
}

func checkValidity(err error) string {
    if HOME_PATH == "" {
        return "You must set ENV $HOME to your home path"
    } else if os.IsNotExist(err) {
        return "You must create a config file in $HOME/.ssh/"
    } else if len(os.Args) < 3 {
        return "You must enter a valid statement"
    } else {
        return ""
    }
}

func processStatement(args []string) {
    option := args[1]
    argument := args[2]

    if option == "show" {
        processShow(argument)
        return
    }
}

func processShow(argument string) {
    if argument == "all" {
        fmt.Println(sshConfig())
        return
    }
}

