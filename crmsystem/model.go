package crmsystem



import "time"
type Repository interface {
	Add(obj *CrmSystem) (*CrmSystem,error)
	GetById(id int64)(*CrmSystem,error)
	Get()([]*CrmSystem,error)
	Delete(obj *CrmSystem) error
	Update(obj *CrmSystem) (*CrmSystem,error)
	GetByClientId(id int64) (*CrmSystem,error)
}
type CrmSystem struct {
	Id        int64  `json:"id" pg:"id,pk"`
	ClientId int64 `json:"client_id,omitempty" pg:"client_id"`
	TypeId int64 `json:"type_id,omitempty" pg:"type_id"`
	SubscriptionId int64 `json:"subscription_id,omitempty" pg:"subscription_id"`
	CreatedAt time.Time `json:"created_at,omitempty" pg:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" pg:"updated_at"`
}