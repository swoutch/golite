package main

import (
	"strings"
	"testing"
)

func TestCli(t *testing.T) {

	type testCase struct {
		input               string
		expectedDisplay     string
		expectedReturnValue bool
	}

	bigString := string(make([]byte, 10000)) + "\n"
	stringWithExitCommandAtByte128 := string(make([]byte, 128)) + ".exit\n"

	testCases := []testCase{
		testCase{"hello\n", "Unrecognized command: hello\n", true},
		testCase{"how are you?\n", "Unrecognized command: how are you?\n", true},
		testCase{bigString, "Unrecognized command: " + bigString, true},
		testCase{".exit\n", "", false},
		testCase{stringWithExitCommandAtByte128, "Unrecognized command: " + stringWithExitCommandAtByte128, true},
	}

	for _, c := range testCases {
		buffer := strings.Builder{}
		actualReturnValue := cli(strings.NewReader(c.input), &buffer)
		actualDisplay := buffer.String()

		if actualDisplay != c.expectedDisplay {
			t.Errorf("Display error: got '%q', expected '%q'", actualDisplay, c.expectedDisplay)
		}
		if actualReturnValue != c.expectedReturnValue {
			t.Errorf("Return value error: got '%v', expected '%v'", actualReturnValue, c.expectedReturnValue)
		}
	}
}
