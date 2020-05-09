package clients

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
	clientsrepo,err:=NewPostgreStore(conf)
	clientsservice:=NewInternship(clientsrepo)
	if err!=nil{
		t.Error(err)
	}
	client:=&Client{
		FirstName: "Tleugazy",
		LastName: "Tleugazy",
		Password: "newpasword",
		Username: "newusersname",
		Email: "newuser@gmail.com",
		Phone: "7086394516",
	}
	newclient,err:=clientsservice.Add(client)
	if err!=nil{
		t.Error(err)
	}
	newid = newclient.Id
	fmt.Println(t.Name()," ended successfully")
}
func TestRepo_GetById(t *testing.T) {

	clientsrepo,err:=NewPostgreStore(conf)
	clientsservice:=NewInternship(clientsrepo)
	if err!=nil{
		t.Error(err)
	}
	object,err:=clientsservice.GetById(newid)
	if err!=nil{
		t.Error(err)
	}
	fmt.Println(object)
	fmt.Println(t.Name()," ended successfully")
}
func TestRepo_Update(t *testing.T) {
	clientsrepo,err:=NewPostgreStore(conf)
	clientsservice:=NewInternship(clientsrepo)
	if err!=nil{
		t.Error(err)
	}
	obj,err:=clientsservice.GetById(newid)
	if err!=nil{
		t.Error(err)
	}
	obj.Username = "newnewusername"
	object,err:=clientsservice.Update(obj)
	if err!=nil{
		t.Error(err)
	}
	fmt.Println(object)
	fmt.Println(t.Name()," ended successfully")
}
func TestRepo_Delete(t *testing.T) {
	clientsrepo,err:=NewPostgreStore(conf)
	clientsservice:=NewInternship(clientsrepo)
	if err!=nil{
		t.Error(err)
	}
	err=clientsservice.Delete(newid)
	if err!=nil{
		t.Error(err)
	}
	fmt.Println(t.Name()," ended successfully")
}