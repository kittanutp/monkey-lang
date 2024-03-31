package main

import (
	"fmt"
	"monkey-lang/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Monkey-Language by MKTP called by %s", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
