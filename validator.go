package main

import (
    "log"
    "os"
    "io/ioutil"
)

func CheckValidity() {
    _, err := ioutil.ReadFile(SSH_CONFIG_PATH)

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

func ValidateChange(change []string) bool {
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

