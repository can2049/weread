package models

import (
	// "errors"
	// "log"
	// "time"
	"fmt"

	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/validation"
	// . "weread/utils"
)

//用户表
type Operation struct {
	Id   uint8  `orm:"pk;unique"`
	Name string `orm:"size(20);unique"`
}

var OPERATION = []Operation{
	{Id: 0, Name: "NULL-OPERATION"},
	// for normal user
	{Id: 1, Name: "UserBorrowBook"},
	{Id: 2, Name: "UserReturnBook"},

	//for admin
	{Id: 11, Name: "AdminAddNewBook"},
	{Id: 12, Name: "AdminModifyBookInfo"},
	{Id: 13, Name: "AdminDeleteBook"},

	{Id: 21, Name: "UserRigister"},
	{Id: 22, Name: "UserModifyInfo"},

	{Id: 31, Name: "AdminAddUser"},
	{Id: 32, Name: "AdminModifyUserInfo"},
	{Id: 33, Name: "AdminDeleteUser"},
}

func init() {
	// this step will create SQL table
	orm.RegisterModel(new(Operation))
}

func InsertInitOperations() {
	o := orm.NewOrm()
	num, err := o.InsertMulti(len(OPERATION), OPERATION)
	fmt.Printf("insert operation, %d items successed.\n", num)
	if err != nil {
		fmt.Println("error: ", err)
	}
}

// func AddOperation(o *Operation)(int64, error){
// 	o:=orm.NewOrm()
// 	operation := new(Operation)
// 	operation.
// }
