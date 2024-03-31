package file_manager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("failed to open file")
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	error2 := scanner.Err()
	if error2 != nil {

		fmt.Println(error2)
		err := file.Close()
		if err != nil {
			fmt.Println("closing file failed")
		}

		return nil, errors.New("failed to read file")
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("creating file failed")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("JSON conversion failed")
	}
	err = file.Close()
	if err != nil {
		return errors.New("closing file failed")
	}
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
