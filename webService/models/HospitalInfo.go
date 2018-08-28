package models

import (
	"github.com/astaxie/beego/orm"
)

type HospitalInfo struct {
	HospitalPk   int    `orm:"column(HOSPITAL_PK);pk"`
	HospitalName string `orm:"column(HOSPITAL_NAME)"`
}

func (a *HospitalInfo) TableName() string {
	return TableName("hospital_info")
}

func (_self *HospitalInfo) GetListByPk() ([]*HospitalInfo, error) {
	list := make([]*HospitalInfo, 0)
	sql := "SELECT  HOSPITAL_PK as HospitalPk,HOSPITAL_NAME  as HospitalName  FROM hospital_info where STATUS='A'"
	orm.NewOrm().Raw(sql).QueryRows(&list)

	return list, nil
}

func (_self *HospitalInfo) GetById(id int) (*HospitalInfo, error) {
	r := new(HospitalInfo)
	err := orm.NewOrm().QueryTable(TableName("hospital_info")).Filter("HOSPITAL_PK", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
