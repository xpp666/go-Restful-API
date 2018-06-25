package main

import (
	"time"
	"fmt"
)
// AddUser 增加新用户
func AddUser(user User) (bool, error) {
	//user.ID = string()
	defer db.Close()
	t1,_ := time.ParseDuration("8h")
	user.CreateTime= time.Now().Add(t1)
	user.UpdateTime=time.Now().Add(t1)
	if err := db.Create(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}

// DelByID 通过ID删除用户
func DelByID(ID int) (bool, error) {
	defer db.Close()
	if err := db.Where("ID=?", ID).Delete(&User{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

// DelByName 通过用户名删除用户
func DelByName(Name string) (bool, error) {
	defer db.Close()
	if err := db.Where("Name=?", Name).Delete(&User{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

//FindAllUser 查询所有用户
func FindAllUser() ([]User, error) {
	defer db.Close()
	users:=[]User{}
	if err := db.Model(&User{}).Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

// FindUserByID 通过ID查找用户
func FindUserByID(ID int) (User, error) {
	// return db.Model(&User{}).Where(&User{ID: ID}).First(&User)
	defer db.Close()
	var u User
	if err := db.Where("ID = ?", ID).Find(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

// FindUserByName 通过用户名查找用户
func FindUserByName(Name string) (User, error) {
	// return db.Model(&User{}).Where(&User{ID: ID}).First(&User)
	defer db.Close()
	var u User
	if err := db.Where("Name = ?", Name).Find(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

// 通过ID修改用户信息
func UpdateUserByID(user *User)(bool,error) {
	defer db.Close()
	fmt.Println("uuuuu", user)
	t1,_ := time.ParseDuration("8h")
	user.UpdateTime = time.Now().Add(t1)
	if err := db.Model(&user).Where("ID = ?", user.ID).Update(&user).Error;err!=nil{
		return false,err
	}
	return true, nil
}

// 通过用户名修改用户信息
func UpdateUserByName(user User)(bool,error) {
	defer db.Close()
	t1,_ := time.ParseDuration("8h")
	user.UpdateTime = time.Now().Add(t1)
	if err := db.Model(&user).Where("Name = ?", user.Name).Update(user).Error;err!=nil{
		return false,err
	}
	return true, nil
}
