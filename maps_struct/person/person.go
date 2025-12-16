package person

import (
	"fmt"
	"strings"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func (p Person) IsValidEmail() bool {
	if !strings.Contains(p.Email, "@") || !strings.Contains(p.Email, ".") {
		return false
	}
	return true
}

func (p Person) IsValidAge() bool {
	return p.Age >= 0 && p.Age <= 150
}

func (p Person) PrintInfo() string {
	return fmt.Sprintf(`
------------------------------------
👤 USER PROFILE
------------------------------------
Name  : %s
Age   : %d
Email : %s
------------------------------------`, p.Name, p.Age, p.Email)
}
