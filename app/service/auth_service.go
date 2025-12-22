package service

import (
	"context"
	"errors"
	"simak-go/app/config"
	"simak-go/app/helper"
	"simak-go/app/models"
	"simak-go/app/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthService interface {
	Register(ctx context.Context, req models.RegisterRequest) (*models.User, error)
	Login(ctx context.Context, req models.LoginRequest) (string, *models.User, error)
}

type authService struct {
	userRepo repository.UserRepository
	config   *config.Config
}

func NewAuthService(userRepo repository.UserRepository, cfg *config.Config) AuthService {
	return &authService{
		userRepo: userRepo,
		config:   cfg,
	}
}

func (s *authService) Register(ctx context.Context, req models.RegisterRequest) (*models.User, error) {
	// Check if email exists
	existingUser, _ := s.userRepo.FindByEmail(ctx, req.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		ID:       primitive.NewObjectID(),
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     req.Role,
	}

	err = s.userRepo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *authService) Login(ctx context.Context, req models.LoginRequest) (string, *models.User, error) {
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	if !helper.CheckPassword(req.Password, user.Password) {
		return "", nil, errors.New("invalid email or password")
	}

	token, err := helper.GenerateToken(user.ID.Hex(), user.Role, s.config.JWTSecret)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
