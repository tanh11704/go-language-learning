package person

import (
	"fmt"
	"testing"
)

func TestPersonMethods(t *testing.T) {
	p := Person{
		Name:  "Nguyen Van A",
		Age:   25,
		Email: "test@gmail.com",
	}

	t.Run("Check Valid Email", func(t *testing.T) {
		if !p.IsValidEmail() {
			t.Error("Expected email to be valid")
		}

		badPerson := Person{Email: "invalid-email"}
		if badPerson.IsValidEmail() {
			t.Error("Expected email to be invalid")
		}
	})

	t.Run("Check Valid Age", func(t *testing.T) {
		if !p.IsValidAge() {
			t.Error("Expected age 25 to be valid")
		}

		oldPerson := Person{Age: 200}
		if oldPerson.IsValidAge() {
			t.Error("Expected age 200 to be invalid")
		}
	})

	t.Run("Check Print Format", func(t *testing.T) {
		info := p.PrintInfo()
		fmt.Println(info)
		if info == "" {
			t.Error("Expected info string, got empty")
		}
	})
}
