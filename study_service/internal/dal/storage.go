package dal

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToMongo(uri string) (*mongo.Client, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	fmt.Println("Connected to MongoDB")
	return client, cancel
}

const (
	minioEndpoint  = "studyService.minio:9000"
	minioAccessKey = "admin"
	minioSecretKey = "password123"
	bucketName     = "uploads"
	location       = "us-east-1"
)

func initMinioClient() (*minio.Client, error) {
	client, err := minio.New(minioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioAccessKey, minioSecretKey, ""),
		Secure: false, // поставь true, если используешь HTTPS
	})
	if err != nil {
		return nil, err
	}

	// Создаем бакет, если его нет
	ctx := context.Background()
	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		return nil, err
	}
	if !exists {
		err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			return nil, err
		}
	}

	return client, nil
}

type CourseRepository struct {
	courseCollection *mongo.Collection
}

func NewCourseRepository(db *mongo.Database) *CourseRepository {
	return &CourseRepository{
		courseCollection: db.Collection("courses"),
	}
}

type LessonRepository struct {
	lessonCollection *mongo.Collection
}

func NewLessonRepository(db *mongo.Database) *LessonRepository {
	return &LessonRepository{
		lessonCollection: db.Collection("lessons"),
	}
}

var Client, Cancel = ConnectToMongo("mongodb://admin:secret@studyService.mongo:27017")

var MinioClient, _ = initMinioClient()
