package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/kasperbe/monkey/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Hello %s!\n", user.Username)
    fmt.Println("Feel free to type in commands")
    repl.Start(os.Stdin, os.Stdout)
}
