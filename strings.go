//	Package datafile allows reading data samples from files.
package datafile

import (
	"bufio"
	"os"
)

//	GetStrings reads a string from each line of a file.
func GetStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}
