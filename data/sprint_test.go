package data

import (
	"testing"

	"github.com/google/uuid"
)

func TestCreateSprint(t *testing.T) {
	test_user := NewUser("test_user_name", "test@email.com")
	s := NewSprint("test_sprint", test_user.id, 2)

	if s.name != "test_sprint" {
		t.Errorf("NewSprint name not set correctly")
	}
}

func TestAddBoard(t *testing.T) {
	test_user := NewUser("test_user_name", "test@email.com")
	s := NewSprint("test_sprint", test_user.id, 2)

	testID, _ := uuid.Parse("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")

	s.addBoard(testID)
	if s.boardRefs[0] != testID {
		t.Errorf("Board id not added correctly to sprint")
	}
	// assert s.boardRefs is equal t testID
}
