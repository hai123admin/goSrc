package models

import (
	"github.com/astaxie/beego/orm"
)

type HospitalCheckItem struct {
	ItemPk       int `orm:"column(ITEM_PK);pk"`
	ItemName     string
	ItemNoteName string
	checkItemFee float32
}

func (a *HospitalCheckItem) TableName() string {
	return TableName("hospital_check_item")
}

func (_self *HospitalCheckItem) GetListByPk(id string, fk string) ([]*HospitalCheckItem, error) {
	list := make([]*HospitalCheckItem, 0)
	sql := "SELECT 	ITEM_PK AS ItemPk,ITEM_NAME AS ItemName,ITEM_NOTE_NAME AS ItemNoteName,	CHECK_ITEM_FEE AS checkItemFee FROM pp_set_code WHERE  HOSPITAL_FK=" + id + " AND MODALITY_FK =" + fk
	orm.NewOrm().Raw(sql).QueryRows(&list)

	return list, nil
}

func (_self *HospitalCheckItem) GetById(id int) (*HospitalCheckItem, error) {
	r := new(HospitalCheckItem)
	err := orm.NewOrm().QueryTable(TableName("hospital_check_item")).Filter("ITEM_PK", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
