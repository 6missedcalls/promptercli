package main

import (
	"errors"
	"fmt"

	promptercli "github.com/manifoldco/promptui"
)

func main() {
	validate := func(input string) error {
		if len(input) < 6 {
			return errors.New("Password must have more than 6 characters")
		}
		return nil
	}

	prompt := promptercli.Prompt{
		Label:    "Password",
		Validate: validate,
		Mask:     '*',
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("Your password is %q\n", result)
}
