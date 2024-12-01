package service

import (
	"context"
	"fmt"

	"github.com/uchupx/dating-api/src/dto"
	"github.com/uchupx/dating-api/src/repo"
)

type PackageService struct {
	PackageRepo *repo.PackageRepo
}

func (PackageService) Name() string {
	return "PackageService"
}

func (s *PackageService) GetPackages(ctx context.Context) (*dto.Response, error) {
	pkg, err := s.PackageRepo.GetPackages(ctx)
	if err != nil {
		return nil, fmt.Errorf("[%s - GetPackages] failed to get packages, error: %w", s.Name(), err)
	}

	return &dto.Response{
		Status:  200,
		Message: "Success",
		Data:    pkg,
	}, nil
}
