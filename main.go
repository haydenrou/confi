package main

import (
	"fmt"
	"log"
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
	"ProxyJump",
	"ProxyCommand",
	"PubkeyAuthentication",
	"RemoteForward",
	"RemoteCommand",
	"RequestTTY",
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
	case "add":
		processAdd(argument, changes)
	case "delete":
		processDelete(argument)
	case "rename":
		processRename(argument, args[3])
	default:
		log.Fatal(option + " is not a valid argument")
	}
}

func processShow(argument string) {
	if argument == "all" {
		// use `log` here? But it shouldnt show exit status, as it wasn't fatal
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
	changedConfigMap := ConfigMap()

	for i := 0; i < len(changes); i++ {
		change := strings.Split(changes[i], "=")

		ValidateChange(change)

		changedConfigMap[argument][change[0]] = change[1]
	}

	changedConfigMap.writeConfig()
}

func processAdd(argument string, changes []string) {
	changedConfigMap := ConfigMap()
	changedConfigMap[argument] = map[string]string{}

	for i := 0; i < len(changes); i++ {
		change := strings.Split(changes[i], "=")

		ValidateChange(change)

		changedConfigMap[argument][change[0]] = change[1]
	}

	changedConfigMap.writeConfig()
}

func processDelete(argument string) {
	changedConfigMap := ConfigMap()

	ValidateExists(argument)

	delete(changedConfigMap, argument)

	changedConfigMap.writeConfig()
}

func processRename(argument string, changedName string) {
	changedConfigMap := ConfigMap()

	ValidateExists(argument)

	changedConfigMap[changedName] = changedConfigMap[argument]

	delete(changedConfigMap, argument)

	changedConfigMap.writeConfig()
}
