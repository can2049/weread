package models

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"weread/utils"
)

//用户表
type User struct {
	Id            uint32    `orm:"pk"`
	Username      string    `orm:"size(32)" form:"Username"  valid:"Required;MaxSize(20);MinSize(6)"`
	EmployId      string    `orm:"unique;size(10)"`
	Password      string    `orm:"size(32)" form:"Password" valid:"Required;MaxSize(20);MinSize(6)"`
	Repassword    string    `orm:"-" form:"Repassword" valid:"Required"`
	Email         string    `orm:"null;size(32)" form:"Email" valid:"Email"`
	Status        uint8     `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
	Lastlogintime time.Time `orm:"null;type(datetime)" form:"-"`
	Createtime    time.Time `orm:"type(datetime);auto_now_add" `
	Role          []*Role   `orm:"rel(m2m)"`
}

var USER = []User{
	{
		Id:       0,
		Username: beego.AppConfig.String("rbac_admin_user"),
		Password: utils.Pwdhash(beego.AppConfig.String("default_admin_passwd")),
		EmployId: "NULL",
		Email:    "null@null",
	},
}

func InsertInitUsers() {
	o := orm.NewOrm()
	num, err := o.InsertMulti(len(USER), USER)
	fmt.Printf("insert users, %d successed\n", num)
	if err != nil {
		println("error: ", err)
	}
}

func (u *User) TableName() string {
	return beego.AppConfig.String("rbac_user_table")
}

func (u *User) Valid(v *validation.Validation) {
	if u.Password != u.Repassword {
		v.SetError("Repassword", "两次输入的密码不一样")
	}
}

//验证用户信息
func checkUser(u *User) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&u)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

func init() {
	orm.RegisterModel(new(User))
}

/************************************************************/

//get user list
func Getuserlist(page int64, page_size int64, sort string) (users []orm.Params, count int64) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable(user)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&users)
	count, _ = qs.Count()
	return users, count
}

//添加用户
func AddUser(u *User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	user := new(User)
	user.Username = u.Username
	user.Password = utils.Pwdhash(u.Password)
	user.Email = u.Email
	user.Status = u.Status

	id, err := o.Insert(user)
	return id, err
}

//更新用户
func UpdateUser(u *User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	user := make(orm.Params)
	if len(u.Username) > 0 {
		user["Username"] = u.Username
	}
	// if len(u.Nickname) > 0 {
	// 	user["Nickname"] = u.Nickname
	// }
	if len(u.Email) > 0 {
		user["Email"] = u.Email
	}
	// if len(u.Remark) > 0 {
	// 	user["Remark"] = u.Remark
	// }
	if len(u.Password) > 0 {
		user["Password"] = utils.Pwdhash(u.Password)
	}
	if u.Status != 0 {
		user["Status"] = u.Status
	}
	if len(user) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table User
	num, err := o.QueryTable(table).Filter("Id", u.Id).Update(user)
	return num, err
}

func DelUserById(Id uint32) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&User{Id: Id})
	return status, err
}

func GetUserByUsername(username string) (user User) {
	user = User{Username: username}
	o := orm.NewOrm()
	o.Read(&user, "Username")
	return user
}

func GetUserById(id uint32) (user User) {
	user = User{Id: id}
	o := orm.NewOrm()
	o.Read(&user, "Id")
	return user
}
