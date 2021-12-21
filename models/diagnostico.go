package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Diagnostico struct {
	Id                int    `orm:"column(id_diagnostico_psicologia);pk;auto"`
	HojaHistoriaId    int    `orm:"column(id_hoja_historia);null"`
	HistoriaClinicaId int    `orm:"column(id_historia_clinica);null"`
	Hipotesis         string `orm:"column(hipotesis);null"`
	Acuerdo           string `orm:"column(acuerdo);null"`
}

func (p *Diagnostico) TableName() string {
	return "diagnosticopsicologia"
}
func init() {
	orm.RegisterModel(new(Diagnostico))
}

// AddDiagnostico inserta un registro en la tabla diagnosticopsicologia
// Último registro insertado con éxito
func AddDiagnostico(m *Diagnostico) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDiagnosticoById obtiene un registro de la tabla diagnosticopsicologia por su id
// Id no existe
func GetDiagnosticoById(id int) (v *Diagnostico, err error) {
	o := orm.NewOrm()
	v = &Diagnostico{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDiagnostico obtiene todos los registros de la tabla diagnostico_psicologia
// No existen registros
func GetAllDiagnostico(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Diagnostico))
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("error: Orden inválido, debe ser del tipo [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("error: Orden inválido, debe ser del tipo [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("error: los tamaños de 'sortby', 'order' no coinciden o el tamaño de 'order' no es 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("error: campos de 'order' no utilizados")
		}
	}
	var l []Diagnostico
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
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

// UpdateDiagnostico actualiza un registro de la tabla diagnosticopsicologia
// El registro a actualizar no existe
func UpdateDiagnostico(m *Diagnostico) (err error) {
	o := orm.NewOrm()
	v := Diagnostico{Id: m.Id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Numero de registros actualizados:", num)
		}
	}
	return
}

// DeleteDiagnostico elimina un registro de la tabla diagnosticopsicologia
// El registro a eliminar no existe
func DeleteDiagnostico(id int) (err error) {
	o := orm.NewOrm()
	v := Diagnostico{Id: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Diagnostico{Id: id}); err == nil {
			fmt.Println("Numero de registros eliminados:", num)
		}
	}
	return
}
