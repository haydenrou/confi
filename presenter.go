package main

import (
	"io/ioutil"
	"strings"
)

func BaseConfig() string {
	data, _ := ioutil.ReadFile(SSH_CONFIG_PATH)

	return string(data)
}

func ConfigMap() map[string]map[string]string {
	splitString := strings.Split(BaseConfig(), "Host ")

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
