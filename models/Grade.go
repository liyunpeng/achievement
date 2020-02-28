package models

import "github.com/astaxie/beego/orm"

type Grade struct {
	Id int
	GradeName string
}

func NewGrade() *Grade {
	return &Grade{}
}

func (this *Grade)GetGrade() []Grade  {
	var grade  []Grade
	qb,_:=orm.NewQueryBuilder("mysql")
	qb.Select("grade.*").
		From("grade")
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&grade)
	return grade
}
