package main

import (
	"bufio"
	"log"
	"os"
)

func WriteConfig(config map[string]map[string]string) {
	file, err := os.Create(SSH_CONFIG_PATH)

	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriter(file)

	for key, value := range config {
		if key == "" {
			continue
		}

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
