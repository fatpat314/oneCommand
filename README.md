[![Go Report Card](https://goreportcard.com/badge/github.com/fatpat314/oneCommand)](https://goreportcard.com/report/github.com/fatpat314/oneCommand)


# H1 oneCommand 



FUNCTIONS

func CmdCreate()
    CmdCreate takes command line input from the user that is used to generate a
    .txt file which stores the commands you with to run.

func CmdHelp()
    CmdHelp is just used to give extra instruction that are printed to the
    command line

func CmdList()
    CmdList filters through the directory to find all .txt and prints the list

        out to the user.

func CmdRun()
    CmdRun takes user input to determine which .txt file to run and exicutes the
    Commands in that file in the order they appear.

func CmdShow()
    CmdShow takes user input to find an exsisting file and prints each line of
    commands so the user can read the list of commands.
