package repositores

import (
	"fmt"

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

	fmt.Println("Repositories OK!")
	return user, nil
}
