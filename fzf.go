package fzfmaybe

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/johnnylee/tui"
)

func Menu(title string, items []string) (string, error) {
	// Does FZF exist?
	_, err := exec.LookPath("fzf")

	// Nope. Fallback to tui.Menu
	if err != nil {
		var menuItems []string
		for i, j := range items {
			menuItems = append(menuItems, []string{strconv.Itoa(i), j}...)
		}
		x := tui.Menu(title, nil, menuItems...)

		if i, err := strconv.Atoi(x); err == nil {
			return items[i], nil
		}
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

	return outb.String(), nil
}
