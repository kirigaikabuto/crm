package clients

import "time"
type Repository interface {
	Add(obj *Client) (*Client,error)
	GetById(id int64)(*Client,error)
	Get()([]*Client,error)
	Delete(obj *Client) error
	Update(obj *Client) (*Client,error)
}
type Client struct {
	Id        int64  `json:"id" pg:"id,pk"`
	FirstName string `json:"first_name,omitempty" pg:"first_name"`
	LastName string `json:"last_name,omitempty" pg:"last_name"`
	Username string `json:"username,omitempty" pg:"username"`
	Password string `json:"password,omitempty" pg:"password"`
	Email string `json:"email,omitempty" pg:"email"`
	Phone string `json:"phone,omitempty" pg:"phone"`
	CreatedAt time.Time `json:"created_at,omitempty" pg:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" pg:"updated_at"`
}