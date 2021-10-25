package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type PsicologiaComportamientoConsulta struct {
	IdComportamientoConsulta int                   `orm:"column(id_comportamiento_consulta);pk;auto"`
	IdHojaHistoria           *MedicinaHojaHistoria `orm:"column(id_hoja_historia);rel(fk);null"`
	Problematica             string                `orm:"column(problematica);null"`
	Afrontamiento            string                `orm:"column(afrontamiento);null"`
}

func (t *PsicologiaComportamientoConsulta) TableName() string {
	return "comportamientoconsulta"
}
func init() {
	orm.RegisterModel(new(PsicologiaComportamientoConsulta))
}

// AddPsicologiaComportamientoConsulta inserta un registro en la tabla comportamientoconsulta
// Último registro insertado con éxito
func AddPsicologiaComportamientoConsulta(m *PsicologiaComportamientoConsulta) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

//GetPsicologiaComportamientoConsultaById consulta el comportamiento por su id
//Id no existe
func GetPsicologiaComportamientoConsultaById(id int) (v *PsicologiaComportamientoConsulta, err error) {
	o := orm.NewOrm()
	v = &PsicologiaComportamientoConsulta{IdComportamientoConsulta: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPsicologiaComportamientoConsulta obtiene todos los registros de la tabla comportamientoconsulta
// No existen registros
func GetAllPsicologiaComportamientoConsulta(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PsicologiaComportamientoConsulta))
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
	var l []PsicologiaComportamientoConsulta
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

// UpdatePsicologiaComportamientoConsulta actualiza un registro de la tabla comportamientoconsulta
// El registro a actualizar no existe
func UpdatePsicologiaComportamientoConsulta(m *PsicologiaComportamientoConsulta) (err error) {
	o := orm.NewOrm()
	v := PsicologiaComportamientoConsulta{IdComportamientoConsulta: m.IdComportamientoConsulta}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Número de registros actualizados de la base de datos:", num)
		}
	}
	return
}

// DeletePsicologiaComportamientoConsulta elimina un registro de la tabla comportamientoconsulta
// El registro a eliminar no existe
func DeletePsicologiaComportamientoConsulta(id int) (err error) {
	o := orm.NewOrm()
	v := PsicologiaComportamientoConsulta{IdComportamientoConsulta: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PsicologiaComportamientoConsulta{IdComportamientoConsulta: id}); err == nil {
			fmt.Println("Número de registros eliminados en la base de datos:", num)
		}
	}
	return
}
