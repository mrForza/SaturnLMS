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

func GetAllFacultaties() ([]models.Facultaty, error) {
	var facultaties []models.Facultaty
	query := "SELECT name, description, university_name FROM facultaty;"

	err := Db.Select(&facultaties, query)
	if err != nil {
		return nil, err
	}

	return facultaties, nil
}

func GetFacultatyByName(name string) (*models.Facultaty, error) {
	var facultaty models.Facultaty

	query := `
        SELECT name, description, university_name
        FROM facultaty
        WHERE name = $1;
    `

	err := Db.Get(&facultaty, query, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with name %s not found", name)
		}
		return nil, fmt.Errorf("failed to fetch facultaty: %w", err)
	}

	return &facultaty, nil
}

func CreateFacultaty(facultaty dtos.CreateFacultatyRequestDto) (string, error) {
	query := "INSERT INTO facultaty (name, description, university_name) VALUES ($1, $2, $3);"

	_, err := Db.Exec(
		query,
		facultaty.Name,
		facultaty.Description,
		facultaty.UniversityName,
	)

	if err != nil {
		log.Fatalf("%s", err.Error())
		return "", errors.New("the facultaty cannot be created because it has already existed")
	}

	return facultaty.Name, nil
}

func DeleteFacultatyByName(name string) error {
	query := `
        DELETE FROM facultaty
        WHERE name = $1;
    `

	result, err := Db.Exec(query, name)
	if err != nil {
		return fmt.Errorf("failed to delete facultaty: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %s not found", name)
	}

	fmt.Printf("facultaty with ID %s successfully deleted\n", name)
	return nil
}

func UpdateFacultatyById(id uuid.UUID, updates map[string]interface{}) error {
	query := "UPDATE facultaty SET "
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
		return fmt.Errorf("failed to update facultaty: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %s not found", id)
	}

	fmt.Printf("facultaty with ID %s successfully updated\n", id)
	return nil
}
