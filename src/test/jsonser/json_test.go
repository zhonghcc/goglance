package jsonser

import (
	"encoding/json"
	"testing"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestSerJson(t *testing.T) {
	user := new(User)
	user.Name = "王biang"
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
