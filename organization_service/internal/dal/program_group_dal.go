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

func GetAllProgramGroups() ([]models.ProgramGroup, error) {
	var programGroups []models.ProgramGroup
	query := "SELECT number, name, course_number, program_name FROM program_group;"

	err := Db.Select(&programGroups, query)
	if err != nil {
		return nil, err
	}

	return programGroups, nil
}

func GetProgramGroupByNumber(number uint16) (*models.ProgramGroup, error) {
	var programGroup models.ProgramGroup

	query := `
        SELECT number, name, course_number, program_name
        FROM program_group
        WHERE number = $1;
    `

	err := Db.Get(&programGroup, query, number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("programGroup with number %d not found", number)
		}
		return nil, fmt.Errorf("failed to fetch programGroup: %w", err)
	}

	return &programGroup, nil
}

func CreateProgramGroup(programGroup dtos.CreateProgramGroupRequestDto) (string, error) {
	query := "INSERT INTO program_group (number, name, course_number, program_name) VALUES ($1, $2, $3, $4);"

	_, err := Db.Exec(
		query,
		programGroup.Number,
		programGroup.Name,
		programGroup.CourseNumber,
		programGroup.ProgramName,
	)

	if err != nil {
		log.Fatalf("%s", err.Error())
		return "", errors.New("the programGroup cannot be created because it has already existed")
	}

	return programGroup.Name, nil
}

func DeleteProgramGroupByName(number uint16) error {
	query := `
        DELETE FROM program_group
        WHERE number = $1;
    `

	result, err := Db.Exec(query, number)
	if err != nil {
		return fmt.Errorf("failed to delete programGroup: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("programGroup with number %d not found", number)
	}

	fmt.Printf("programGroup with number %d successfully deleted\n", number)
	return nil
}

func UpdateProgramGroupById(id uuid.UUID, updates map[string]interface{}) error {
	query := "UPDATE program_group SET "
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
		return fmt.Errorf("failed to update programGroup: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %s not found", id)
	}

	fmt.Printf("programGroup with ID %s successfully updated\n", id)
	return nil
}
