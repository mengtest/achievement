package models

import "github.com/astaxie/beego/orm"

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