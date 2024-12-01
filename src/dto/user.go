package dto

import (
	"strconv"
	"strings"
	"time"

	"github.com/uchupx/dating-api/pkg/helper"
	"github.com/uchupx/dating-api/src/model"
)

type UserRequest struct {
	Name    *string `json:"name"`
	Gender  *string `json:"gender"`
	Address *string `json:"address"`
	DOB     *string `json:"date_of_birth"`
	Phone   *string `json:"phone"`
}

type User struct {
	ID          string     `json:"id"`
	Password    string     `json:"-"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	Name        string     `json:"name"`
	Gender      string     `json:"gender"`
	Address     string     `json:"address"`
	DOB         *time.Time `json:"date_of_birth"`
	Phone       string     `json:"phone"`
	Features    []string   `json:"features"`
	ClientAppId string     `json:"-"`
	Created     time.Time  `json:"created"`
	Updated     *time.Time `json:"updated"`
}

func (u *User) Update(p *UserRequest) {
	if p.Name != nil {
		u.Name = *p.Name
	}

	if p.Address != nil {
		u.Address = *p.Address
	}

	if p.Gender != nil {
		u.Gender = *p.Gender
	}

	if p.DOB != nil {
		t, _ := time.Parse("2006-01-02", *p.DOB)
		u.DOB = &t
	}

	if p.Phone != nil {
		u.Phone = *p.Phone
	}

}

func (u *User) Model(p *model.User) {
	u.ID = p.ID.String
	u.Username = p.Username.String
	u.Name = p.Name.String
	u.Gender = p.Gender.String
	u.Address = p.Address.String
	u.Phone = p.Phone.String
	u.ClientAppId = p.ClientAppID.String
	u.Email = p.Email.String
	u.Created = p.CreatedAt.Time

	if p.UpdatedAt.Valid {
		u.Updated = &p.UpdatedAt.Time
	}

	if p.DOB.Valid {
		u.DOB = &p.DOB.Time
	}

	if p.Features.Valid {
		f := strings.Split(p.Features.String, ",")
		for _, v := range f {
			i, _ := strconv.Atoi(v)

			u.Features = append(u.Features, helper.FEATURE_MAP[int8(i)])
		}
	}
}

func (u *User) ToModel() model.User {
	var m model.User

	m.ID.String = u.ID
	m.ClientAppID.String = u.ClientAppId
	m.Username.String = u.Username
	m.Password.String = u.Password
	m.Email.String = u.Email
	m.Gender.String = u.Gender
	m.Address.String = u.Address
	m.Phone.String = u.Phone
	m.CreatedAt.Time = u.Created

	if u.Updated != nil {
		m.UpdatedAt.Time = *u.Updated
	}

	if u.DOB != nil {
		m.DOB.Time = *u.DOB
	}

	return m
}
