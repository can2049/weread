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

//图书馆表，便于对馆进行调整，新增、合并、撤销
//一个图书馆可以对应多个ID，便于兼容历史图书馆的撤销
type Library struct {
	Id       uint16 `orm:"pk"`
	City     string `orm:"size(40)"`
	Building string `orm:"size(10)"`
	Note     string `orm:"null;size(500)"`
}

var LIBRARY = []Library{
	// inserted Id should bigger than 1
	{Id: 0, City: "NULL-LIBRARY", Building: "NULL"},
	{Id: 1, City: "SongShanHu", Building: "F1"},
	{Id: 2, City: "SongShanHu", Building: "M2"},
}

func init() {
	// this step will create SQL table
	orm.RegisterModel(new(Library))
}

func InsertInitLibraries() {
	o := orm.NewOrm()
	num, err := o.InsertMulti(len(LIBRARY), LIBRARY)
	fmt.Printf("insert library, %d items successed.\n", num)
	if err != nil {
		fmt.Println("error: ", err)
	}
}
