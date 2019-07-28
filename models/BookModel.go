package models

import (
	// "errors"
	// "log"
	"fmt"
	"time"

	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/validation"
	// . "weread/utils"
)

//图书信息表
type Book struct {
	Id          uint32    `orm:"pk"`
	Name        string    `orm:"size(40)"`
	Isbn        uint64    `orm:"" form:"Isbn"`
	Library     uint16    `orm:""`
	CoverPath   string    `orm:"null"`
	Added       time.Time `orm:"auto_now_add;type(datetime)"`
	Description string    `orm:"null;size(1000)"`
}

func init() {
	// this step will create SQL table
	orm.RegisterModel(new(Book))
}

var BOOK = []Book{
	{Id: 0, Name: "NULL-BOOK", Isbn: 0, Library: 0},
}

func InsertInitBooks() {
	o := orm.NewOrm()
	num, err := o.InsertMulti(len(BOOK), BOOK)
	fmt.Printf("insert books, %d success\n", num)
	if err != nil {
		fmt.Println("error: ", err)
	}
}

//添加图书
func AddBook(b *Book) (int64, error) {
	o := orm.NewOrm()
	book := new(Book)
	//check book exist or not
	//TODO check if book exist, return error
	book.Name = b.Name
	book.Isbn = b.Isbn
	book.Library = b.Library
	book.CoverPath = b.CoverPath

	id, err := o.Insert(book)
	return id, err
}

func UpdateBook(b *Book) (int64, error) {
	o := orm.NewOrm()

	book := make(orm.Params)
	if len(b.Name) > 0 {
		book["Name"] = b.Name
	}
	if b.Isbn > 0 {
		book["Isbn"] = b.Isbn
	}
	if b.Library > 0 {
		book["Library"] = b.Library
	}
	if len(b.CoverPath) > 0 {
		book["CoverPath"] = b.CoverPath
	}
	var table Book
	num, err := o.QueryTable(table).Filter("Id", b.Id).Update(book)
	return num, err
}

func DelBookById(Id uint32) (int64, error) {
	o := orm.NewOrm()
	// TODO check book exist or not, if exist, return error
	status, err := o.Delete(&Book{Id: Id})
	return status, err
}

func DelBookByName(Name string) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Book{Name: Name})
	return status, err
}
