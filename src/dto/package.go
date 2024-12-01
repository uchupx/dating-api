package dto

import (
	"strconv"
	"strings"
	"time"

	"github.com/uchupx/dating-api/pkg/helper"
	"github.com/uchupx/dating-api/src/model"
)

type Package struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Price       int64      `json:"price"`
	Description *string    `json:"description"`
	Features    []string   `json:"features"`
	Status      bool       `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func (u *Package) Model(p *model.Package) {
	u.ID = p.ID.String
	u.Name = p.Name.String
	u.Price = p.Price.Int64
	u.Status = p.Status.Bool
	u.CreatedAt = p.CreatedAt.Time

	if p.Description.Valid {
		u.Description = &p.Description.String
	}

	if p.Features.Valid {
		f := strings.Split(p.Features.String, ",")
		for _, v := range f {
			i, _ := strconv.Atoi(v)

			u.Features = append(u.Features, helper.FEATURE_MAP[int8(i)])
		}
	}

	if p.UpdatedAt.Valid {
		u.UpdatedAt = &p.UpdatedAt.Time
	}

	if p.DeletedAt.Valid {
		u.DeletedAt = &p.DeletedAt.Time
	}
}

func (u *Package) ToModel() model.Package {
	var m model.Package
	m.ID.String = u.ID
	m.Name.String = u.Name
	m.Price.Int64 = u.Price
	m.Status.Bool = u.Status
	m.CreatedAt.Time = u.CreatedAt

	if u.Description != nil {
		m.Description.String = *u.Description
	}

	if len(u.Features) > 0 {
		m.Features.String = strings.Join(u.Features, ",")
	}

	if u.DeletedAt != nil {
		m.DeletedAt.Time = *u.DeletedAt
	}
	return m
}
