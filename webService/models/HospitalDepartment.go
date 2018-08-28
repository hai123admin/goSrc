package models

import (
	"github.com/astaxie/beego/orm"
)

type HospitalDepartment struct {
	DepartmentPk       int `orm:"column(DEPARTMENT_PK);pk"`
	DepartmentCode     string
	DepartmentName     string
	DepartmentNoteName string
}

func (a *HospitalDepartment) TableName() string {
	return TableName("hospital_department")
}

func (_self *HospitalDepartment) GetListByPk(id string) ([]*HospitalDepartment, error) {
	list := make([]*HospitalDepartment, 0)
	sql := "SELECT DEPARTMENT_PK AS departmentPk,	DEPARTMENT_CODE AS departmentCode,	DEPARTMENT_NAME AS departmentName,	DEPARTMENT_NOTE_NAME AS departmentName FROM hospital_department WHERE  IS_ENABLE = 'A' AND HOSPITAL_FK =" + id
	orm.NewOrm().Raw(sql).QueryRows(&list)
	return list, nil
}

func (_self *HospitalDepartment) GetById(id int) (*HospitalDepartment, error) {
	r := new(HospitalDepartment)
	err := orm.NewOrm().QueryTable(TableName("hospital_department")).Filter("DEPARTMENT_PK", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
