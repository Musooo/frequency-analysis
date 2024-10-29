package pkg

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileAndReturnString(fileName string) string{
	file, err := os.Open(fileName)
	var f string
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f += scanner.Text()
	}

	return f
}