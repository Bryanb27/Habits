package main

import (
	"testing"
)

// Test case for creating a user
func TestCreateUser(t *testing.T) {
	user, err := createUser()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	//Check if the user is created with the correct data
	if user.name == "" || user.age == 0 || user.email == "" || user.password == "" {
		t.Errorf("User creation failed, some fields are empty")
	}
}

func TestCreateNewHabit(t *testing.T) {

	habit, err := createNewHabit()

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if habit == nil {
		t.Errorf("Expected a habit, but got none")
	}
}
