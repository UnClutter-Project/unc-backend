package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"unc/services/unc-service/application/client"
	"unc/services/unc-service/domain/repository"
	"unc/services/unc-service/domain/request"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
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

func (s *ClothingServiceImpl) CreateClothing(ctx context.Context, user_id string, req *request.CreateClothingRequest) error {
	var pgUUID pgtype.UUID
	err := pgUUID.Scan(user_id)
	if err != nil {
		return err
	}
	user, err := s.repository.GetUserById(ctx, pgUUID)

	if err != nil {
		return fmt.Errorf("invalid user: %w", err)
	}

	mainColor1, err := s.findOrCreateColor(ctx, req.MainColor1)
	if err != nil {
		return fmt.Errorf("failed to resolve main_color_1: %w", err)
	}

	var mainColor2ID pgtype.UUID
	if req.MainColor2 != "" {
		mainColor2, err := s.findOrCreateColor(ctx, req.MainColor2)
		if err != nil {
			return fmt.Errorf("failed to resolve main_color_2: %w", err)
		}
		mainColor2ID = mainColor2.ID
	}

	var accentColorID pgtype.UUID
	if req.AccentColor != "" {
		accentColor, err := s.findOrCreateColor(ctx, req.AccentColor)
		if err != nil {
			return fmt.Errorf("failed to resolve accent_color: %w", err)
		}
		accentColorID = accentColor.ID
	}

	category, err := s.findOrCreateCategory(ctx, user.ID, req.Category)
	if err != nil {
		return fmt.Errorf("failed to resolve category: %w", err)
	}

	imageLink, err := s.uploadImageToS3(ctx, req.Image)
	if err != nil {
		return fmt.Errorf("failed to upload image: %w", err)
	}
	fmt.Println(imageLink)

	_, err = s.repository.CreateClothing(ctx, &repository.CreateClothingParams{
		UserID:             user.ID,
		MainColor1ID:       mainColor1.ID,
		MainColor2ID:       mainColor2ID,
		AccentColorID:      accentColorID,
		ClothingCategoryID: category.ID,
		Brand:              pgtype.Text{String: req.Brand},
		Style:              pgtype.Text{String: req.Style},
		ImageLink:          pgtype.Text{String: imageLink, Valid: true},
	})
	if err != nil {
		return fmt.Errorf("failed to create clothing: %w", err)
	}

	return nil
}

func (s *ClothingServiceImpl) findOrCreateCategory(ctx context.Context, user_id pgtype.UUID, category string) (*repository.ClothingCategory, error) {
	exist_category, err := s.repository.GetClothingCategoryByValueAndUserID(ctx, &repository.GetClothingCategoryByValueAndUserIDParams{
		UserID: user_id,
		Value:  category,
	})
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	if err == nil {
		return exist_category, nil
	}

	new_category, err := s.repository.CreateClothingCategory(ctx, &repository.CreateClothingCategoryParams{
		UserID: user_id,
		Value:  category,
	})
	if err != nil {
		return nil, err
	}

	return new_category, nil
}

func (s *ClothingServiceImpl) findOrCreateColor(ctx context.Context, hex string) (*repository.Color, error) {
	exist_color, err := s.repository.GetColorByHex(ctx, hex)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	if err == nil {
		return exist_color, nil
	}

	// TODO: replace with go-colorful logic once ready
	colorGroupName := "black"

	newColor, err := s.repository.CreateColor(ctx, &repository.CreateColorParams{
		HexValue:       hex,
		ColorGroupName: colorGroupName,
	})
	if err != nil {
		return nil, err
	}

	return newColor, nil
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

	fileKey := fmt.Sprintf("clothing/%s%s", uuid.New().String(), filepath.Ext(file.Filename))

	_, err = s.storageClient.UploadFile(ctx, fileKey, fileData)
	if err != nil {
		return "", fmt.Errorf("failed to upload image to S3: %w", err)
	}
	imageLink, err := s.storageClient.DownloadFile(ctx, fileKey)
	// imageLink = config.GetConfig().BucketLink + "/" + imageLink

	return imageLink, nil
}
