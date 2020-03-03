package testingdemo2

import (
	"testing"
)

func TestMonster_Store(t *testing.T) {
	monster := Monster{
		Name:  "蒋冬莲",
		Age:   26,
		Skill: "测试",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("store()测试失败，希望为%v,实际为%v", true, false)
	}
	t.Logf("store()测试成功！")
}

func TestMonster_Restore(t *testing.T) {
	monster := Monster{}
	res := monster.Restore()
	if !res {
		t.Fatalf("restore()错误，希望为%v,实际为%v", true, false)
	}
	if monster.Name != "蒋冬莲" {
		t.Fatalf("restore()错误，希望为%v,实际为%v", "蒋冬莲", monster.Name)
	}
	t.Logf("restore()测试成功！")
}
