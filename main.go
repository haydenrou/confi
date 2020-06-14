package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

const (
    CONFIG_PATH = "/.ssh/config"
)

var HOME_PATH = os.Getenv("HOME")
var SSH_CONFIG_PATH = HOME_PATH + CONFIG_PATH

func main() {
    _, err := ioutil.ReadFile(SSH_CONFIG_PATH)

    errMessage := checkValidity(err)

    if errMessage != "" {
        fmt.Println(errMessage)
        return
    } else {
        processStatement(os.Args)
    }
}

func sshConfig() string {
    data, _ := ioutil.ReadFile(SSH_CONFIG_PATH)

    return string(data)
}

func mappedSsh() map[string]map[string]string {
    splitString := strings.Split(sshConfig(), "Host ")

    splitConfigs := make([][]string, len(splitString), len(splitString))

    sshMap := map[string]map[string]string{}

    for i := 0; i < len(splitString); i++ {
        splitConfigs[i] = strings.Split(splitString[i], "\n")
    }

    for i := 0; i < len(splitConfigs); i++ {
        valueMap := map[string]string{}

        for x := 0; x < len(splitConfigs[i]); x++ {
            vals := strings.Split(strings.TrimSpace(string(splitConfigs[i][x])), " ")

            if len(vals) > 1 {
                valueMap[vals[0]] = vals[1]
            }
        }

        sshMap[strings.TrimSpace(splitConfigs[i][0])] = valueMap
    }

    return sshMap
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

    if v, ok := mappedSsh()[argument]; ok {
        fmt.Printf("Host\t%v\n", argument)

        for k, v := range v {
            fmt.Printf("\t%v\t %v\n", k, v)
        }
    }
}

