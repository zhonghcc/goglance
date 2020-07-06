package jsonser

import (
	"encoding/json"
	"testing"
)

type User struct {
	Name    string    `json:"name"`
	Age     int       `json:"age"`
	Address []Address `json:"address"`
}

type Address struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

func TestSerJson(t *testing.T) {
	user := new(User)
	user.Name = "王你好"
	user.Age = 18
	var address Address
	address.Name = "家庭住址"
	address.Detail = "家庭住址"
	user.Address = append(user.Address, address)
	t.Logf("%+v", user)
	b, e := json.Marshal(user)
	if e == nil {
		t.Log(string(b))
	} else {
		t.Error(e)
	}

}

func TestUnSerJson(t *testing.T) {
	str := "{\"name\":\"1\"}"
	t.Log(str)
	byteArr := []byte(str)
	var user User
	err := json.Unmarshal(byteArr, &user)
	if err == nil {
		t.Logf("%+v", user)
	} else {
		t.Error(err)
	}

}

func TestUnserUnknownJson(t *testing.T) {
	str := `
	{
		"name":"王富贵",
		"age":11,
		"children":[
			{
				"name":"王哈哈"
			},
			{
				"name":"王婀娜"
			}
		]
	}
	
	`
	byteArr := []byte(str)
	var known interface{}
	err := json.Unmarshal(byteArr, &known)
	if err == nil {
		t.Logf("%+v", known)
	} else {
		t.Error(err)
	}
}
