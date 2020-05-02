package crmclientmanager


import "time"
type Repository interface {
	Add(obj *CrmClientManager) (*CrmClientManager,error)
	GetById(id int64)(*CrmClientManager,error)
	Get()([]*CrmClientManager,error)
	Delete(obj *CrmClientManager) error
	Update(obj *CrmClientManager) (*CrmClientManager,error)
	GetByCrmId(id int64)([]*CrmClientManager,error)
}
type CrmClientManager struct {
	Id        int64  `json:"id" pg:"id,pk"`
	FirstName string `json:"first_name,omitempty" pg:"first_name"`
	LastName string `json:"last_name,omitempty" pg:"last_name"`
	Username string `json:"username,omitempty" pg:"username"`
	Password string `json:"password,omitempty" pg:"password"`
	Email string `json:"email,omitempty" pg:"email"`
	Phone string `json:"phone,omitempty" pg:"phone"`
	CrmId int64 `json:"crm_id,omitempty" pg:"crm_id"`
	CreatedAt time.Time `json:"created_at,omitempty" pg:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" pg:"updated_at"`
}