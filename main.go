package main

import (
	"crm/clients"
	"crm/config"
	"crm/crmclientmanager"
	"crm/crmsystem"
	"crm/crmtype"
	"crm/redis_connect"
	"crm/subsription"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	PATH string  = ""
	conf  config.PostgreConfig
	PORT string = "8000"
)
var flags []cli.Flag=[]cli.Flag{
	&cli.StringFlag{
		Name: "config,c",
		Usage:"Load configuration from `FILE`",
		Destination: &PATH,
	},
	&cli.StringFlag{
		Name:"Host",
		Usage:"Postgre host",
		Destination: &conf.Host,
	},
	&cli.StringFlag{
		Name:"Port",
		Usage:"Postgre port",
		Destination: &conf.Port,
	},
	&cli.StringFlag{
		Name:"User",
		Usage:"Postgre user",
		Destination: &conf.User,
	},
	&cli.StringFlag{
		Name:"Password",
		Usage: "Postgre password",
		Destination: &conf.Password,
	},
	&cli.StringFlag{
		Name:"Database",
		Usage:"Postgre database",
		Destination: &conf.Database,
	},
}
func main(){
	app :=cli.NewApp()
	app.Name = "HH Rest Api"
	app.Flags = flags
	app.Action = runRestApi
	fmt.Println(app.Run(os.Args))
}


func extractConfigPostgre(path string, conf *config.PostgreConfig) error{
	file, _ := ioutil.ReadFile(path)
	err:=json.Unmarshal(file, &conf)
	if err!=nil{
		return err
	}
	return nil
}
func runRestApi(*cli.Context) error{
	if PATH!=""{
		err:=extractConfigPostgre(PATH,&conf)
		if err!=nil{
			return err
		}
	} else if conf.Host=="" {
		return errors.New("Nothing found in host")
	} else if conf.Port==""{
		return errors.New("Nothing found in port")
	} else if conf.User==""{
		return errors.New("Nothing found in user")
	} else if conf.Host ==""{
		return errors.New("Nothing found in host")
	} else if conf.Database ==""{
		return errors.New("Nothing found in database")
	}
	router:=mux.NewRouter()
	//clients
	clientsrepo,err:=clients.NewPostgreStore(conf)
	if err!=nil{
		return err
	}
	clientsservice:=clients.NewInternship(clientsrepo)
	clientsendpoints:=clients.NewEndpoints(clientsservice)
	router.Methods("POST").Path("/clients/").HandlerFunc(clientsendpoints.Add())
	router.Methods("GET").Path("/clients/").HandlerFunc(clientsendpoints.Get())
	router.Methods("GET").Path("/clients/{id}").HandlerFunc(clientsendpoints.GetById("id"))
	router.Methods("PUT").Path("/clients/{id}").HandlerFunc(clientsendpoints.Update("id"))
	router.Methods("DELETE").Path("/clients/{id}").HandlerFunc(clientsendpoints.Delete("id"))
	//crm_types
	crmtypesrepo,err:=crmtype.NewPostgreStore(conf)
	if err!=nil{
		return nil
	}
	crmtypesservice:=crmtype.NewInternship(crmtypesrepo)
	crmtypesendpoints:=crmtype.NewEndpoints(crmtypesservice)
	router.Methods("POST").Path("/crmtypes/").HandlerFunc(crmtypesendpoints.Add())
	router.Methods("GET").Path("/crmtypes/").HandlerFunc(crmtypesendpoints.Get())
	router.Methods("GET").Path("/crmtypes/{id}").HandlerFunc(crmtypesendpoints.GetById("id"))
	router.Methods("PUT").Path("/crmtypes/{id}").HandlerFunc(crmtypesendpoints.Update("id"))
	router.Methods("DELETE").Path("/crmtypes/{id}").HandlerFunc(crmtypesendpoints.Delete("id"))
	//subscriptions
	subsriptionsrepo,err:=subsription.NewPostgreStore(conf)
	if err!=nil{
		return nil
	}
	subscriptionsservice:=subsription.NewInternship(subsriptionsrepo)
	subscriptionensdpoints:=subsription.NewEndpoints(subscriptionsservice)
	router.Methods("POST").Path("/subscriptions/").HandlerFunc(subscriptionensdpoints.Add())
	router.Methods("GET").Path("/subscriptions/").HandlerFunc(subscriptionensdpoints.Get())
	router.Methods("GET").Path("/subscriptions/{id}").HandlerFunc(subscriptionensdpoints.GetById("id"))
	router.Methods("PUT").Path("/subscriptions/{id}").HandlerFunc(subscriptionensdpoints.Update("id"))
	router.Methods("DELETE").Path("/subscriptions/{id}").HandlerFunc(subscriptionensdpoints.Delete("id"))
	//crmsystems
	redisconnect := redis_connect.ConnectRedis()
	crmsystems,err:=crmsystem.NewPostgreStore(conf)
	if err!=nil{
		return err
	}
	crmsystemsservice:=crmsystem.NewInternship(crmsystems,redisconnect)
	crmsystemsendpoints:=crmsystem.NewEndpoints(crmsystemsservice)
	router.Methods("POST").Path("/crmsystems/").HandlerFunc(crmsystemsendpoints.Add())
	router.Methods("GET").Path("/crmsystems/").HandlerFunc(crmsystemsendpoints.Get())
	router.Methods("GET").Path("/crmsystems/{id}").HandlerFunc(crmsystemsendpoints.GetById("id"))
	router.Methods("GET").Path("/crmsystems/{id}/byclientid/").HandlerFunc(crmsystemsendpoints.GetByClientId("id"))
	router.Methods("PUT").Path("/crmsystems/{id}").HandlerFunc(crmsystemsendpoints.Update("id"))
	router.Methods("DELETE").Path("/crmsystems/{id}").HandlerFunc(crmsystemsendpoints.Delete("id"))
	//crmclientmanagers
	crmclientmanagerrepo,err:=crmclientmanager.NewPostgreStore(conf)
	if err!=nil{
		return err
	}
	crmclientmanagerservice:=crmclientmanager.NewInternship(crmclientmanagerrepo)
	crmclientmanageendpoints:=crmclientmanager.NewEndpoints(crmclientmanagerservice)
	router.Methods("POST").Path("/crmclientmanagers/").HandlerFunc(crmclientmanageendpoints.Add())
	router.Methods("GET").Path("/crmclientmanagers/").HandlerFunc(crmclientmanageendpoints.Get())
	router.Methods("GET").Path("/crmclientmanagers/{id}").HandlerFunc(crmclientmanageendpoints.GetById("id"))
	router.Methods("GET").Path("/crmclientmanagers/{id}/bycrmid/").HandlerFunc(crmclientmanageendpoints.GetByCrmId("id"))
	router.Methods("PUT").Path("/crmclientmanagers/{id}").HandlerFunc(crmclientmanageendpoints.Update("id"))
	router.Methods("DELETE").Path("/crmclientmanagers/{id}").HandlerFunc(crmclientmanageendpoints.Delete("id"))
	fmt.Println("Server is running on port "+PORT)
	err = http.ListenAndServe(":"+PORT,router)
	if err!=nil{
		return err
	}
	return nil
}