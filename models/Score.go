package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Score struct {
	Examid int
	Studentid int
	Courseid int
	Score int
	Name string
}

type ScoreInformation struct {
	Number string
	Name string
	Sex string
	ClazzName string
	GradeName string
	ExamName string
	CourseName string
	Score int
	Time string
}

func NewScore() *Score  {
	return &Score{}
}

func (this *Score)GetId(name string) int  {
	var id int
	qb,_ := orm.NewQueryBuilder("mysql")
	qb.Select("course.id").
		From("course").
		Where("name = ?")
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql,name).QueryRow(&id)
	return id
}

//批量添加成绩
func (this *Score)AddBatchSCore(score []Score) error  {
	var error error
	qb,_:=orm.NewQueryBuilder("mysql")
	qb.InsertInto("escore","examid","studentid","courseid","score").
		Values("?","?","?","?" )
	sql := qb.String()
	o := orm.NewOrm()
	for _, scores := range score {
		_,error =o.Raw(sql,scores.Examid,scores.Studentid,scores.Courseid,scores.Score).Exec()
	}
	return error
}


//获取学生的成绩
func (this *Score)GetAll(examids,studentids int) []Score  {
	var id []Score
	qb,_ := orm.NewQueryBuilder("mysql")
	qb.Select("course.name","escore.score").
		From("escore").
		InnerJoin("course").
		Where("escore.courseid = course.id").
		And("escore.studentid = ?").
		And("escore.examid = ?")
	sql := qb.String()
	o := orm.NewOrm()
	_,error:=o.Raw(sql,studentids,examids).QueryRows(&id)
	fmt.Println(error)
	fmt.Println(id)
	return id
}

/*

type ScoreInformation struct {
	Number string
	Name string
	Sex string
	ClazzName string
	GradeName string
	ExamName string
	CourseName string
	Score int
	Time string
}

*/
//查询学生的成绩
func (this *Score)SelectAll(page,limit, class int,school string) []ScoreInformation {
	var score []ScoreInformation
	page = (page-1)*10
	qb,_ := orm.NewQueryBuilder("mysql")
	qb.Select("student.number","student.name","student.sex","clazz.clazz_name","grade.grade_name","escore.score","course.name as course_name","exam.exam_name","exam.time").
		From("student,clazz,grade,escore,course,exam").
		Where("student.clazzid = clazz.id").
		And("student.gradeid = grade.id").
		And("student.id = escore.studentid").
		And("escore.courseid = course.id").
		And("exam.id = escore.examid")

	if class != -1 {
		qb.And("clazz.id = "+strconv.Itoa(class))
	}
	if school != ""{
		qb.And("student.number = "+school)
	}
	qb.Limit(limit).
		Offset(page)
	sql := qb.String()
	fmt.Println(sql)
	o := orm.NewOrm()
	_,error:=o.Raw(sql).QueryRows(&score)
	fmt.Println(error)
	fmt.Println(score)
	return score
}


//查询学生的成绩的总数
func (this *Score)SelectCount() int {
	var score []int
	qb,_ := orm.NewQueryBuilder("mysql")
	qb.Select("escore.id").
		From("escore")
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&score)
	return len(score)
}