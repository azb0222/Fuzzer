package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	file, err := os.ReadFile("./test02.pdf")
	if err != nil {
		fmt.Printf("%v", err)
	}

	for i := 0; i < len(file); i++ {
		ChangeByte(i)
	}

}

func ChangeByte(byteNum int) {
	// Make a new test file
	cmd := exec.Command("cp", "./test02.pdf", "./test.pdf")
	cmd.Run()

	// Read file
	file, err := os.ReadFile("./test.pdf")
	if err != nil {
		fmt.Printf("%v", err)
	}

	// Change byte to 0
	file[byteNum] = 0x00
	err = os.WriteFile("./test.pdf", file, 0644)
	if err != nil {
		fmt.Printf("%v", err)
	}

	// Record crashes
	cmd = exec.Command("./pdf_parser", "test.pdf")
	runError := cmd.Run()
	if runError != nil {
		// Record what byte the program crashes at
		fmt.Printf("CRASHED AT %d:\n", byteNum)
		// Record what caused the program to crash
		fmt.Printf("ERROR: %v\n", runError)

		// Open a new file to keep track of crashes
		crashTracker, err := os.OpenFile("crashes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("%v", err)
		}
		defer crashTracker.Close()

		output := fmt.Sprintf("CRASHED AT %d: \n ERROR: %v \n", byteNum, runError)

		if _, err := crashTracker.WriteString(output); err != nil {
			fmt.Printf("ERROR: %v \n", err)
		}
	}

}
