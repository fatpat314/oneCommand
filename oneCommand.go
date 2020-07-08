package main

import (
	"bytes"
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
// func Cmd_write(){
//
// }
func Cmd_create(){
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Name new command file: ")
    fileName, _ := reader.ReadString('\n')
    fileName = strings.TrimSpace(fileName)
    fileName = fileName + ".txt"
    os.Create(fileName)

    // Cmd_write
    writer := bufio.NewReader(os.Stdin)
    fmt.Print("Write Commands: ", "\n")
    fileWrite, _ := writer.ReadString('\n')
    f := strings.Replace(fileWrite, ", ", "\n", -1)
    fmt.Print(f, "\n")

    s := []byte(f)
    err := ioutil.WriteFile(fileName, s, 0777)
    if err != nil {
        fmt.Println(err)
    }
}

func Cmd_Show(){
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter command file name: ")
    fmt.Print("\n")
    comFile, _ := reader.ReadString('\n')
    comFile = strings.TrimSpace(comFile)
    comFile = comFile + ".txt"
    content, err := ioutil.ReadFile(comFile)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Print("", "\n")
    fmt.Print("Command list: ", "\n" + string(content), "\n")
}

func Cmd_Run() {

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter command file name: ")
    comFile, _ := reader.ReadString('\n')
    fmt.Print("Running command list: " + comFile, "\n")
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
    for n < valueList {
        values := string(commandMap[n])
        list := strings.Split(values, " ")

        switch {
        case len(list) == 1:
            if list[0] == ""{
                // do nothing
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
            fmt.Print("\n")
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
    starter := bufio.NewReader(os.Stdin)
    fmt.Print("Enter 'Run' to do run a command file.", "\n", "Enter 'New' to create a new command file.", "\n", "Enter 'Show' to see the contents of a command file.", "\n", "Enter 'Help' for more information.", "\n", "Enter: ")
    // fmt.Print("Enter 'N' to create a new file.", "\n")
    a, _ := starter.ReadString('\n')
    a = strings.TrimSpace(a)
    // fmt.Print(a)
    switch {
    case string(a) == "RUN":
        Cmd_Run()
    case string(a) == "Run":
        Cmd_Run()
    case string(a) == "run":
        Cmd_Run()

    case string(a) == "NEW":
        Cmd_create()
    case string(a) == "New":
        Cmd_create()
    case string(a) == "new":
        Cmd_create()

    case string(a) == "Help":
        fmt.Print("\n", "When you are creating a new command file the name of that file is used to create a .txt file. The command lines you enter are then processed to create each command as a seperate line in the command file written in sequence", "\n")
        fmt.Print("\n", "When you are running an already existing command file the lines of commands that are stored in the corresponding .txt are executed in sequence.", "\n")
    case string(a) == "help":
        fmt.Print("\n", "When you are creating a new command file the name of that file is used to create a .txt file. The command lines you enter are then processed to create each command as a seperate line in the command file written in sequence", "\n")
        fmt.Print("\n", "When you are running an already existing command file the lines of commands that are stored in the corresponding .txt are executed in sequence.", "\n")

    case string(a) == "SHOW":
        Cmd_Show()
    case string(a) == "Show":
        Cmd_Show()
    case string(a) == "show":
        Cmd_Show()
}
    // readFile()
    // ExampleCmd_Run()

    // reader := bufio.NewReader(os.Stdin)
    // fmt.Print("Enter your city: ")
    // city, _ := reader.ReadString('\n')
    // fmt.Print("You live in " + city)

}
