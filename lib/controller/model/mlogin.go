package model

type Mlogin struct {
	Uid string
	Pwd string
}

type Mloginrespond struct {
	Title string `json:"title"`
	Rcode int `json:"rcode"`
}
