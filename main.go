package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
    "log"
    "bufio"
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
    _, err := ioutil.ReadFile(SSH_CONFIG_PATH)

    checkValidity(err)

    processStatement(os.Args)
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

func checkValidity(err error) {
    var errorMessage string

    if HOME_PATH == "" {
        errorMessage = "You must set ENV $HOME to your home path"
    } else if os.IsNotExist(err) {
        errorMessage = "You must create a config file in $HOME/.ssh/"
    } else if os.Args[1] == "edit" && len(os.Args) == 3 {
        errorMessage = "You must provide changes"
    } else if len(os.Args) < 3 {
        errorMessage = "You must enter a valid statement"
    }

    if errorMessage != "" {
        log.Fatal(errorMessage)
    }
}

func validateChange(change []string) bool {
    if change[1] == "" {
        log.Fatalf("Cannot set %v to blank", change[0])

        return false
    }

    for _, val := range VALID_CONFIGS {
        if change[0] == val {
            return true
        }
    }

    log.Fatalf("%v is not a valid configuration", change[0])

    return false
}

func writeConfig(config map[string]map[string]string) {
    file, err := os.Create(SSH_CONFIG_PATH)

    if err != nil {
        log.Fatal(err)
    }

    writer := bufio.NewWriter(file)

    for key, value := range config {
        if key == "" { continue }

        _, err := writer.WriteString("\nHost " + key + "\n")

        if err != nil {
            log.Fatalf("Error whilst writing to file: %s", err.Error())
        }

        for config, confValue := range value {
            _, err := writer.WriteString("  " + config + " " + confValue + "\n")

            if err != nil {
                log.Fatalf("Error whilst writing to file: %s", err.Error())
            }
        }
    }

    writer.Flush()
}

func processStatement(args []string) {
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
        fmt.Println(sshConfig())
        return
    }

    if v, ok := mappedSsh()[argument]; ok {
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

        validateChange(change)

        splitChanges = append(splitChanges, change)
    }

    changedMap := mappedSsh()

    for i := 0; i < len(splitChanges); i++ {
        changedMap[argument][splitChanges[i][0]] = splitChanges[i][1]
    }

    writeConfig(changedMap)
}
