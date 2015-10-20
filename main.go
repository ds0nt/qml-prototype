package main

import (
	"fmt"
	"os"
	"os/exec"

	"gopkg.in/qml.v1"
)

func main() {
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

}

func run() error {
	engine := qml.NewEngine()

	controls, err := engine.LoadFile("main.qml")
	if err != nil {
		return err
	}

	window := controls.CreateWindow(nil)
	window.On("visibleChanged", func(visible bool) {
		if visible {
			fmt.Println("Width:", window.Int("width"))
			fmt.Println("Height:", window.Int("height"))
		}
	})

	submit := window.ObjectByName("submit")
	query := window.ObjectByName("query")
	results := window.ObjectByName("results")

	submit.On("clicked", func() {
		go func() {
			cmd := "/home/dsont/dot-files/commands/search-aur.sh"
			querystring := fmt.Sprintf("%s", query.Property("text"))
			result, err := exec.Command(cmd, querystring).Output()
			if err != nil {
				results.Set("text", fmt.Sprintf("Errors be appeared all over: %v", err))
			}
			results.Set("text", string(result))
		}()
	})

	window.Show()
	window.Wait()
	return nil
}
