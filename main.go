package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter scryfall query: ")
	text, _ := reader.ReadString('\n')

	criteria := Parse(text)
	jcriteria, _ := json.MarshalIndent(criteria, "", "\t")
	fmt.Printf("%+v\n", string(jcriteria))
}
