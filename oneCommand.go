package main

import (
	"bytes"
	// "context"
	// "encoding/json"
    "bufio"
	"fmt"
	// "io"
	// "io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	// "time"
)

func ExampleCmd_Run() {
    cmd := exec.Command("git", "init")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(out.String())
}

func main(){

    ExampleCmd_Run()

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your city: ")
    city, _ := reader.ReadString('\n')
    fmt.Print("You live in " + city)

}
