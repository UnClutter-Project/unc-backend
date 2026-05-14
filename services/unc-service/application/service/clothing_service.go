package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"unc/services/unc-service/application/client"
	"unc/services/unc-service/config"
	"unc/services/unc-service/domain/repository"
	"unc/services/unc-service/domain/request"

	"github.com/google/uuid"
)

type ClothingService interface {
	CreateClothing(ctx context.Context, user_id string, request *request.CreateClothingRequest) error
}

type ClothingServiceImpl struct {
	repository    repository.Querier
	storageClient client.StorageClient
}

func NewClothingService(repository repository.Querier, storageClient client.StorageClient) ClothingService {
	return &ClothingServiceImpl{
		repository:    repository,
		storageClient: storageClient,
	}
}

func (s *ClothingServiceImpl) CreateClothing(ctx context.Context, userID string, req *request.CreateClothingRequest) error {
	// 1. Upload image to S3
	imageLink, err := s.uploadImageToS3(ctx, req.Image)
	if err != nil {
		return fmt.Errorf("failed to upload image: %w", err)
	}
	fmt.Printf("link = %s", imageLink)
	return nil
}

func (s *ClothingServiceImpl) uploadImageToS3(ctx context.Context, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open image: %w", err)
	}
	defer src.Close()

	fileData, err := io.ReadAll(src)
	if err != nil {
		return "", fmt.Errorf("failed to read image: %w", err)
	}

	fileKey := fmt.Sprintf("%s/clothing/%s%s", config.GetConfig().BucketLink, uuid.New().String(), filepath.Ext(file.Filename))

	imageLink, err := s.storageClient.UploadFile(ctx, fileKey, fileData)
	if err != nil {
		return "", fmt.Errorf("failed to upload image to S3: %w", err)
	}

	return imageLink, nil
}
