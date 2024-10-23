package service

import "mygo/internal/model"

type Repository interface {
	GetBanners() ([]model.Banner, error)
	CreateBanner(banner model.Banner) error
	UpdateBanner(banner model.Banner) error
	DeleteBanner(id int) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetBanners() ([]model.Banner, error) {
	return s.repo.GetBanners()
}

func (s *Service) CreateBanner(banner model.Banner) error {
	return s.repo.CreateBanner(banner)
}

func (s *Service) UpdateBanner(banner model.Banner) error {
	return s.repo.UpdateBanner(banner)
}

func (s *Service) DeleteBanner(id int) error {
	return s.repo.DeleteBanner(id)
}
