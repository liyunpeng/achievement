package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)
type User struct {
	Id int
	Account string
	Password string
	Name string
	Type int
	Student     *Student   `orm:"rel(one)"` //一对一
}

//登陆判断
func (u *User)LoginJudge(account string,password string)  []User {
	var users []User
	qb,_:=orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("user.account").
		From("user").
		Where("account = ?").
		And("password = ?").And("type != 2")
	//返回sql语句
	sql := qb.String()
	// 执行 SQL 语句
	o := orm.NewOrm()
	o.Raw(sql,account,password).QueryRows(&users)
	return users
}

//登陆判断
func (u *User)ReceptionLoginJudge(account string,password string)  []User {
	var users []User
	qb,_:=orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("user.account").
		From("user").
		Where("account = ?").
		And("password = ?").And("type = 2")
	//返回sql语句
	sql := qb.String()
	// 执行 SQL 语句
	o := orm.NewOrm()
	o.Raw(sql,account,password).QueryRows(&users)
	return users
}

func NewUser() *User  {
	return &User{}
}

func (u *User)GetMenager(page,limit int) []User  {
	var users []User
	page = (page-1)*10
	qb,_:=orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("user.id,user.account,user.name").
		From("user").
		Where("type = 3").
		Limit(limit).
		Offset(page)
	//返回sql语句
	sql := qb.String()
	// 执行 SQL 语句
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&users)
	return users
}

func (u *User)GetMenagerConut() int  {
	var users []User
	qb,_:=orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("user.id,user.account,user.name").
		From("user").
		Where("type = 3")
	//返回sql语句
	sql := qb.String()
	// 执行 SQL 语句
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&users)
	return len(users)
}

func (u *User)DeleteMenagerConut(id int) error  {
	qb,_:=orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Delete("user").
		From("user").
		Where("id = ?")
	//返回sql语句
	sql := qb.String()
	fmt.Println(sql)
	// 执行 SQL 语句
	o := orm.NewOrm()
	_,error := o.Raw(sql,id).Exec()
	return error
}

func (u *User)AddMenagerExit(account string) []User  {
	var users []User
	qb,_:=orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("user.*").
		From("user").
		Where("account = ?")
	//返回sql语句
	sql := qb.String()
	fmt.Println(sql)
	// 执行 SQL 语句
	o := orm.NewOrm()
	o.Raw(sql,account).QueryRows(&users)
	return users
}

func (u *User)AddMenager(account,password,name string) error  {
	qb,_:=orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.InsertInto("user","user.account","user.password","user.name","user.type").
		Values("?","?","?","?")
	//返回sql语句
	sql := qb.String()
	fmt.Println(sql)
	// 执行 SQL 语句
	o := orm.NewOrm()
	_,error := o.Raw(sql,account,password,name,3).Exec()
	return error
}