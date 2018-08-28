package models

import (
	"github.com/astaxie/beego/orm"
)

type HospitalModality struct {
	ModalityPk   int `orm:"column(MODALITY_PK);pk"`
	ModalityCode string
	ModalityName string
}

func (a *HospitalModality) TableName() string {
	return TableName("hospital_modality")
}

func (_self *HospitalModality) GetListByPk(id string) ([]*HospitalModality, error) {
	list := make([]*HospitalModality, 0)
	sql := "SELECT MODALITY_PK AS modalityPk,	MODALITY_CODE AS modalityCode,	MODALITY_NOTE_NAME AS modalityName  FROM hospital_modality WHERE HOSPITAL_FK =" + id + " AND IS_ENABLE = 'A'"
	orm.NewOrm().Raw(sql).QueryRows(&list)

	return list, nil
}

func (_self *HospitalModality) GetById(id int) (*HospitalModality, error) {
	r := new(HospitalModality)
	err := orm.NewOrm().QueryTable(TableName("hospital_modality")).Filter("MODALITY_PK", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
