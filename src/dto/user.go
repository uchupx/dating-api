package dto

import (
	"time"

	"github.com/uchupx/dating-api/src/model"
)

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
	Verified    bool       `json:"verified"`
	ClientAppId string     `json:"-"`
	Created     time.Time  `json:"created"`
	Updated     *time.Time `json:"updated"`
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
