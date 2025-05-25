package routers

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/mrForza/SaturnLMS/study_service/internal/dal"
	"github.com/mrForza/SaturnLMS/study_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/study_service/internal/usecases"
)

func RegisterLessonRouter(baseRouter *gin.Engine) {
	LessonRouters := baseRouter.Group("/lessons")
	{
		LessonRouters.GET("/:id", GetLessonById)
		LessonRouters.POST("/", CreateLessonRoute)
		LessonRouters.PATCH("/:id", UpdateLessonById)
		LessonRouters.DELETE("/:id", DeleteLessonRoute)
		LessonRouters.POST("/:id/upload_file", UploadFileInLesson)
		LessonRouters.DELETE("/:id/delete_file", DeleteFileFromLesson)
		LessonRouters.GET("/:id/download_file/:file_name", DownloadFileFromLesson)
	}
}

func GetLessonById(ctx *gin.Context) {
	bucketName := ctx.Param("id")

	exists, err := dal.MinioClient.BucketExists(context.Background(), bucketName)
	if err != nil || !exists {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Бакет не найден"})
		return
	}

	tempFile, err := os.CreateTemp("", fmt.Sprintf("%s_*.zip", bucketName))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать временный архив"})
		return
	}
	defer os.Remove(tempFile.Name())

	zipWriter := zip.NewWriter(tempFile)
	defer zipWriter.Close()

	for object := range dal.MinioClient.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{Recursive: true}) {
		if object.Err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка при чтении файла %s", object.Key)})
			return
		}

		objReader, err := dal.MinioClient.GetObject(context.Background(), bucketName, object.Key, minio.GetObjectOptions{})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Не удалось получить файл %s", object.Key)})
			return
		}
		defer objReader.Close()

		header := &zip.FileHeader{
			Name:   object.Key,
			Method: zip.Deflate,
		}
		dstWriter, err := zipWriter.CreateHeader(header)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении файла в архив"})
			return
		}

		if _, err := io.Copy(dstWriter, objReader); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при копировании данных в архив"})
			return
		}
	}

	if err := zipWriter.Close(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при завершении архива"})
		return
	}

	ctx.Header("Content-Type", "application/zip")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s.zip", bucketName))
	ctx.Header("Content-Length", fmt.Sprintf("%d", getFileSize(tempFile)))

	http.ServeFile(ctx.Writer, ctx.Request, tempFile.Name())
}

func getFileSize(file *os.File) int64 {
	stat, _ := file.Stat()
	return stat.Size()
}

func CreateLessonRoute(ctx *gin.Context) {
	name := ctx.PostForm("name")
	description := ctx.PostForm("description")
	lessonType := ctx.PostForm("type") == "true"
	// homework := ctx.PostForm("homework")

	form, _ := ctx.MultipartForm()
	files := form.File["files"]

	dto := dtos.CreateLessonRequestDto{
		Name:        name,
		Description: description,
		Type:        lessonType,
		Files:       files,
	}

	ctx.JSON(http.StatusCreated, usecases.CreateLesson(dto))
}

func UpdateLessonById(ctx *gin.Context) {

}

func DeleteLessonRoute(ctx *gin.Context) {
	id := ctx.Param("id")

	err := usecases.DeleteLesson(dtos.DeleteLessonByIdRequest{Id: id})
	if err != nil {
		ctx.JSON(http.StatusNoContent, gin.H{"message": err.Error()})
		return
	}

	err2 := dal.DeleteBucketAndAllFiles(dal.MinioClient, id)
	if err2 != nil {
		ctx.JSON(http.StatusNoContent, gin.H{"message": err2.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func UploadFileInLesson(ctx *gin.Context) {
	id := ctx.Param("id")
	form, _ := ctx.MultipartForm()
	files := form.File["files"]

	uuidID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"messsage": "Incorrect lesson_id format"})
	}

	dto := usecases.UploadFileIntoLessonById(dtos.UploadFileRequestDto{
		BucketId: uuidID,
		Files:    files,
	})

	if dto.IsError == true {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": dto.Message})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"messsage": dto.Message})
}

func DeleteFileFromLesson(ctx *gin.Context) {
	id := ctx.Param("id")
	fileName := ctx.PostForm("file_name")

	uuidID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"messsage": "Incorrect lesson_id format"})
	}

	dto := usecases.DeleteFileFromLessonById(dtos.DeleteFileRequestDto{
		BucketId: uuidID,
		FileName: fileName,
	})

	if dto.IsError == true {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": dto.Message})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"messsage": dto.Message})
}

func DownloadFileFromLesson(ctx *gin.Context) {
	bucketId := ctx.Param("id")
	fileName := ctx.Param("file_name")

	uuidID, err := uuid.Parse(bucketId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"messsage": "Incorrect lesson_id format"})
	}

	dto := usecases.DownloadFileFromLessonById(dtos.DownloadFileRequestDto{
		BucketId: uuidID,
		FileName: fileName,
	})

	if dto.IsError {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "No such file"})
		return
	}

	ctx.Header("Content-Type", "application/zip")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s.zip", fileName))
	// ctx.Header("Content-Length", fmt.Sprintf("%d", getFileSize(*dto.File)))

	ctx.Data(http.StatusOK, "application/octet-stream", *dto.File)
}
