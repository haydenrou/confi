package main

import(
    "io/ioutil"
    "strings"
    "os"
    "bufio"
    "log"
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

func WriteConfig(config map[string]map[string]string) {
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

