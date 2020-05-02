package crmclientmanager



import (
	"errors"
	"time"
)

type Internship interface {
	Add(obj *CrmClientManager) (*CrmClientManager,error)
	Get()([]*CrmClientManager,error)
	GetById(id int64)(*CrmClientManager,error)
	Update(obj *CrmClientManager)(*CrmClientManager,error)
	Delete(id int64) error
	GetByCrmId(id int64)([]*CrmClientManager,error)
}
type internshipStruct struct {
	Repo Repository
}
func NewInternship(repo Repository) Internship{
	return &internshipStruct{
		Repo: repo,
	}
}
func(inter *internshipStruct) checkForEmpty(obj *CrmClientManager) error{
	if obj.FirstName=="" {
		return errors.New("first Name is empty")
	}else if obj.LastName==""{
		return errors.New("last Name is empty")
	}else if obj.Username==""{
		return errors.New("username is empty")
	}else if obj.Password==""{
		return errors.New("password is empty")
	}else if obj.Email==""{
		return errors.New("email is empty")
	}else if obj.Phone==""{
		return errors.New("phone is empty")
	}else if obj.CrmId==0{
		return errors.New("crm id is empty")
	}

	return nil
}
func(inter *internshipStruct) checkForUsername(obj *CrmClientManager) error{
	objects,err:=inter.Repo.Get()
	if err!=nil{
		return err
	}
	for _,v := range objects{
		if v.Username == obj.Username && v.Id!=obj.Id{
			return errors.New("user with that username already exist")
		}
	}
	return nil
}
func(inter *internshipStruct) checkForEmail(obj *CrmClientManager) error {
	objects,err:=inter.Repo.Get()
	if err!=nil{
		return err
	}
	for _,v := range objects{
		if v.Email == obj.Email{
			return errors.New("user with that email already exist")
		}
	}
	return nil
}
func(inter *internshipStruct) Add(obj *CrmClientManager) (*CrmClientManager,error){
	err:=inter.checkForEmpty(obj)
	if err!=nil{
		return nil, err
	}
	err=inter.checkForUsername(obj)
	if err!=nil{
		return nil, err
	}
	err=inter.checkForEmail(obj)
	if err!=nil{
		return nil, err
	}
	obj.CreatedAt = time.Now()
	obj.UpdatedAt = time.Now()
	newobj,err:=inter.Repo.Add(obj)
	return newobj,err
}
func(inter *internshipStruct) Get()([]*CrmClientManager,error){
	objects,err:=inter.Repo.Get()
	if err!=nil{
		return nil, err
	}
	return objects,err
}
func(inter *internshipStruct) Update(obj *CrmClientManager)(*CrmClientManager,error){
	err:=inter.checkForUsername(obj)
	if err!=nil{
		return nil, err
	}
	obj.UpdatedAt=time.Now()
	obj,err=inter.Repo.Update(obj)
	if err!=nil{
		return nil,err
	}
	return obj,err
}
func(inter *internshipStruct) GetById(id int64)(*CrmClientManager,error){
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
func(inter *internshipStruct) GetByCrmId(id int64)([]*CrmClientManager,error){
	objects,err:=inter.Repo.GetByCrmId(id)
	if err!=nil{
		return nil,err
	}
	return objects,nil
}