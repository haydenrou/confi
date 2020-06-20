package main

import (
	"io/ioutil"
	"strings"
	"bufio"
	"log"
	"os"
	"sort"
)

type config map[string]map[string]string

func BaseConfig() string {
	data, _ := ioutil.ReadFile(SSH_CONFIG_PATH)

	return string(data)
}

func ConfigMap() config {
	splitString := strings.Split(BaseConfig(), "Host ")

	splitConfigs := make([][]string, len(splitString), len(splitString))

	sshMap := make(config)

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

func (c config) writeConfig() {
	conf := c.alphabeticalize()

	file, err := os.Create(SSH_CONFIG_PATH)

	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriter(file)

	// We don't want to have a newline at the top of the file...
	// this isn't great - but it does the job!
	var count int

	for key, value := range conf {
		if key == "" {
			continue
		}

		var hostString string

		switch count {
		case 0:
			hostString = "Host "
		default:
			hostString = "\nHost "
		}
		count++

		_, err := writer.WriteString(hostString + key + "\n")

		if err != nil {
			log.Fatalf("Error whilst writing to file: %s", err.Error())
		}

		for configKey, confValue := range value {
			_, err := writer.WriteString("  " + configKey + " " + confValue + "\n")

			if err != nil {
				log.Fatalf("Error whilst writing to file: %s", err.Error())
			}
		}
	}

	writer.Flush()
}

func (c config) alphabeticalize() config {
	newConf := make(config)
	keys := make([]string, 0, len(c))

	for key, _ := range c {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		newConf[key] = c[key]
	}

	return newConf
}
