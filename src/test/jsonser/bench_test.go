package jsonser

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

// type User struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

func serJSON(user User) string {
	b, e := json.Marshal(user)
	if e == nil {
		return string(b)
	}
	return ""
}

func unSerJSON(str string) User {
	byteArr := []byte(str)
	var user User
	err := json.Unmarshal(byteArr, &user)
	if err == nil {
		return user
	}
	return user
}

func serJSON2(user User) string {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	b, e := json.Marshal(user)
	if e == nil {
		return string(b)
	}
	return ""
}

func unSerJSON2(str string) User {
	byteArr := []byte(str)
	var user User
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal(byteArr, &user)
	if err == nil {
		return user
	}
	return user
}

func BenchmarkSerJSON(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var user User
		user.Name = "王你好"
		user.Age = 18
		var address Address
		address.Name = "家庭住址"
		address.Detail = "家庭住址"
		user.Address = append(user.Address, address)
		serJSON(user)
	}
}

func BenchmarkSerJSON2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var user User
		user.Name = "王你好"
		user.Age = 18
		serJSON2(user)
		var address Address
		address.Name = "家庭住址"
		address.Detail = "家庭住址"
		user.Address = append(user.Address, address)
	}
}
func BenchmarkUnserJSON(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str := `{"name":"王你好","age":18,"address":[{"name":"家庭住址","detail":"家庭住址"}]}`
		unSerJSON(str)
	}
}

func BenchmarkUnserJSON2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str := `{"name":"王你好","age":18,"address":[{"name":"家庭住址","detail":"家庭住址"}]}`
		unSerJSON2(str)
	}
}
