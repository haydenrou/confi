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

	// We don't want to have a newline at the top of the file...
	// this isn't great - but it does the job!
	var count int

	for key, value := range config {
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

		for config, confValue := range value {
			_, err := writer.WriteString("  " + config + " " + confValue + "\n")

			if err != nil {
				log.Fatalf("Error whilst writing to file: %s", err.Error())
			}
		}
	}

	writer.Flush()
}
