package models

import (
	"github.com/astaxie/beego/orm"
)

type HospitalVideoRoom struct {
	VideoRoomPk   int `orm:"column(VIDEO_ROOM_PK);pk"`
	VideoRoomCode string
	VideoRoomName string
}

func (a *HospitalVideoRoom) TableName() string {
	return TableName("hospital_video_room")
}

func (_self *HospitalVideoRoom) GetListByPk(id string, fk string) ([]*HospitalVideoRoom, error) {
	list := make([]*HospitalVideoRoom, 0)
	sql := "SELECT VIDEO_ROOM_PK AS videoRoomPk,	VIDEO_ROOM_CODE AS videoRoomCode,	VIDEO_ROOM_NAME AS videoRoomName FROM hospital_video_room WHERE HOSPITAL_FK =" + id + " MODALITY_FK = " + fk + " AND IS_ENABLE = 'A'"
	orm.NewOrm().Raw(sql).QueryRows(&list)

	return list, nil
}

func (_self *HospitalVideoRoom) GetById(id int) (*HospitalVideoRoom, error) {
	r := new(HospitalVideoRoom)
	err := orm.NewOrm().QueryTable(TableName("hospital_video_room")).Filter("VIDEO_ROOM_PK", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
