package api

import (

	"gopkg.in/mgo.v2/bson"
	"../../mongo/db"
	model "../../mongo/model"

	"fmt"
)

func GetDeviceByDid(data string)  (model.Device,bool){

	id := data
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result := model.Device{}

	if err := Db.C("devices").Find(bson.M{"did": id}).One(&result); err != nil {
		print("Fail")
		return model.Device{},true
	}

	print(result.Did)
	fmt.Printf("%+v\n",result)
	return result,false
	
}

func GetAllDeviceByUid(uid string)  {

}

func UpdateStatusDevice(did string,status int)  bool{
	id := did
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	colQuerier := bson.M{"did": id}
	change := bson.M{"$set": bson.M{"status": status}}

	if err := Db.C("devices").Update(colQuerier,change); err != nil {
		print("Fail")
		return true
	}

	return false
}