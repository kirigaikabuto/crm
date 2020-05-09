package crmsystem
import (
	"crm/config"
	"crm/redis_connect"
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
	redisconnect := redis_connect.ConnectRedis()
	crmsystems,err:=NewPostgreStore(conf)
	if err!=nil{
		t.Error(err)
	}
	crmsystemsservice:=NewInternship(crmsystems,redisconnect)
	object:=&CrmSystem{
		ClientId: 1,
		TypeId: 1,
		SubscriptionId: 1,
	}
	newobject,err:=crmsystemsservice.Add(object)
	if err!=nil{
		t.Error(err)
	}
	newid=newobject.Id
	fmt.Println(newobject)
	fmt.Println(t.Name()," ended successfully")
}
func TestRepo_GetById(t *testing.T) {
	redisconnect := redis_connect.ConnectRedis()
	crmsystems,err:=NewPostgreStore(conf)
	if err!=nil{
		t.Error(err)
	}
	crmsystemsservice:=NewInternship(crmsystems,redisconnect)
	object,err:=crmsystemsservice.GetById(newid)
	fmt.Println(object)
	fmt.Println(t.Name()," ended successfully")
}
func TestRepo_Update(t *testing.T) {
	redisconnect := redis_connect.ConnectRedis()
	crmsystems,err:=NewPostgreStore(conf)
	if err!=nil{
		t.Error(err)
	}
	crmsystemsservice:=NewInternship(crmsystems,redisconnect)
	object,err:=crmsystemsservice.GetById(newid)
	if err!=nil{
		t.Error(err)
	}
	object.ClientId=2
	updated_object,err:=crmsystemsservice.Update(object)
	fmt.Println(updated_object)
	fmt.Println(t.Name()," ended successfully")
}
func TestRepo_Delete(t *testing.T) {
	redisconnect := redis_connect.ConnectRedis()
	crmsystems,err:=NewPostgreStore(conf)
	if err!=nil{
		t.Error(err)
	}
	crmsystemsservice:=NewInternship(crmsystems,redisconnect)
	err = crmsystemsservice.Delete(newid)
	fmt.Println(t.Name()," ended successfully")
}
