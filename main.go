package main

import (
    "fmt"
    "os"
    "strings"
)

const (
    CONFIG_PATH = "/.ssh/config"
)

var HOME_PATH = os.Getenv("HOME")
var SSH_CONFIG_PATH = HOME_PATH + CONFIG_PATH

var VALID_CONFIGS = []string{
    "BindAddress",
    "ForwardAgent",
    "ForwardX11",
    "ForwardX11Trusted",
    "GatewayPorts",
    "HostName",
    "IdentityFile",
    "IdentitiesOnly",
    "LocalCommand",
    "LocalForward",
    "LogLevel",
    "PasswordAuthentication",
    "Port",
    "PreferredAuthentications",
    "Protocol",
    "ProxyCommand",
    "PubkeyAuthentication",
    "RemoteForward",
    "Tunnel",
    "TunnelDevice",
    "UsePrivilegedPort",
    "User",
    "UserKnownHostsFile",
    "VerifyHostKeyDNS",
    "VisualHostKey",
    "XAuthLocation",
}

func main() {
    CheckValidity()

    processStatement()
}

func processStatement() {
    args := os.Args

    option := args[1]
    argument := args[2]
    changes := args[3:]

    switch option {
    case "show":
        processShow(argument)
    case "edit":
        processEdit(argument, changes)
    }
}

func processShow(argument string) {
    if argument == "all" {
        fmt.Println(BaseConfig())
        return
    }

    if v, ok := ConfigMap()[argument]; ok {
        fmt.Printf("Host %v\n", argument)

        for key, val := range v {
            fmt.Printf("  %v %v\n", key, val)
        }
    }
}

func processEdit(argument string, changes []string) {
    splitChanges := [][]string{}

    for i := 0; i < len(changes); i++ {
        change := strings.Split(changes[i], "=")

        ValidateChange(change)

        splitChanges = append(splitChanges, change)
    }

    changedMap := ConfigMap()

    for i := 0; i < len(splitChanges); i++ {
        changedMap[argument][splitChanges[i][0]] = splitChanges[i][1]
    }

    WriteConfig(changedMap)
}
