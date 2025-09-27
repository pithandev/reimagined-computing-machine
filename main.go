// package declaration
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scan := bufio.NewScanner(os.Stdin)

	quit := false

	for !quit {

		fmt.Println()
		fmt.Println("TRY SAVE SOMETHING IN TXT FILE.")
		fmt.Println("1) Save.")
		fmt.Println("2) Show data.")
		fmt.Println("0) QUIT!!!")
		fmt.Println()

		scan.Scan()
		opt := scan.Text()

		switch opt {
		case "1":

			textScan := bufio.NewScanner(os.Stdin)
			textScan.Scan()

			text := textScan.Text()

			save(text)
		case "2":

		case "0":
			quit = true
		}

	}

}

func save(data string) string {

	file, err := os.OpenFile("storage.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)

	}

	if _, err := file.Write([]byte(data)); err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	return "Data stored!"
}
