package clients

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"crm/config"
)

type repo struct {
	db *pg.DB
}
func NewPostgreStore(config config.PostgreConfig)(Repository,error){
	db:= pg.Connect(&pg.Options{
		Addr: config.Host+":"+config.Port,
		User: config.User,
		Password: config.Password,
		Database: config.Database,
	})
	err:= createSchema(db)
	if err!=nil{
		return nil, err
	}
	return &repo{db: db},nil
}
func createSchema(db *pg.DB) error{
	for _, model := range []interface{}{(*Client)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists:true,
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (rep *repo) Add(obj *Client) (*Client,error){
	err:=rep.db.Insert(obj)
	if err!=nil{
		return nil, err
	}
	return obj,nil
}
func (rep *repo) GetById(id int64)(*Client,error){
	obj:=&Client{Id: id}
	err:=rep.db.Select(obj)
	if err!=nil{
		return nil,err
	}
	return obj,nil
}
func (rep *repo) Get()([]*Client,error){
	var objects []*Client
	err:=rep.db.Model(&objects).Select()
	return objects,err
}
func (rep *repo) Update(obj *Client) (*Client,error){
	err:=rep.db.Update(obj)
	if err!=nil{
		return nil,err
	}
	return obj,err
}
func (rep *repo) Delete(obj *Client) error{
	err:=rep.db.Delete(obj)
	if err!=nil{
		return err
	}
	return nil
}
