package dal

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/mrForza/SaturnLMS/study_service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *CourseRepository) GetAllCourses() ([]models.Course, error) {
	var courses []models.Course
	cursor, err := r.courseCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var course models.Course
		if err := cursor.Decode(&course); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *CourseRepository) GetCourseById(id string) (*models.Course, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid course id format")
	}

	var course models.Course
	err = r.courseCollection.FindOne(context.Background(), bson.M{"id": uuidID}).Decode(&course)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("course not found")
		}
		return nil, err
	}

	return &course, nil
}

func (r *CourseRepository) CreateCourse(id uuid.UUID, dto models.Course) (uuid.UUID, error) {
	dto.Id = id
	_, err := r.courseCollection.InsertOne(context.Background(), dto)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (r *CourseRepository) DeleteCourseById(id string) error {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid course id format")
	}

	result, err := r.courseCollection.DeleteOne(context.Background(), bson.M{"id": uuidID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("no course was deleted")
	}

	return nil
}

func (r *CourseRepository) UpdateCourseById(id uuid.UUID, updates map[string]interface{}) error {
	return nil
}
