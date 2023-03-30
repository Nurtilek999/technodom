package service

import (
	"fmt"
	"technodom/internal/app/config"
	"technodom/internal/entity"
	"technodom/internal/repository"
	"technodom/internal/util/logger"
)

type Service struct {
	repo   repository.Repository
	cache  repository.Cache
	config *config.TomlConfig
	logger logger.Logger
}

// NewService Отдельный файл сервис для каждой бизнес модули
func NewService(repo repository.Repository, cache repository.Cache, config *config.TomlConfig, logger logger.Logger) *Service {
	return &Service{repo, cache, config, logger}
}

type InterfaceService interface {
	FillCache() error
	GetList(limit int, offset int) ([]entity.Url, error)
	GetByID(id string) (entity.Url, error)
	Post(url entity.Url) error
	Patch(id, activeLink string) error
	Delete(id string) error
	Get(id string) (entity.Url, error)
}

func (s Service) FillCache() error {
	urls, err := s.GetList(1000, 1)
	if err != nil {
		return err
	}
	for _, url := range urls {
		s.cache.Add(url.HistoryLink, url.ActiveLink, s.config.Technodom.Ttl)
	}
	return nil
}
func (s Service) GetList(limit int, offset int) ([]entity.Url, error) {
	rows, err := s.repo.GetList(limit, offset)
	if err != nil {
		errMsg := fmt.Errorf("error in GetList method %w", err).Error()
		s.logger.Error(errMsg, map[string]interface{}{})
		return nil, fmt.Errorf("error in GetList method: %s", err.Error())
	}
	var urls []entity.Url
	for rows.Next() {
		var url entity.Url
		if err := rows.Scan(&url.ActiveLink, &url.HistoryLink); err != nil {
			errMsg := fmt.Errorf("error while scanning GetList %w", err).Error()
			s.logger.Error(errMsg, map[string]interface{}{})
			return nil, fmt.Errorf("error in scanning GetList, %s", err.Error())
		}
		urls = append(urls, url)
	}

	return urls, nil
}

func (s Service) GetByID(id string) (entity.Url, error) {
	row := s.repo.GetByID(id)
	var url entity.Url
	if err := row.Scan(&url.ActiveLink, &url.HistoryLink); err != nil {
		errMsg := fmt.Errorf("error in scanning GetByID %w", err).Error()
		s.logger.Error(errMsg, map[string]interface{}{})
		return entity.Url{}, fmt.Errorf("error in scanning GetByID, %s", err.Error())
	}
	return url, nil

}

func (s Service) Post(url entity.Url) error {
	err := s.repo.Post(url)
	if err != nil {
		errMsg := fmt.Errorf("error in Post method %w", err).Error()
		s.logger.Error(errMsg, map[string]interface{}{})
		return fmt.Errorf("error in Post mehtod: %s", err.Error())
	}
	return nil
}

func (s Service) Patch(id, activeLink string) error {
	err := s.repo.Update(id, activeLink)
	return err
}

func (s Service) Delete(id string) error {
	err := s.repo.Delete(id)
	return err
}

func (s Service) Get(id string) (entity.Url, error) {
	row := s.repo.GetByID(id)
	var url entity.Url
	if err := row.Scan(&url.ActiveLink, &url.HistoryLink); err != nil {
		errMsg := fmt.Errorf("error in scanning Get method %w", err).Error()
		s.logger.Error(errMsg, map[string]interface{}{})
		return entity.Url{}, fmt.Errorf("Error in scanning Get, %s", err.Error())
	}

	return url, nil
}
