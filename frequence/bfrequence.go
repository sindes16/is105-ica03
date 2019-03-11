package bfrequence

import(
	"fmt"
	"os"
	"bufio"
)


func Countlines(filename string) {
file, _ := os.Open(filename)
fileScanner := bufio.NewScanner(file)
lineCount := 0
for fileScanner.Scan() {
    lineCount++
}
fmt.Println("number of lines:", lineCount)
}
