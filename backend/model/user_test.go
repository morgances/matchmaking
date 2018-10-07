package model

import (
	"fmt"
	"testing"
	"time"
)

func Test_user(t *testing.T) {
	defer DB.Close()
	user1 := User{
		NickName: "zhanghao",
		Birthday: time.Now(),
		Sex:      true,
		Wechat:   "123",
		Phone:    "321",
	}
	user2 := User{
		NickName: "zhanghao2",
		Birthday: time.Now(),
		Sex:      false,
		Wechat:   "1234",
		Phone:    "4321",
	}
	update2 := User{
		NickName: "zhanghao2",
		Birthday: time.Now(),
		Sex:      false,
		Wechat:   "1234",
		Phone:    "4321",
		Charm:    200,
		CreateAt: time.Now(),
	}
	user3 := User{
		NickName: "zhanghao3",
		Birthday: time.Now(),
		Sex:      false,
		Wechat:   "12345",
		Phone:    "54321",
		Charm:    100,
	}
	var err error
	if err = UserService.Insert(&user1); err != nil {
		t.Error(err)
	}
	if err = UserService.Insert(&user2); err != nil {
		t.Error(err)
	}
	if err = UserService.Insert(&user3); err != nil {
		t.Error(err)
	}

	if _, err = UserService.FindByPhone("4321"); err != nil {
		t.Error(err)
	}
	var users []User
	if users, err = UserService.RecommendByCharm(); err != nil {
		t.Error(err)
	}
	fmt.Println(users)
	if err = UserService.Update(&update2); err != nil {
		t.Error(err)
	}
	if users, err = UserService.RecommendByCharm(); err != nil {
		t.Error(err)
	}
	fmt.Println(users)
	if err = UserService.DeleteByPhone("54321"); err != nil {
		t.Error(err)
	}
}
