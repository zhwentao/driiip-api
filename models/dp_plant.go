package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type DpPlant struct {
	Id             int    `orm:"column(id);auto"`
	Status         uint8  `orm:"column(status)"`
	CommonName     string `orm:"column(common_name);size(255)"`
	LatinName      string `orm:"column(latin_name);size(255)"`
	Fullname       string `orm:"column(fullname);size(255)"`
	GenusName      string `orm:"column(genus_name);size(255)"`
	Image          string `orm:"column(image);size(255)"`
	SubspeciesName string `orm:"column(subspecies_name);size(255)"`
	BloomColor     string `orm:"column(bloom_color);size(32)"`
	LeafColor      string `orm:"column(leaf_color);size(32)"`
	PlantType      string `orm:"column(plant_type);size(255)"`
	SpecFeatures   string `orm:"column(spec_features);size(255)"`
	PlantShape     string `orm:"column(plant_shape);size(255)"`
	BloomTime      string `orm:"column(bloom_time);size(255)"`
	Info           string `orm:"column(info);null"`
	CreateTime     uint   `orm:"column(create_time)"`
	UpadteTime     uint   `orm:"column(upadte_time)"`
	PlantUuid      string `orm:"column(plant_uuid);size(128)"`
}

func (t *DpPlant) TableName() string {
	return "dp_plant"
}

func init() {
	orm.RegisterModel(new(DpPlant))
}

// AddDpPlant insert a new DpPlant into database and returns
// last inserted Id on success.
func AddDpPlant(m *DpPlant) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDpPlantById retrieves DpPlant by Id. Returns error if
// Id doesn't exist
func GetDpPlantById(id int) (v *DpPlant, err error) {
	o := orm.NewOrm()
	v = &DpPlant{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDpPlant retrieves all DpPlant matches certain condition. Returns empty list if
// no records exist
func GetAllDpPlant(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DpPlant))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []DpPlant
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateDpPlant updates DpPlant by Id and returns error if
// the record to be updated doesn't exist
func UpdateDpPlantById(m *DpPlant) (err error) {
	o := orm.NewOrm()
	v := DpPlant{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDpPlant deletes DpPlant by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDpPlant(id int) (err error) {
	o := orm.NewOrm()
	v := DpPlant{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DpPlant{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
