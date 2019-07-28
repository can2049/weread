//记载所有图书馆活动，包括管理员添加图书、用户借阅、归还图书
// 这个表很庞大，要考虑分表之类的事情
package models

import (
	// "errors"
	// "log"
	"time"

	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/validation"
	// . "weread/utils"
)

//各种操作记录表
type Record struct {
	Id          uint64
	UserId      uint32 `orm:""`
	BookId      uint32 `orm:"null"`
	OperationId uint8  `orm:""`
	LibraryId   uint64 `orm:"null"`
	// TODO  add sql comment
	Date time.Time `orm:"auto_now_add;type(datetime);comment(comment);description(UTC time)"`
	// 备用，记事
	Note string `orm:"null;size(500)"`
}

func init() {
	// this step will create SQL table
	orm.RegisterModel(new(Record))
}

func AddRecord(r *Record) (int64, error) {
	o := orm.NewOrm()
	record := new(Record)
	record.UserId = r.UserId
	record.BookId = r.BookId
	record.OperationId = r.OperationId
	record.LibraryId = r.LibraryId

	id, err := o.Insert(record)
	return id, err
}

func insertTestRecord() {
	r1 := Record{UserId: 1, BookId: 1, OperationId: 1, LibraryId: 1}
	r2 := Record{UserId: 2, BookId: 3, OperationId: 4, LibraryId: 5}
	r3 := Record{UserId: 3, OperationId: 6}

	AddRecord(&r1)
	// time.Sleep(1 * time.Second)
	AddRecord(&r2)
	AddRecord(&r3)
}
