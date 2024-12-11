package repositores

import (
	"fmt"
	"strings"

	"github.com/GrubbyRM/CyberGo/cmd/models"
	"github.com/GrubbyRM/CyberGo/cmd/storage"
)

func CreateUser(user models.User) (models.User, error) {
	db := storage.GetDB()

	query := `INSERT INTO users_table (name, password, role) VALUES (?, ?, ?)` // Używamy "?" jako placeholderów w MySQL
	_, err := db.Exec(query, user.Name, user.Password, user.Role)
	if err != nil {
		return user, err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&user.Id)
	if err != nil {
		return user, err
	}

	fmt.Println("User created")
	return user, nil
}

func DeleteUser(userName string) error {
	db := storage.GetDB()

	query := `DELETE FROM users_table WHERE name=?`
	_, err := db.Exec(query, userName)
	if err != nil {
		return err
	}
	fmt.Println("User removed")
	return nil
}

func ShowUsers() ([]string, error) {
	db := storage.GetDB()

	query := `SELECT name FROM users_table`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	var users []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		users = append(users, name)
	}

	return users, nil
}

func UpdatePassword(userName string, newPassword string) error {
	db := storage.GetDB()

	query := `UPDATE users_table 
			  SET password = ?
			  WHERE user = ?`

	_, err := db.Exec(query, userName, newPassword)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(userName string, userData map[string]interface{}) error {
	db := storage.GetDB()

	setClauses := []string{}
	args := []interface{}{}

	for k, v := range userData {
		setClauses = append(setClauses, fmt.Sprintf("%s = ?", k))
		args = append(args, v)
	}

	args = append(args, userName)
	query := fmt.Sprintf("UPDATE users_table SET %s WHERE name = ?", strings.Join(setClauses, ", "))

	_, err := db.Exec(query, args...)

	if err != nil {
		return fmt.Errorf("Database error: %w", err)
	}

	return nil

}
