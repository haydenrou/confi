package main

import(
    "os"
    "io/ioutil"
    "fmt"
)

func main() {
    var HOME_PATH string = os.Getenv("HOME")
    var CONFIG_PATH string = "/.ssh/config"
    var SSH_CONFIG_PATH = HOME_PATH + CONFIG_PATH

    data, err := ioutil.ReadFile(SSH_CONFIG_PATH)


    if HOME_PATH == "" {
        fmt.Println("You must set ENV $HOME to your home path")
        return
    } else if os.IsNotExist(err) {
        fmt.Println("You must create a config file in $HOME/.ssh/")
        return
    }

    var SSH_CONFIG = string(data)

    fmt.Println(SSH_CONFIG)

    if len(os.Args) == 1 {
        fmt.Println("Please enter some arguments")
        return
    }

    fmt.Printf("Your first argument is: %v\n", os.Args[1])
}
