package api

import (
	modelctrl "Go_Server_tripleS/lib/controller/model"
	dbapi "Go_Server_tripleS/lib/mongo/api"
	"encoding/json"

)

func MGetAllDevice(payload []byte)  (modelctrl.Mgethomerespond,bool) {
	m:= modelctrl.Mgethome{}
	bytes:=	[]byte(payload)

	ret:= modelctrl.Mgethomerespond{Title:"RMGETDEVICE"}

	err:=	json.Unmarshal(bytes,&m)
	if err!=nil {
		println("parse json fail")
	}

	data,queryerr:= dbapi.MGetAllDevice(m.Uid)

	if queryerr {
		println("Query err")
		ret.Rcode=201
		return ret,true
	}
	ret.Rcode=200
	ret.Lhome=data.Lhome
	ret.Ldevice=data.Ldevice
	ret.UID=data.UID
	ret.Uname=data.Uname
	ret.Permission=data.Permission

	//fmt.Print("Return value:",ret)

	return ret,false

}

func HomeGetAllDevice(payload []byte)  (result modelctrl.LHomeDeviceRsp){
	rqmodel:= modelctrl.LHomeDevice{}
	bytes:=	[]byte(payload)

	result.Title="RGETDEVICE"

	if err:=json.Unmarshal(bytes,&rqmodel); err!=nil {
		println("parse json fail")
		result.Rcode=100
		return result
	}

	ldev,err:= dbapi.GetAllDevice(rqmodel.Hid)

	if err {
		result.Rcode=201
		return result
	}

	result.Rcode=200
	result.Ldevice=ldev

	return result

}