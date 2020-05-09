package subsription



import (
	"errors"
	"time"
)

type Internship interface {
	Add(obj *Subscription) (*Subscription,error)
	GetById(id int64)(*Subscription,error)
	Get()([]*Subscription,error)
	Delete(id int64) error
	Update(obj *Subscription) (*Subscription,error)
}
type internshipStruct struct {
	Repo Repository
}
func NewInternship(repo Repository) Internship{
	return &internshipStruct{
		Repo: repo,
	}
}
func(inter *internshipStruct) checkForEmpty(obj *Subscription) error{
	if obj.Name == ""{
		return errors.New("client id is empty")
	}else if obj.Payment ==0{
		return errors.New("payment is empty")
	}else if obj.TimeDuration ==0{
		return errors.New("time duration is empty")
	}
	return nil
}

func(inter *internshipStruct) Add(obj *Subscription) (*Subscription,error){
	err:=inter.checkForEmpty(obj)
	if err!=nil{
		return nil, err
	}
	obj.CreatedAt = time.Now()
	obj.UpdatedAt = time.Now()
	newobj,err:=inter.Repo.Add(obj)
	return newobj,err
}
func(inter *internshipStruct) Get()([]*Subscription,error){
	objects,err:=inter.Repo.Get()
	if err!=nil{
		return nil, err
	}
	return objects,err
}
func(inter *internshipStruct) Update(obj *Subscription)(*Subscription,error){
	obj.UpdatedAt=time.Now()
	obj,err:=inter.Repo.Update(obj)
	if err!=nil{
		return nil,err
	}
	return obj,err
}
func(inter *internshipStruct) GetById(id int64)(*Subscription,error){
	obj,err:=inter.Repo.GetById(id)
	if err!=nil{
		return nil,err
	}
	return obj,err
}
func(inter *internshipStruct) Delete(id int64) error{
	obj,err:=inter.Repo.GetById(id)
	if err!=nil{
		return err
	}
	err = inter.Repo.Delete(obj)
	if err!=nil{
		return err
	}
	return nil
}