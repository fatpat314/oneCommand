package main

import (
	"bytes"
    "bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
    "flag"

    "github.com/briandowns/spinner"
    "github.com/gookit/color"
)

// CmdCreate takes command line input from the user that is used to generate
// a .txt file which stores the commands you with to run.
func CmdCreate(){
    reader := bufio.NewReader(os.Stdin)
    color.Green.Print("Name new command file: ")
    fileName, _ := reader.ReadString('\n')
    fileName = strings.TrimSpace(fileName)
	if fileName == ""{
		fileName = "TEST"
	}
    fileName = fileName + ".txt"
    os.Create(fileName)

    // Cmdwrite
	// CmdWrite takes command line input and uses it to write Commands
	// into the newly created .txt file
    writer := bufio.NewReader(os.Stdin)
    color.Green.Print("Write Commands: ", "\n")
    fileWrite, _ := writer.ReadString('\n')
    f := strings.Replace(fileWrite, ", ", "\n", -1)
    fmt.Printf(f, "\n")

    s := []byte(f)
    err := ioutil.WriteFile(fileName, s, 0777)
    if err != nil {
        color.Red.Println(err)
    }
}
// CmdShow takes user input to find an exsisting file and prints each line
// of commands so the user can read the list of commands.
func CmdShow(){
    reader := bufio.NewReader(os.Stdin)
    color.Green.Print("Enter command file name: ")
    fmt.Print("\n")
    comFile, _ := reader.ReadString('\n')
    comFile = strings.TrimSpace(comFile)
    comFile = comFile + ".txt"
    content, err := ioutil.ReadFile(comFile)
    if err != nil {
        color.Red.Print(err)
    }
    fmt.Print("", "\n")
    color.Green.Print(comFile, " Command list: ")
    fmt.Print("\n" + string(content), "\n")
}

// CmdList filters through the directory to find all .txt and prints the list
// 	out to the user.
func CmdList(){
    dir := flag.String("dir", ".", "Find all.txt")
    file, err := os.Open(*dir)
    if err != nil{
        color.Red.Print("failed opening directory: %s", err)
    }
    defer file.Close()
    list, _ := file.Readdirnames(0)
    i := 0
    for i < len(list){
        if strings.Contains(list[i], ".txt"){
            fmt.Print(list[i], "\n")
        }
        i += 1
    }
    fmt.Print("\n")
}

//CmdHelp is just used to give extra instruction that are printed to the command line
func CmdHelp(){
    color.Green.Print("\n", "ENTER: NEW", "\n", "Creates a new command file. You can then name the command file and enter each command, line by line, separated by a ', ' for each command line.", "\n")
    color.Green.Print("\n", "ENTER: RUN", "\n", "Runs a existing command file. Choose the command to run by the name of the command file.", "\n")
    color.Green.Print("\n", "ENTER: LIST", "\n", "Shows a list of all of the of the existing command files", "\n")
    color.Green.Print("\n", "ENTER: SHOW", "\n", "Shows the contents of an existing command file. The file that is shown is chosen by name", "\n")
    color.Green.Print("\n", "ENTER: END", "\n", "Exits the program", "\n", "\n")
}

// CmdRun takes user input to determine which .txt file to run and exicutes the Commands
// in that file in the order they appear.
func CmdRun() {
    reader := bufio.NewReader(os.Stdin)
    color.Green.Print("Enter command file name: ")
    comFile, _ := reader.ReadString('\n')
    color.Green.Print("Running command list: ")
    fmt.Print(comFile, "\n")
    comFile = strings.TrimSpace(comFile)

	// Pretty colors
    s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)  // Build our new spinner
    s.Color("red")
    s.Start()                                                    // Start the spinner
    time.Sleep(1 * time.Second)                                  // Run for some time to simulate work
    s.Color("green")
    time.Sleep(100*time.Millisecond)
    s.Stop()

    fileContents, err3 := ioutil.ReadFile(comFile + ".txt")
    if err3 != nil {
        color.Red.Print(err3)
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
                    color.Red.Print(err)
                    color.Red.Print(test)
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
                color.Red.Print(err)
                color.Red.Print(test)
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
                color.Red.Print(err)
                color.Red.Print(test)
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
                color.Red.Print(err)
                color.Red.Print(test)
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
                color.Red.Print(err)
                color.Red.Print(test)
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
                color.Red.Print(err)
                color.Red.Print(test)
        	}
            fmt.Printf(out.String())
        }
        n += 1
    }
}

func main(){

    t := true
    for t == true {
        starter := bufio.NewReader(os.Stdin)
        color.Green.Print("\n", "Enter 'Run' to do run a command file.", "\n", "Enter 'Help' for more information.", "\n", "Enter 'End' to end the program.", "\n", "\n", "Enter: ")
        a, _ := starter.ReadString('\n')
        a = strings.TrimSpace(a)
        switch {
        case string(a) == "RUN":
            CmdRun()
        case string(a) == "Run":
            CmdRun()
        case string(a) == "run":
            CmdRun()

        case string(a) == "NEW":
            CmdCreate()
        case string(a) == "New":
            CmdCreate()
        case string(a) == "new":
            CmdCreate()

        case string(a) == "HELP":
            CmdHelp()
        case string(a) == "Help":
            CmdHelp()
        case string(a) == "help":
            CmdHelp()

        case string(a) == "SHOW":
            CmdShow()
        case string(a) == "Show":
            CmdShow()
        case string(a) == "show":
            CmdShow()

        case string(a) == "LIST":
            CmdList()
        case string(a) == "List":
            CmdList()
        case string(a) == "list":
            CmdList()

        case string(a) == "END":
            t = false
        case string(a) == "End":
            t = false
        case string(a) == "end":
            t = false
        }
    }
}
