// package declaration
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scan := bufio.NewScanner(os.Stdin)

	quit := false

	for !quit {

		fmt.Println()
		fmt.Println("TRY SAVE SOMETHING IN TXT FILE.")
		fmt.Println("1) Save.")
		fmt.Println("2) Show.")
		fmt.Println("3) Update.")
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

			text := textScan.Text()

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
			newText := textScan.Text()

			update(int(id), newText)
			fmt.Println("==============================")

		case "0":
			quit = true
		}

	}

}

func save(data string) string {

	writeFile, err := os.OpenFile("storage.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

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
	readFile, err := os.ReadFile("storage.txt")

	if err != nil {
		fmt.Println(err)
	}

	return string(readFile)
}

func update(id int, newText string) bool {

	updateFile, err := os.OpenFile("storage.txt", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return false
	}

	updateFile.Seek(int64(id), 0)

	updateFile.Write([]byte(newText))

	defer updateFile.Close()

	return true
}
