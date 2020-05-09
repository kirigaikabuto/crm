package subsription

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
	subsriptionsrepo,err:=NewPostgreStore(conf)
	if err!=nil{
		t.Error(err)
	}
	subscriptionsservice:=NewInternship(subsriptionsrepo)
	subscription:=&Subscription{
		Name: "TestSubscription",
		TimeDuration: 60,
		Payment: 1000,

	}
	newsubscription,err:=subscriptionsservice.Add(subscription)
	if err!=nil{
		t.Error(err)
	}
	newid  = newsubscription.Id
	fmt.Println(newsubscription)
	fmt.Println(t.Name()," ended successfully")
}
func TestRepo_Update(t *testing.T) {
	subsriptionsrepo,err:=NewPostgreStore(conf)
	if err!=nil{
		t.Error(err)
	}
	subscriptionsservice:=NewInternship(subsriptionsrepo)
	subscription,err:=subscriptionsservice.GetById(newid)
	if err!=nil{
		t.Error(err)
	}
	subscription.Name="NewTestSuscription"
	_,err=subscriptionsservice.Update(subscription)
	if err!=nil{
		t.Error(err)
	}
	fmt.Println(t.Name()," ended successfully")
}
func TestRepo_GetById(t *testing.T) {
	subsriptionsrepo,err:=NewPostgreStore(conf)
	if err!=nil{
		t.Error(err)
	}
	subscriptionsservice:=NewInternship(subsriptionsrepo)
	subscription,err:=subscriptionsservice.GetById(newid)
	if err!=nil{
		t.Error(err)
	}
	fmt.Println(subscription)
	fmt.Println(t.Name()," ended successfully")
}
func TestRepo_Delete(t *testing.T) {
	subsriptionsrepo,err:=NewPostgreStore(conf)
	if err!=nil{
		t.Error(err)
	}
	subscriptionsservice:=NewInternship(subsriptionsrepo)
	err = subscriptionsservice.Delete(newid)
	fmt.Println(t.Name()," ended successfully")
}