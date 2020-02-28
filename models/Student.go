package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Student struct{
	Id int
	Number string
	Name string
	Sex string
	Phone string
	Qq string
	Clazzid int
	Gradeid int
	ClazzName string
	GradeName string
}

func NewStudent() *Student  {
	return &Student{}
}

//获取等登陆学生的信息
func (this *Student)GetInformation(account string) Student  {
	var student Student
	qb,_:=orm.NewQueryBuilder("mysql")
	qb.Select("student.*,clazz.clazz_name,grade.grade_name").
		From("student").
		InnerJoin("clazz").On("student.clazzid = clazz.id").
		InnerJoin("grade").On("student.gradeid = grade.id").
		Where("student.number = ?")
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql,account).QueryRow(&student)
	return student
}

//添加学生
func (this *Student)AddStudent(number,name,sex,phone,qq string,clazzid,gradeid int) error  {
	qb,_:=orm.NewQueryBuilder("mysql")
	qb.InsertInto("student","number","name","sex","phone","qq","clazzid","gradeid").
		Values("?","?","?","?","?","?","?" )
	sql := qb.String()
	o := orm.NewOrm()
	_,error :=o.Raw(sql,number,name,sex,phone,qq,clazzid,gradeid).Exec()
	return error
}

//批量添加学生
func (this *Student)AddBatchStudent(student []Student) error  {
	var error error
	qb,_:=orm.NewQueryBuilder("mysql")
	qb.InsertInto("student","number","name","sex","phone","qq","clazzid","gradeid").
		Values("?","?","?","?","?","?","?" )
	sql := qb.String()
	o := orm.NewOrm()
	for _, student := range student {
		_,error =o.Raw(sql,student.Number,student.Name,student.Sex,student.Phone,student.Qq,student.Clazzid,student.Clazzid).Exec()
	}
	return error
}

//修改密码
func (this *Student)ModifyBatchStudent(qq,password string) error  {
	var error error
	qb,_:=orm.NewQueryBuilder("mysql")
	qb.Update("user").
		Set("password=?").
		Where("account = ?")
	sql := qb.String()
	fmt.Println(password)
	fmt.Println(qq)
	fmt.Println(sql)
	o := orm.NewOrm()
	_,error =o.Raw(sql,password,qq).Exec()
	return error
}

//获取账号
func (this *Student)GetAccount(qq string) Student  {
	var student Student
	qb,_:=orm.NewQueryBuilder("mysql")
	qb.Select("student.*").
		From("student").
		Where("student.qq = ?")
	sql := qb.String()
	o := orm.NewOrm()

	o.Raw(sql,qq).QueryRow(&student)
	return student
}

//获取iD
func (this *Student)GetId(account string) int  {
	var student int
	qb,_:=orm.NewQueryBuilder("mysql")
	qb.Select("student.id").
		From("student").
		Where("student.number = ?")
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql,account).QueryRow(&student)
	return student
}
