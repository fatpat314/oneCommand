package main

import (
	"bytes"
	// "context"
	// "encoding/json"
    // "bufio"
	"fmt"
	// "io"
	"io/ioutil"
	"log"
	// "os"
	"os/exec"
	"strings"
	// "time"
)

func readFile(){
    fileContents, err := ioutil.ReadFile("text2.txt")
    if err != nil {
        panic(err)
    }
    fmt.Print(string(fileContents))
    return
}

func ExampleCmd_Run() {
    fileContents, err3 := ioutil.ReadFile("text2.txt")
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
    fmt.Print("valueList: ", valueList)
    for n < valueList {
        values := string(commandMap[n])
        fmt.Print("VALUES: ", values, "\n")
        list := strings.Split(values, " ")
        fmt.Print("LIST: ", list, "\n")
        fmt.Print("\n")

        switch {
        case len(list) == 2:
            fmt.Print("2", "\n")
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
            fmt.Print("\n")

        case len(list) == 3:
            fmt.Print("3", "\n")
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
            fmt.Print("\n")

        case len(list) == 4:
            fmt.Print("4", "\n")
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
            fmt.Print("\n")

        case len(list) == 5:
            fmt.Print("5", "\n")
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
            if err != nil {
                log.Fatal(err)
            }

            fmt.Printf(out.String())
            fmt.Print("\n")

        case len(list) == 6:
            fmt.Print("6", "\n")
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
            com6 = strings.TrimSpace(com5)
            cmd = exec.Command(string(com1), string(com2), string(com3), string(com4), string(com5), string(com6) )
            var out bytes.Buffer
            cmd.Stdout = &out
            err := cmd.Run()
            if err != nil {
                log.Fatal(err)
            }

            fmt.Printf(out.String())
            fmt.Print("\n")

        }

        // if len(list) == 2{
        //     print("CEHCK", "\n")
        //     com1 = list[0]
        //     com2 = list[1]
        //     com1 = strings.TrimSpace(com1)
        //     com2 = strings.TrimSpace(com2)
        //     cmd = exec.Command(string(com1), string(com2))
        // }


    	// var out bytes.Buffer
    	// cmd.Stdout = &out
        //
        // err := cmd.Run()
        // if err != nil {
    	// 	log.Fatal(err)
    	// }
        //
        // fmt.Printf(out.String())
        //
        //
        //
        // fmt.Print("\n")
        //
        n += 1
    }



    // commandValue := strings.Split
    // value := 0


    // commandValue := strings.Split(commandKey, ",")
    // fmt.Print("Value ", commandValue)


    // commandMap := make(map[int]string)
    // commandMap[commandKey[0]] = "check"
    // fmt.Print(s)

    // keyLength := len(commandKey)
    // key := 0
    // valueLength := len(commandValue)
    // value := 0
    // for key < keyLength {
    //     // commandMap[key] = commandValue[value]
    //     for value < valueLength{
    //         commandValue := strings.Split(command, ",")
    //         commandMap[key] = commandValue[value]
    //         value += 1
    //     }
    //
    //     // commandMap[key] = commandValue[n]
    //     // fmt.Print("MAP: ", commandMap, '\n')
    //
    //     // com1 = commandMap[0]
    //     // com2 = commandMap[1]
    //
    //     // command = commandKey[n]
    //     // if command == commandKey[0]{
    //     //     com1 = command
    //     // }
    //     // if command == commandKey[1]{
    //     //     com2 = command
    //     // }
    //     key += 1
    // // }
    // fmt.Print("MAP: ", commandMap)
    //
    // fmt.Println("END: ", com1, com2)
    // // com1 = commandMap[0]
    // // com1 = "cat"
    // com1 = strings.TrimSpace(com1)
    // // com2 = commandMap[0:1]
    // // com2 = "text.txt"
    // com2 = strings.TrimSpace(com2)
    //
    //
    //
    // cmd := exec.Command(string(com1), string(com2))
	// var out bytes.Buffer
	// cmd.Stdout = &out
    //
    // // cmd2 := exec.Command("git", "commit", "-m", "automated")
    // // var out2 bytes.Buffer
    // // cmd2.Stdout = &out2
    // //
    // // cmd3 := exec.Command("git", "push", "origin", "master")
    // // var out3 bytes.Buffer
    // // cmd3.Stdout = &out3
    //
    //
	// err := cmd.Run()
    // // err2 := cmd2.Run()
    // // err3 := cmd3.Run()
    //
    //
	// if err != nil {
	// 	log.Fatal(err)
	// }
    //
    // // if err2 != nil {
    // //     log.Fatal(err2)
    // // }
    // //
    // // if err3 != nil {
    // //     log.Fatal(err3)
    // // }
    //
	// fmt.Printf(out.String())
    // fmt.Printf(out2.String())
    // fmt.Printf(out3.String())
}

func main(){
    // readFile()
    ExampleCmd_Run()

    // reader := bufio.NewReader(os.Stdin)
    // fmt.Print("Enter your city: ")
    // city, _ := reader.ReadString('\n')
    // fmt.Print("You live in " + city)

}
