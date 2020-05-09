package subsription


import "time"
type Repository interface {
	Add(obj *Subscription) (*Subscription,error)
	GetById(id int64)(*Subscription,error)
	Get()([]*Subscription,error)
	Delete(obj *Subscription) error
	Update(obj *Subscription) (*Subscription,error)
}
type Subscription struct {
	Id        int64  `json:"id" pg:"id,pk"`
	Name string `json:"name,omitempty" pg:"name"`
	TimeDuration int64 `json:"time_duration" pg:"time_duration"`
	Payment int64 `json:"payment"  pg:"payment"`
	CreatedAt time.Time `json:"created_at,omitempty" pg:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" pg:"updated_at"`
}