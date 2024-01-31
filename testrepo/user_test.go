package test

import (
	"ks/domain"
	"ks/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetByName(t *testing.T) {
	testUser := domain.SignUp{
		UserName: "preyas",
		Password: "1234",
	}
	err := repository.GetByName(testUser)
	if err != nil {
		if err.Error() != "username already exists" {
			t.Errorf("expcted error 'username already exists',but got '%s'", err.Error())
		} else {
			t.Errorf("expected error got 'username already exists',got no error ")
		}
	}
}
func TestCreateUser(t *testing.T) {
	testUser := domain.SignUp{
		UserName: "testuser",
		Password: "testpassword",
	}

	// Create user
	err := repository.CreateUser(testUser)
	defer func() {
		// Clean up: Delete user even if assertions fail
		deleteErr := repository.DeleteUser(testUser)
		assert.NoError(t, deleteErr, "Error deleting user: %v", deleteErr)
	}()

	// Assertion 1: Check if user creation was successful
	assert.NoError(t, err, "Error creating user: %v", err)

	// Assertion 2: Check if user deletion was successful
	deleteErr := repository.DeleteUser(testUser)
	assert.NoError(t, deleteErr, "Error deleting user: %v", deleteErr)

	// Assertion 3: Check if user retrieval was successful
	getErr := repository.GetByName(testUser)
	assert.NoError(t, getErr, "Error getting user: %v", getErr)
}
func TestLogin(t *testing.T)  {
	testuser:=domain.SignUp{
		UserName: "preya",
		Password: "345",
	}
	err:=repository.Login(testuser)
	assert.Error(t,err,"expected Error")
	assert.EqualError(t,err,"invalid username or password", "Expected error message 'invalid username or password'")
}