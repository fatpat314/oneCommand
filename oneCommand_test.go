package main

import "testing"

func TestCmdShow(t *testing.T) {
	var tests = []struct {
		expected string
		err      error
	}{
		{expected: CmdShow(), err: nil},
	}

	for _, test := range tests {
		// var output = CmdCreate()
		if output := CmdShow(); output != test.expected {
			t.Error("Test Failed:{} expected, recived", test.expected, output)
		}
	}
}

func TestCmdCreate(t *testing.T) {
	var tests = []struct {
		expected string
		err      error
	}{
		{expected: CmdCreate(), err: nil},
	}

	for _, test := range tests {
		if output := CmdCreate(); output != test.expected {
			t.Error("Test Failed:{} expected, recived", test.expected, output)
		}
	}
}

func BenchmarkCmdCreate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CmdCreate()

	}
}

func BenchmarkCmdShow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CmdShow()

	}
}
