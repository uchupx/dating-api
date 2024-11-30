package dto

import (
	"time"

	"github.com/uchupx/dating-api/src/model"
)

type Client struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Key       string    `json:"key"`
	Secret    string    `json:"secret"`
	CreatedAt time.Time `json:"created_at"`
}
type ClientPost struct {
	Name string `json:"name"`
}

func (u *Client) Model(p *model.Client) {
	u.ID = p.ID.String
	u.Name = p.Name.String
	u.Key = p.Key.String
	u.Secret = p.Secret.String
	u.CreatedAt = p.CreatedAt.Time
}

func (u *Client) ToModel() model.Client {
	var m model.Client

	m.ID.String = u.ID
	m.Name.String = u.Name
	m.Key.String = u.Key
	m.Secret.String = u.Secret
	m.CreatedAt.Time = u.CreatedAt

	return m
}
