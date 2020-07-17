package main

import "testing"

func TestCmdCreate(t *testing.T){
    var tests = []struct {
        input   string
        expected    string
        err     error
    }{
        { input: "something", expected: "something.txt", err: nil },
    }

    for _, test := range tests {
        // var output = CmdCreate()
        if  output := CmdCreate(test.input); output != test.expected {
            t.Error("Test Failed: {} inputted, {} expected, recived", test.input, test.expected, output)
        }
    }
}

// func BenchmarkCmdCreate(b *testing.B) {
//     for i := 0; i < b.N; i++ {
//         CmdCreate()
//
//     }
// }
