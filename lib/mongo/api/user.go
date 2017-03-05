package api

import (

	"gopkg.in/mgo.v2/bson"
	"../../mongo/db"
	model "../../mongo/model"
	"fmt"
)

type Param struct {
	Id string
	Pwd string
}


func GetUserByID(data model.User) {
	id := data.Uid
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result := model.User{}

	if err := Db.C("user").Find(bson.M{"id": id}).One(&result); err != nil {
		fmt.Fprint("fail")
		return
	}

	fmt.Fprint(result)


}