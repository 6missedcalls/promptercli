package main

import (
	"fmt"

	promptercli "github.com/manifoldco/promptui"
)

func main() {
	prompt := promptercli.Prompt{
		Label:     "Delete Resource",
		IsConfirm: true,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}
