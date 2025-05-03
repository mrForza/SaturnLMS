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

func GetAllPrograms() ([]models.Program, error) {
	var programs []models.Program
	query := "SELECT name, description, type, languages, facultaty_name FROM program;"

	err := Db.Select(&programs, query)
	if err != nil {
		return nil, err
	}

	return programs, nil
}

func GetProgramByName(name string) (*models.Program, error) {
	var program models.Program

	query := `
        SELECT name, description, type, languages, facultaty_name
        FROM program
        WHERE name = $1;
    `

	err := Db.Get(&program, query, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("program with name %s not found", name)
		}
		return nil, fmt.Errorf("failed to fetch program: %w", err)
	}

	return &program, nil
}

func CreateProgram(program dtos.CreateProgramRequestDto) (string, error) {
	query := "INSERT INTO program (name, description, type, languages, facultaty_name) VALUES ($1, $2, $3, $4, $5);"

	_, err := Db.Exec(
		query,
		program.Name,
		program.Description,
		program.Type,
		program.Languages,
		program.FacultatyNme,
	)

	if err != nil {
		log.Fatalf("%s", err.Error())
		return "", errors.New("the program cannot be created because it has already existed")
	}

	return program.Name, nil
}

func DeleteProgramByName(name string) error {
	query := `
        DELETE FROM program
        WHERE name = $1;
    `

	result, err := Db.Exec(query, name)
	if err != nil {
		return fmt.Errorf("failed to delete program: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("program with name %s not found", name)
	}

	fmt.Printf("program with ID %s successfully deleted\n", name)
	return nil
}

func UpdateProgramById(id uuid.UUID, updates map[string]interface{}) error {
	query := "UPDATE program SET "
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
		return fmt.Errorf("failed to update program: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %s not found", id)
	}

	fmt.Printf("program with ID %s successfully updated\n", id)
	return nil
}
