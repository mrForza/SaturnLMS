package dal

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/mrForza/SaturnLMS/study_service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *LessonRepository) GetAllLessons() ([]models.Lesson, error) {
	var Lessons []models.Lesson
	cursor, err := r.lessonCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var Lesson models.Lesson
		if err := cursor.Decode(&Lesson); err != nil {
			return nil, err
		}
		Lessons = append(Lessons, Lesson)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return Lessons, nil
}

func (r *LessonRepository) GetLessonById(id string) (*models.Lesson, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid Lesson id format")
	}

	var Lesson models.Lesson
	err = r.lessonCollection.FindOne(context.Background(), bson.M{"id": uuidID}).Decode(&Lesson)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("Lesson not found")
		}
		return nil, err
	}

	return &Lesson, nil
}

func (r *LessonRepository) CreateLesson(id uuid.UUID, dto models.Lesson) (uuid.UUID, error) {
	dto.Id = id
	_, err := r.lessonCollection.InsertOne(context.Background(), dto)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (r *LessonRepository) DeleteLessonById(id string) error {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid Lesson id format")
	}

	result, err := r.lessonCollection.DeleteOne(context.Background(), bson.M{"id": uuidID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("no Lesson was deleted")
	}

	return nil
}

func (r *LessonRepository) UpdateLessonById(id uuid.UUID, updates map[string]interface{}) error {
	return nil
}

func CreateLessonBucketByUUID(client *minio.Client, bucketUUID uuid.UUID) error {
	ctx := context.Background()
	bucketName := bucketUUID.String()

	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("bucket %s has already existed", bucketName)
	}

	err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		return err
	}

	return nil
}

func UploadFilesToLessonBucket(client *minio.Client, bucketUUID uuid.UUID, files []*multipart.FileHeader) ([]string, error) {
	ctx := context.Background()
	bucketName := bucketUUID.String()

	var uploadedFileNames []string

	for _, fileHeader := range files {
		fileReader, err := fileHeader.Open()
		if err != nil {
			return nil, fmt.Errorf("cannot open a file: %v", err)
		}
		defer fileReader.Close()

		objectName := fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(fileHeader.Filename))

		_, err = client.PutObject(ctx, bucketName, objectName, fileReader, fileHeader.Size, minio.PutObjectOptions{})
		if err != nil {
			return nil, fmt.Errorf("error while upload a file %s: %v", fileHeader.Filename, err)
		}

		uploadedFileNames = append(uploadedFileNames, objectName)
	}

	return uploadedFileNames, nil
}

func DeleteBucketAndAllFiles(client *minio.Client, bucketName string) error {
	ctx := context.Background()

	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("error while checking file existance: %v", err)
	}
	if !exists {
		return fmt.Errorf("file with id: %s does not exist", bucketName)
	}

	objectCh := make(chan string)

	go func() {
		defer close(objectCh)
		for object := range client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{Recursive: true}) {
			if object.Err != nil {
				fmt.Printf("Error while getting the file: %v\n", object.Err)
				continue
			}
			objectCh <- object.Key
		}
	}()

	for objectKey := range objectCh {
		err := client.RemoveObject(ctx, bucketName, objectKey, minio.RemoveObjectOptions{})
		if err != nil {
			return fmt.Errorf("erorr while deleting the file %s: %v", objectKey, err)
		}
		fmt.Printf("File has been deleted: %s\n", objectKey)
	}

	err = client.RemoveBucket(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("erorr while deleting the file: %v", err)
	}

	fmt.Printf("File with id: %s all it's content were successfully removed\n", bucketName)
	return nil
}

func UploadFileToLesson(client *minio.Client, bucketName string, fileHeader *multipart.FileHeader) error {
	// Открываем временный файл из заголовка
	src, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer src.Close()

	objectName := fileHeader.Filename

	_, err = client.PutObject(context.Background(), bucketName, objectName, src, fileHeader.Size, minio.PutObjectOptions{
		ContentType: fileHeader.Header.Get("Content-Type"),
	})
	if err != nil {
		return fmt.Errorf("ошибка при загрузке файла в MinIO: %v", err)
	}

	fmt.Printf("Файл '%s' успешно загружен в бакет '%s'\n", objectName, bucketName)
	return nil
}

func DeleteFileFromMinIO(client *minio.Client, bucketName string, objectName string) error {
	err := client.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("ошибка при удалении файла '%s' из бакета '%s': %v", objectName, bucketName, err)
	}

	fmt.Printf("Файл '%s' успешно удален из бакета '%s'\n", objectName, bucketName)
	return nil
}

func DownloadFileFromMinIO(client *minio.Client, bucketName string, objectName string) ([]byte, error) {
	reader, err := client.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении файла '%s' из бакета '%s': %v", objectName, bucketName, err)
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении содержимого файла '%s': %v", objectName, err)
	}

	fmt.Printf("Файл '%s' успешно загружен из бакета '%s'\n", objectName, bucketName)
	return data, nil
}
