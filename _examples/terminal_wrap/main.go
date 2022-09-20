package main

import (
	"fmt"

	promptercli "github.com/manifoldco/promptui"
)

func main() {
	longPrompt := promptercli.Prompt{
		Label: "This is a very long prompt that will wrap to the next line",
	}
	value, err := longPrompt.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Value:", value)
}
