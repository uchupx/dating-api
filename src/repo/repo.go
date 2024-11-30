package repo

import (
	"github.com/gofrs/uuid"
	"github.com/uchupx/dating-api/pkg/helper"
	"github.com/uchupx/kajian-api/pkg/logger"
)

type BaseRepo struct{}

func (BaseRepo) name() string {
	return "BaseRepo"
}

func (m BaseRepo) ID() *string {
	val, err := uuid.NewV7()
	if err != nil {
		logger.Logger.Errorf("[%s - ID] failed to generating uuid v7, %+v", m.name(), err)
		return nil
	}

	return helper.StringToPointer(val.String())
}
