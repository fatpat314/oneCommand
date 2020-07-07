package main

import (
	"bytes"
	// "context"
	// "encoding/json"
    "bufio"
	"fmt"
	// "io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

    "github.com/briandowns/spinner"
)

func readFile(){
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter command file name: ")
    comFile, _ := reader.ReadString('\n')
    fmt.Print("Running command list " + comFile)
    comFile = strings.TrimSpace(comFile)
    return
}

func ExampleCmd_Run() {

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter command file name: ")
    comFile, _ := reader.ReadString('\n')
    fmt.Print("Running command list " + comFile, "\n")
    comFile = strings.TrimSpace(comFile)

    s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)  // Build our new spinner
    s.Color("red")
    s.Start()                                                    // Start the spinner
    time.Sleep(1 * time.Second)                                  // Run for some time to simulate work
    s.Color("green")
    time.Sleep(100*time.Millisecond)
    s.Stop()

    fileContents, err3 := ioutil.ReadFile(comFile + ".txt")
    if err3 != nil {
        panic(err3)
    }
    fmt.Print(string(fileContents))
    command := string(fileContents)

    com1 := " "
    com2 := " "
    com3 := " "
    com4 := " "
    com5 := " "
    com6 := " "
    cmd := exec.Command(string(com1), string(com2))

    // for each line in text
    commandMap := make(map[int]string)
    commandKey := strings.Split(command, "\n")
    keyLength := len(commandKey)

    key := 0
    for key < keyLength {
        commandMap[key] = commandKey[key]
        key += 1
    }

    n := 0
    valueList := len(commandMap)
    // fmt.Print("valueList: ", valueList, "\n")
    for n < valueList {
        values := string(commandMap[n])
        // fmt.Print("VALUES: ", values, "\n")
        list := strings.Split(values, " ")
        // fmt.Print("LIST: ", list, "\n")
        // fmt.Print("\n")

        switch {

        case len(list) == 1:
            if list[0] == ""{

            }else{
                fmt.Print("\n")
                com1 = list[0]
                com1 = strings.TrimSpace(com1)
                cmd = exec.Command(string(com1))
                var out bytes.Buffer
                cmd.Stdout = &out
                err := cmd.Run()
                test := cmd.Wait()
                if err != nil {
                    log.Fatal(err)
                    log.Fatal(test)
                }
                fmt.Printf(out.String())
            }

        case len(list) == 2:
            // fmt.Print("2", "\n")
            com1 = list[0]
            com2 = list[1]
            com1 = strings.TrimSpace(com1)
            com2 = strings.TrimSpace(com2)
            cmd = exec.Command(string(com1), string(com2))
            var out bytes.Buffer
            cmd.Stdout = &out
            err := cmd.Run()
            test := cmd.Wait()
            if err != nil {
        		log.Fatal(err)
                log.Fatal(test)
        	}
            fmt.Printf(out.String())
            // fmt.Print("\n")

        case len(list) == 3:
            fmt.Print("\n")
            com1 = list[0]
            com2 = list[1]
            com3 = list[2]
            com1 = strings.TrimSpace(com1)
            com2 = strings.TrimSpace(com2)
            com3 = strings.TrimSpace(com3)
            cmd = exec.Command(string(com1), string(com2), string(com3))
            var out bytes.Buffer
            cmd.Stdout = &out
            err := cmd.Run()
            test := cmd.Wait()
            if err != nil {
        		log.Fatal(err)
                log.Fatal(test)
        	}
            fmt.Printf(out.String())
            // fmt.Print("\n")

        case len(list) == 4:
            fmt.Print("\n")
            com1 = list[0]
            com2 = list[1]
            com3 = list[2]
            com4 = list[3]
            com1 = strings.TrimSpace(com1)
            com2 = strings.TrimSpace(com2)
            com3 = strings.TrimSpace(com3)
            com4 = strings.TrimSpace(com4)
            cmd = exec.Command(string(com1), string(com2), string(com3), string(com4))
            var out bytes.Buffer
            cmd.Stdout = &out
            err := cmd.Run()
            test := cmd.Wait()
            if err != nil {
        		log.Fatal(err)
                log.Fatal(test)
        	}

            fmt.Printf(out.String())
            // fmt.Print("\n")

        case len(list) == 5:
            fmt.Print("\n")
            com1 = list[0]
            com2 = list[1]
            com3 = list[2]
            com4 = list[3]
            com5 = list[4]
            com1 = strings.TrimSpace(com1)
            com2 = strings.TrimSpace(com2)
            com3 = strings.TrimSpace(com3)
            com4 = strings.TrimSpace(com4)
            com5 = strings.TrimSpace(com5)
            cmd = exec.Command(string(com1), string(com2), string(com3), string(com4), string(com5))
            var out bytes.Buffer
            cmd.Stdout = &out
            err := cmd.Run()
            test := cmd.Wait()
            if err != nil {
        		log.Fatal(err)
                log.Fatal(test)
        	}

            fmt.Printf(out.String())
        case len(list) == 6:
            fmt.Print("\n")
            com1 = list[0]
            com2 = list[1]
            com3 = list[2]
            com4 = list[3]
            com5 = list[4]
            com6 = list[5]
            com1 = strings.TrimSpace(com1)
            com2 = strings.TrimSpace(com2)
            com3 = strings.TrimSpace(com3)
            com4 = strings.TrimSpace(com4)
            com5 = strings.TrimSpace(com5)
            com6 = strings.TrimSpace(com6)
            cmd = exec.Command(string(com1), string(com2), string(com3), string(com4), string(com5), string(com6) )
            var out bytes.Buffer
            cmd.Stdout = &out
            err := cmd.Run()
            test := cmd.Wait()
            if err != nil {
        		log.Fatal(err)
                log.Fatal(test)
        	}
            fmt.Printf(out.String())
        }
        n += 1
    }
}

func main(){
    // readFile()
    ExampleCmd_Run()

    // reader := bufio.NewReader(os.Stdin)
    // fmt.Print("Enter your city: ")
    // city, _ := reader.ReadString('\n')
    // fmt.Print("You live in " + city)

}
