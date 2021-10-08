package fzfmaybe

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

func Menu(title string, items []string) (string, error) {
	// Does FZF exist?
	_, err := exec.LookPath("fzf")

	// Nope. Fallback to tui.Menu
	if err != nil {
		searcher := func(input string, index int) bool {
			item := items[index]
			name := strings.Replace(strings.ToLower(item), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)

			return strings.Contains(name, input)
		}
		prompt := promptui.Select{
			Label:    title,
			Items:    items,
			Size:     len(items),
			Searcher: searcher,
		}

		_, result, err := prompt.Run()

		return result, err
	}

	return fzfMenu(title, items)
}

func fzfMenu(title string, items []string) (string, error) {

	app := "fzf"

	args := []string{"--header", title, "--reverse"}

	cmd := exec.Command(app, args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, strings.Join(items, "\n"))
	}()
	cmd.Stderr = os.Stderr
	var outb bytes.Buffer
	cmd.Stdout = &outb

	err = cmd.Run()

	if err != nil {
		return "", err
	}

	selected := strings.TrimSpace(outb.String())

	return selected, nil
}
