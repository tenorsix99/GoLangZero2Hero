// ✅ Part 2: Custom Error ด้วย Struct

package main

import (
	"fmt"
)

type NotFoundError struct {
	Resource string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("🔍 %s not found", e.Resource)
}

func findUser(id string) error {
	if id == "" {
		return &NotFoundError{Resource: "User...."}
	}
	return nil
}

func main() {
	err := findUser("")
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("Success found user in Database")
	}
}
