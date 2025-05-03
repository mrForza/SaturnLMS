package dal

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/mrForza/SaturnLMS/organization_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/organization_service/internal/models"
)

func GetAllUniversities() ([]models.University, error) {
	var universities []models.University
	query := "SELECT name, description, legal_address, actual_address, inn, bank_name, owner_id FROM university;"

	err := Db.Select(&universities, query)
	if err != nil {
		return nil, err
	}

	return universities, nil
}

func GetUniversityByName(name string) (*models.University, error) {
	var university models.University

	query := `
        SELECT name, description, legal_address, actual_address, inn, bank_name, owner_id
        FROM university
        WHERE name = $1;
    `

	err := Db.Get(&university, query, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with name %s not found", name)
		}
		return nil, fmt.Errorf("failed to fetch university: %w", err)
	}

	return &university, nil
}

func CreateUniversity(university dtos.CreateUniversityRequestDto) (string, error) {
	query := "INSERT INTO university (name, description, legal_address, actual_address, inn, bank_name, owner_id) VALUES ($1, $2, $3, $4, $5, $6, $7);"

	_, err := Db.Exec(
		query,
		university.Name,
		university.Description,
		university.LegalAddress,
		university.ActualAddress,
		university.Inn,
		university.BankName,
		university.OwnerId,
	)

	if err != nil {
		log.Fatalf("%s", err.Error())
		return "", errors.New("the University cannot be created because this user has already existed")
	}

	return university.Name, nil
}

func DeleteUniversityByName(name string) error {
	query := `
        DELETE FROM university
        WHERE name = $1;
    `

	result, err := Db.Exec(query, name)
	if err != nil {
		return fmt.Errorf("failed to delete university: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %s not found", name)
	}

	fmt.Printf("User with ID %s successfully deleted\n", name)
	return nil
}

func UpdateUniversityById(id uuid.UUID, updates map[string]interface{}) error {
	query := "UPDATE university SET "
	var args []interface{}
	i := 1

	for key, value := range updates {
		query += fmt.Sprintf("%s = $%d, ", key, i)
		args = append(args, value)
		i++
	}

	query = query[:len(query)-2] + " WHERE id = $" + fmt.Sprint(i)
	args = append(args, id)

	result, err := Db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to update university: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %s not found", id)
	}

	fmt.Printf("User with ID %s successfully updated\n", id)
	return nil
}
