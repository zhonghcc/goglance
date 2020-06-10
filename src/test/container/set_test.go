package container

import "testing"

func getSetExist(m map[string]string, key string) bool {
	if _, exist := m[key]; exist == false {
		return false
	} else {
		return true
	}
}

func setSet(m map[string]string, key string) {
	if m != nil {
		m[key] = "1"
	}
}

func TestSet(t *testing.T) {
	set := map[string]string{}
	t.Log(getSetExist(set, "key"))
	setSet(set, "abc")
	t.Log(getSetExist(set, "abc"))
	t.Log(len(set))
	setSet(set, "abc")
	t.Log(getSetExist(set, "abc"))
	t.Log(len(set))
	setSet(set, "abcd")
	setSet(set, "abcde")
	t.Log(getSetExist(set, "abcd"))
	t.Log(len(set))

}
