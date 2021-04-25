package main

import (
	"fmt"
	"log"

	"github.com/lucymhdavies/go-fzfmaybe"
	"github.com/moby/moby/pkg/namesgenerator"
)

func main() {

	items := []string{}
	for i := 0; i <= 3; i++ {
		// Generate some random names, borrowing this handy Moby package
		// The example use case is picking from a list of AWS accounts in an Org
		name := namesgenerator.GetRandomName(0)

		// Suffix each of these with standard dev/test/etc.
		items = append(items, name+"/dev")
		items = append(items, name+"/test")
		items = append(items, name+"/staging")
		items = append(items, name+"/prod")
	}

	selected, err := fzfmaybe.Menu("Select Item From List", items)
	if err != nil {
		log.Fatalf("Could not get item from menu: %v", err)
	}

	fmt.Println("Selected account:", selected)
}
