package main

import (
	"fmt"
	"os"
	"os/user"

	"./repl"
)

//main function, obviously , this is the ENTRY to this program.
func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hey, %s! XInterpreter's here.\n", user.Username)
	fmt.Printf("\n")
	repl.Start(os.Stdin, os.Stdout)
}
