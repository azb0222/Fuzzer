package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	filePath = "/home/asritha/HackUmass/TestInput/testInput"
)

func main() {
	data, err := os.ReadFile("testInput")

	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	for i := 0; i < len(data); i++ {
		go ChangeByte(i)
	}
}

func ChangeByte(byteNum int) {
	//Copy file into TestingInput
	testFilePath := "./TestingInput/test"
	cmd := exec.Command("cp", "./testInput", testFilePath)
	cmd.Run()

	//Read file
	data, err := os.ReadFile(testFilePath)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	// Change byte to 0
	data[byteNum] = 0x00
	err = os.WriteFile(testFilePath, data, 0644)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	//Open new file
	crashTracker, err := os.Create("./testOutput/crashes.txt")
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	defer crashTracker.Close()

	//Record if crash in a separate file
	cmd = exec.Command("./TestingInput/test")
	err = cmd.Run()
	if err != nil {
		//Record what byte
		fmt.Printf("CRASHED AT %d:\n", byteNum)
		fmt.Printf("ERROR: %v\n", err)
		crashTracker.WriteString(err.Error())
	}

}
