package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/uchupx/dating-api/pkg/errors"
	"github.com/uchupx/dating-api/src/dto"
	"github.com/uchupx/dating-api/src/repo"
	"github.com/uchupx/kajian-api/pkg/db"
)

type PackageService struct {
	DB          *db.DB
	PackageRepo *repo.PackageRepo
}

func (PackageService) Name() string {
	return "PackageService"
}

func (s *PackageService) GetPackages(ctx context.Context) (*dto.Response, *errors.ErrorMeta) {
	data, err := s.PackageRepo.GetPackages(ctx)
	if err != nil {
		return nil, serviceError(500, fmt.Errorf("[%s - GetPackages] failed to get packages, error: %w", s.Name(), err))
	}

	var pkgs []dto.Package
	for _, d := range data {
		var pkg dto.Package
		pkg.Model(&d)
		pkgs = append(pkgs, pkg)
	}

	return &dto.Response{
		Status:  200,
		Message: "Success",
		Data:    pkgs,
	}, nil
}

func (s *PackageService) Purchase(ctx context.Context, id string) (*dto.Response, *errors.ErrorMeta) {
	userId := ctx.Value("userData").(string)
	data, err := s.PackageRepo.GetPackageByID(ctx, id)
	if err != nil {
		return nil, serviceError(500, fmt.Errorf("[%s - Purchase] failed to purchase package, error: %w", s.Name(), err))
	}

	err = s.DB.FTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		for _, d := range strings.Split(data.Features.String, ",") {
			isExist, err := s.PackageRepo.GetActivePackageByUser(ctx, userId, d)
			if err != nil {
				return fmt.Errorf("[%s - Purchase] failed to purchase package, error: %w", s.Name(), err)
			}

			if isExist != nil {
				return fmt.Errorf("[%s - Purchase] failed to purchase package, error: %s", s.Name(), "package already purchased")
			}

			if err := s.PackageRepo.InsertUserPackage(ctx, userId, d, true, nil); err != nil {
				return fmt.Errorf("[%s - Purchase] failed to purchase package, error: %w", s.Name(), err)
			}
		}

		return nil
	}, nil)

	if err != nil {
		return nil, serviceError(500, err)
	}

	return &dto.Response{
		Status:  200,
		Message: "Success purchase package",
	}, nil

}
