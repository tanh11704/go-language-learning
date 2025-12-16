package phonebook

import (
	"testing"
)

func TestAddressBook(t *testing.T) {
	book := NewAddressBook()

	book.AddContact("Alice", "111-111")
	book.AddContact("Charlie", "333-333")
	book.AddContact("Bob", "222-222")

	t.Run("Find Existing Contact", func(t *testing.T) {
		contact, found := book.FindContact("Alice")
		if !found || contact.Phone != "111-111" {
			t.Errorf("Expected to find Alice with phone 111-111")
		}
	})

	t.Run("Find Non-existent Contact", func(t *testing.T) {
		_, found := book.FindContact("David")
		if found {
			t.Error("Expected not to find David")
		}
	})

	t.Run("List Sorted Contacts", func(t *testing.T) {
		list := book.ListContacts()

		if len(list) != 3 {
			t.Errorf("Expected 3 contacts, got %d", len(list))
		}

		expectedOrder := []string{"Alice", "Bob", "Charlie"}
		for i, contact := range list {
			if contact.Name != expectedOrder[i] {
				t.Errorf("Index %d: expected %s, got %s", i, expectedOrder[i], contact.Name)
			}
		}
	})

	t.Run("Delete Contact", func(t *testing.T) {
		book.DeleteContact("Bob")
		_, found := book.FindContact("Bob")
		if found {
			t.Error("Bob should have been deleted")
		}
	})
}
