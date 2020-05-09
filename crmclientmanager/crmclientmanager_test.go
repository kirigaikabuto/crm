package crmclientmanager

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
func TestRepo_Add(t *testing.T){
	crmclientmanagerrepo,err:=NewPostgreStore(conf)
	if err!=nil{
		t.Error(err)
	}
	crmclientmanagerservice:=NewInternship(crmclientmanagerrepo)
	object:=&CrmClientManager{
		FirstName: "manager_test",
		LastName: "manager_test",
		Username: "manager_test",
		Password: "manager_test",
		Email: "manager_test@gmail.com",
		Phone: "8123213",
		CrmId: 1,

	}
	newobject,err:=crmclientmanagerservice.Add(object)
	if err!=nil{
		t.Error(err)
	}
	newid = newobject.Id
	fmt.Println(t.Name()," ended successfully")
}
func TestRepo_GetById(t *testing.T) {
	crmclientmanagerrepo,err:=NewPostgreStore(conf)
	if err!=nil{
		t.Error(err)
	}
	crmclientmanagerservice:=NewInternship(crmclientmanagerrepo)
	object,err:=crmclientmanagerservice.GetById(newid)

	fmt.Println(t.Name()," ended successfully")
	fmt.Println(object)
}
func TestRepo_Update(t *testing.T) {
	crmclientmanagerrepo,err:=NewPostgreStore(conf)
	if err!=nil{
		t.Error(err)
	}
	crmclientmanagerservice:=NewInternship(crmclientmanagerrepo)
	object,err:=crmclientmanagerservice.GetById(newid)
	object.Username="test_manager_test"
	updated_object,err:=crmclientmanagerservice.Update(object)
	if err!=nil{
		t.Error(err)
	}
	fmt.Println(t.Name()," ended successfully")
	fmt.Println(updated_object)

}
func TestRepo_Delete(t *testing.T) {
	crmclientmanagerrepo,err:=NewPostgreStore(conf)
	if err!=nil{
		t.Error(err)
	}
	crmclientmanagerservice:=NewInternship(crmclientmanagerrepo)
	err = crmclientmanagerservice.Delete(newid)
	if err!=nil{
		t.Error(err)
	}
	fmt.Println(t.Name()," ended successfully")
}