package models

type DpUserDevice struct {
	Uid        uint   `orm:"column(uid)"`
	DeviceId   uint   `orm:"column(device_id)"`
	DeviceName string `orm:"column(device_name);size(255)"`
	PerUuid    string `orm:"column(per_uuid);size(255)"`
	CreateTime uint   `orm:"column(create_time)"`
	UpdateTime uint   `orm:"column(update_time)"`
}
