// package declaration
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var FILE_NAME = "storage.txt"
var CREATED_AT = time.Now().Format(time.RFC822)
var ID int

func main() {

	ID = 0

	scan := bufio.NewScanner(os.Stdin)

	quit := false

	for !quit {

		fmt.Println()
		fmt.Println("TRY SAVE SOMETHING IN TXT FILE.")
		fmt.Println("1) Save.")
		fmt.Println("2) Show.")
		fmt.Println("3) Update.")
		fmt.Println("4) Delete.")
		fmt.Println("0) QUIT!!!")
		fmt.Println()

		scan.Scan()
		opt := scan.Text()

		switch opt {
		case "1":
			fmt.Println("==============================")
			fmt.Println("              SAVE             ")
			fmt.Print("Type: ")

			textScan := bufio.NewScanner(os.Stdin)
			textScan.Scan()

			ID++

			text := fmt.Sprint(ID) + ") " + textScan.Text() + " | " + CREATED_AT

			fmt.Println("==============================")

			save(text)
		case "2":
			fmt.Println("==============================")
			fmt.Println("         SHOWING ALL          ")

			fmt.Println(getAll())
			fmt.Println("==============================")

		case "3":
			fmt.Println("==============================")
			fmt.Println("            UPDATE           ")

			fmt.Println("ID: ")
			idScan := bufio.NewScanner(os.Stdin)
			idScan.Scan()
			id, err := strconv.ParseInt(idScan.Text(), 10, 64)

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("UPDATE: ")
			textScan := bufio.NewScanner(os.Stdin)
			textScan.Scan()
			newText := fmt.Sprint(ID) + ") " + textScan.Text() + " | " + CREATED_AT

			update(int(id), newText)
			fmt.Println("==============================")

		case "4":

			fmt.Println("==============================")
			fmt.Print("LINE: ")

			lineScan := bufio.NewScanner(os.Stdin)
			lineScan.Scan()

			line, err := strconv.ParseInt(lineScan.Text(), 10, 0)

			if err != nil {
				fmt.Println(err)
			}

			delete(int(line))
			fmt.Println("==============================")
			fmt.Println("            DELETED           ")
			fmt.Println("==============================")

		case "0":
			quit = true
		}

	}

}

func save(data string) string {

	writeFile, err := os.OpenFile(FILE_NAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
	}

	if _, err := writeFile.WriteString(data + "\n"); err != nil {
		fmt.Println(err)
	}

	defer writeFile.Close()

	return "Data stored!"
}

func getAll() string {
	readFile, err := os.ReadFile(FILE_NAME)

	if err != nil {
		fmt.Println(err)
	}

	return string(readFile)
}

func update(id int, text string) bool {

	content, err := os.ReadFile(FILE_NAME)

	if err != nil {
		fmt.Println(err)
	}

	//transform to lines
	lines := strings.Split(string(content), "\n")

	//check if line exists
	if id < 1 || id > len(lines) {
		fmt.Println("Line doesnt exist")
	}

	//replace lines
	lines[id-1] = text

	//rewrite the file
	newText := strings.Join(lines, "\n")

	os.WriteFile(FILE_NAME, []byte(newText), 0644)

	return true
}

func delete(line int) bool {

	content, err := os.ReadFile(FILE_NAME)

	if err != nil {
		fmt.Println(err)
	}

	//transform to lines
	lines := strings.Split(string(content), "\n")
	//check if line exists
	if line < 1 || line > len(lines) {
		fmt.Println("Line doesnt exist")
	}

	//replace lines
	lines[line-1] = ""

	//rewrite the file
	newText := strings.Join(lines, "\n")

	os.WriteFile(FILE_NAME, []byte(newText), 0644)

	return true
}
