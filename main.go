package main

import (
	"bufio"
	"fmt"
	"os"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter scryfall query: ")
	text, _ := reader.ReadString('\n')

	criteria := Parse(text)

	fmt.Println(criteria)
}
