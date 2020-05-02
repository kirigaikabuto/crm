package crmsystem



import (
	"errors"
	"time"
)

type Internship interface {
	Add(obj *CrmSystem) (*CrmSystem,error)
	GetById(id int64)(*CrmSystem,error)
	Get()([]*CrmSystem,error)
	Delete(id int64) error
	Update(obj *CrmSystem) (*CrmSystem,error)
	GetByClientId(id int64) (*CrmSystem,error)
}
type internshipStruct struct {
	Repo Repository
}
func NewInternship(repo Repository) Internship{
	return &internshipStruct{
		Repo: repo,
	}
}
func(inter *internshipStruct) checkForEmpty(obj *CrmSystem) error{
	if obj.ClientId==0{
		return errors.New("client id is empty")
	}else if obj.TypeId==0{
		return errors.New("type id is empty")
	}else if obj.SubscriptionId==0{
		return errors.New("subscription id is empty")
	}
	return nil
}

func(inter *internshipStruct) Add(obj *CrmSystem) (*CrmSystem,error){
	err:=inter.checkForEmpty(obj)
	if err!=nil{
		return nil, err
	}
	obj.CreatedAt = time.Now()
	obj.UpdatedAt = time.Now()
	newobj,err:=inter.Repo.Add(obj)
	return newobj,err
}
func(inter *internshipStruct) Get()([]*CrmSystem,error){
	objects,err:=inter.Repo.Get()
	if err!=nil{
		return nil, err
	}
	return objects,err
}
func(inter *internshipStruct) Update(obj *CrmSystem)(*CrmSystem,error){
	obj.UpdatedAt=time.Now()
	obj,err:=inter.Repo.Update(obj)
	if err!=nil{
		return nil,err
	}
	return obj,err
}
func(inter *internshipStruct) GetById(id int64)(*CrmSystem,error){
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
func(inter *internshipStruct) GetByClientId(id int64)(*CrmSystem,error){
	object,err:=inter.Repo.GetByClientId(id)
	if err!=nil{
		return nil,err
	}
	return object,nil
}