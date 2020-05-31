package main

import(
    "os"
    "io/ioutil"
    "fmt"
    "strings"
)

func main() {
    var HOME_PATH string = os.Getenv("HOME")
    var CONFIG_PATH string = "/.ssh/config"
    var SSH_CONFIG_PATH = HOME_PATH + CONFIG_PATH

    data, err := ioutil.ReadFile(SSH_CONFIG_PATH)

    if HOME_PATH == "" {
        fmt.Println("You must set ENV $HOME to your home path")
        return
    } else if strings.Contains(err.Error(), "no such file or directory") {
        fmt.Println("You must create a config file in $HOME/.ssh/")
        return
    }

    fmt.Println(data)

    if len(os.Args) == 1 {
        fmt.Println("Please enter some arguments")
        return
    }

    fmt.Printf("Your first argument is: %v\n", os.Args[1])
}
