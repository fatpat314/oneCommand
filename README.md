[![Go Report Card](https://goreportcard.com/badge/github.com/fatpat314/oneCommand)](https://goreportcard.com/report/github.com/fatpat314/oneCommand)

Project Pitch Slides:
https://docs.google.com/presentation/d/10gHsr8NU0YF8xOThCITkY3GfHkU_TBObHJQV45cbxS0/edit?usp=sharing

# oneCommand

oneCommand is a simple command line tool that distill sequences of commands into a single word or line.


## Usage
RUN: Runs an existing command sequence.

NEW: Lets a user create a new command sequence.

HELP: Gives the user more instruction.

END: Ends the program.



### FUNCTIONS

func CmdCreate()
    CmdCreate takes command line input from the user that is used to generate a
    .txt file which stores the commands you with to run.

func CmdHelp()
    CmdHelp is just used to give extra instruction that are printed to the
    command line

func CmdList()
    CmdList filters through the directory to find all .txt and prints the list out to the user.

func CmdRun()
    CmdRun takes user input to determine which .txt file to run and exicutes the
    Commands in that file in the order they appear.

func CmdShow()
    CmdShow takes user input to find an exsisting file and prints each line of
    commands so the user can read the list of commands.
