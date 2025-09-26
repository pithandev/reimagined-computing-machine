// package declaration
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("TRY SAVE SOMETHING IN TXT FILE.")
		fmt.Println("type \"0\" to exit the program.")
		fmt.Println()

		reader.Scan()

		opt := reader.Text()

		switch opt {
		case "1":
			fmt.Println("Type what u wanna save: ")
			scan := bufio.NewReader(os.Stdin)
			text, err := scan.ReadString('\n')

			if err != nil {
				fmt.Print(err)
			}

			save(text)
		case "0":
			return
		}

	}

}

func save(data string) string {

	err := os.WriteFile("storage.txt", []byte(data), 0644)

	if err != nil {
		fmt.Println(err)

	}

	return "Data stored!"
}
