package crmtype

import (
	"crm/config"
	"fmt"
	"testing"
)

var (
	conf =config.PostgreConfig{
		"localhost",
		"crm",
		"postgres",
		"passanya",
		"5432",
	}
	newid int64 = 0
)

func TestRepo_Add(t *testing.T) {
	crmtypesrepo,err:=NewPostgreStore(conf)
	if err!=nil{
		t.Error(err)
	}
	crmtypesservice:=NewInternship(crmtypesrepo)
	crm_type:=&CrmType{
		Name: "testType",
	}
	newcrm_type,err:=crmtypesservice.Add(crm_type)
	if err!=nil{
		t.Error(err)
	}
	newid  = newcrm_type.Id
	fmt.Println(newcrm_type)
	fmt.Println(t.Name()," ended successfully")
}
func TestRepo_Update(t *testing.T) {
	crmtypesrepo,err:=NewPostgreStore(conf)
	if err!=nil{
		t.Error(err)
	}
	crmtypesservice:=NewInternship(crmtypesrepo)
	crm_type,err:= crmtypesservice.GetById(newid)
	if err!=nil{
		t.Error(err)
	}
	crm_type.Name = "NewTestType"
	_,err=crmtypesservice.Update(crm_type)
	if err!=nil{
		t.Error(err)
	}
	fmt.Println(t.Name()," ended successfully")
}
func TestRepo_GetById(t *testing.T) {
	crmtypesrepo,err:=NewPostgreStore(conf)
	if err!=nil{
		t.Error(err)
	}
	crmtypesservice:=NewInternship(crmtypesrepo)
	crm_type,err:= crmtypesservice.GetById(newid)
	fmt.Println(crm_type)
	fmt.Println(t.Name()," ended successfully")
}
func TestRepo_Delete(t *testing.T) {
	crmtypesrepo,err:=NewPostgreStore(conf)
	if err!=nil{
		t.Error(err)
	}
	crmtypesservice:=NewInternship(crmtypesrepo)
	err=crmtypesservice.Delete(newid)
	if err!=nil{
		t.Error(err)
	}
	fmt.Println(t.Name()," ended successfully")
}