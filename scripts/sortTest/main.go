package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// sortTest allows to sort function definitions
// depending on their call order defined in the test suite.
// It is useful to reorganise file after modification in the suite.
func main() {

	if len(os.Args) != 2 || os.Args[1] == "" {
		log.Fatalf("invalid argument")
	}

	name := os.Args[1]

	buf, err := ioutil.ReadFile("tests/integration/" + strings.ToLower(name) + "_test.go")
	if err != nil {
		log.Fatalf("failed to read file")
	}

	suiteName := strings.Title(name) + "Suite"
	testName := "Test" + strings.Title(name)
	funcName := "func (s *" + suiteName + ") " + testName + "()"
	index := strings.Index(string(buf), funcName)
	if index == -1 {
		log.Fatalf("failed to find test suite function")
	}
	firstPart, buf := buf[:index], buf[index:]

	testName = strings.SplitN(string(buf), "\"", 3)[1]
	funcName = "func (s *" + suiteName + ") " + testName + "()"
	index = strings.Index(string(buf), funcName)
	if index == -1 {
		log.Fatalf("failed to find test function")
	}
	secondPart, thirdPart := buf[:index], buf[index:]

	testList := []string{}
	testParts := strings.Split(string(secondPart), "\"")
	for i, test := range testParts {
		if i%2 == 1 {
			testList = append(testList, test)
		}
	}

	sortedThirdPart := []string{}
	functionList := strings.SplitAfter(string(thirdPart), "}\n\n")
	for i := 0; i < len(functionList); i++ {
		for _, function := range functionList {
			if strings.Contains(function, testList[i]) {
				sortedThirdPart = append(sortedThirdPart, function)
			}
		}
	}

	buf = bytes.Join([][]byte{firstPart, secondPart, []byte(strings.Join(sortedThirdPart, ""))}, []byte(""))

	err = ioutil.WriteFile("tests/integration/"+strings.ToLower(name)+"_generate_test.go", buf, 0644)
	if err != nil {
		log.Fatalf("failed to write file")
	}
}
