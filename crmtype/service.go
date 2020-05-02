package crmtype



import (
	"errors"
	"time"
)

type Internship interface {
	Add(obj *CrmType) (*CrmType,error)
	GetById(id int64)(*CrmType,error)
	Get()([]*CrmType,error)
	Delete(id int64) error
	Update(obj *CrmType) (*CrmType,error)
}
type internshipStruct struct {
	Repo Repository
}
func NewInternship(repo Repository) Internship{
	return &internshipStruct{
		Repo: repo,
	}
}
func(inter *internshipStruct) checkForEmpty(obj *CrmType) error{
	if obj.Name == ""{
		return errors.New("client id is empty")
	}
	return nil
}

func(inter *internshipStruct) Add(obj *CrmType) (*CrmType,error){
	err:=inter.checkForEmpty(obj)
	if err!=nil{
		return nil, err
	}
	obj.CreatedAt = time.Now()
	obj.UpdatedAt = time.Now()
	newobj,err:=inter.Repo.Add(obj)
	return newobj,err
}
func(inter *internshipStruct) Get()([]*CrmType,error){
	objects,err:=inter.Repo.Get()
	if err!=nil{
		return nil, err
	}
	return objects,err
}
func(inter *internshipStruct) Update(obj *CrmType)(*CrmType,error){
	obj.UpdatedAt=time.Now()
	obj,err:=inter.Repo.Update(obj)
	if err!=nil{
		return nil,err
	}
	return obj,err
}
func(inter *internshipStruct) GetById(id int64)(*CrmType,error){
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