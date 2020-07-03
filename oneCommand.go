package main

import (
	"bytes"
	// "context"
	// "encoding/json"
    // "bufio"
	"fmt"
	// "io"
	// "io/ioutil"
	"log"
	// "os"
	"os/exec"
	"strings"
	// "time"
)

func ExampleCmd_Run() {

    cmd := exec.Command("git", "add", ".")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out

    cmd2 := exec.Command("git", "commit", "-m", "automated")

    var out2 bytes.Buffer
    cmd2.Stdout = &out2


	err := cmd.Run()
    err2 := cmd2.Run()


	if err != nil {
		log.Fatal(err)
	}

    if err2 != nil {
        log.Fatal(err2)
    }
	fmt.Printf(out.String())
    fmt.Printf(out2.String())
}

func main(){

    ExampleCmd_Run()

    // reader := bufio.NewReader(os.Stdin)
    // fmt.Print("Enter your city: ")
    // city, _ := reader.ReadString('\n')
    // fmt.Print("You live in " + city)

}
