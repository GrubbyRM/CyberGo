package services

import (
	"fmt"

	"github.com/GrubbyRM/CyberGo/cmd/repositores"
)

func CheckUser(userName string) (bool, error) {
	if userName == "" {
		return false, fmt.Errorf("user name is required")
	}

	users, err := repositores.ShowUsers()

	if err != nil {
		return false, err
	}

	for _, user := range users {
		if user == userName {
			return true, nil
		}
	}
	return false, nil
}
func DeleteUser(userName string) error {

	exists, err := CheckUser(userName)
	if err != nil {
		return fmt.Errorf("error occured: %w", err)
	}
	if !exists {
		return fmt.Errorf("user does not exist")
	}

	err = repositores.DeleteUser(userName)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

// TODO: add form
func UpdateUser(userName string, userData map[string]interface{}) error {
	exists, err := CheckUser(userName)
	if err != nil {
		return fmt.Errorf("error occured: %w", err)
	}
	if !exists {
		return fmt.Errorf("user not found")
	}
	err = repositores.UpdateUser(userName, userData)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil

}
