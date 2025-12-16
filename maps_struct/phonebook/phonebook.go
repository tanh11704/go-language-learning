package phonebook

import "sort"

type Contact struct {
	Name  string
	Phone string
}

type AddressBook struct {
	contacts map[string]Contact
}

func NewAddressBook() *AddressBook {
	return &AddressBook{
		contacts: make(map[string]Contact),
	}
}

func (a *AddressBook) AddContact(name, phone string) {
	a.contacts[name] = Contact{
		Name:  name,
		Phone: phone,
	}
}

func (a *AddressBook) FindContact(name string) (Contact, bool) {
	contact, exists := a.contacts[name]
	return contact, exists
}

func (a *AddressBook) DeleteContact(name string) {
	delete(a.contacts, name)
}

func (a *AddressBook) ListContacts() []Contact {
	var list []Contact
	for _, contact := range a.contacts {
		list = append(list, contact)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})

	return list
}
