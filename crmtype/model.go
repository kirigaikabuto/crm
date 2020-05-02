package crmtype


import "time"
type Repository interface {
	Add(obj *CrmType) (*CrmType,error)
	GetById(id int64)(*CrmType,error)
	Get()([]*CrmType,error)
	Delete(obj *CrmType) error
	Update(obj *CrmType) (*CrmType,error)
}
type CrmType struct {
	Id        int64  `json:"id" pg:"id,pk"`
	Name string `json:"name,omitempty" pg:"name"`
	CreatedAt time.Time `json:"created_at,omitempty" pg:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" pg:"updated_at"`
}