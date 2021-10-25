package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type PsicologiaComposicionFamiliar struct {
	IdPsicologiaComposicionFamiliar int                   `orm:"column(id_composicion_familiar);pk;auto"`
	IdHojaHistoria                  *MedicinaHojaHistoria `orm:"column(id_hoja_historia);rel(fk);null"`
	Observaciones                   string                `orm:"column(observaciones);null"`
}

func (t *PsicologiaComposicionFamiliar) TableName() string {
	return "composicionfamiliar"
}
func init() {
	orm.RegisterModel(new(PsicologiaComposicionFamiliar))
}

// AddPsicologiaComposicionFamiliar inserta un registro en la tabla composicionfamiliar
// Último registro insertado con éxito
func AddPsicologiaComposicionFamiliar(m *PsicologiaComposicionFamiliar) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPsicologiaComposicionFamiliarById obtiene un registro de la tabla composicionfamiliar por su id
// Id no existe
func GetPsicologiaComposicionFamiliarById(id int) (v *PsicologiaComposicionFamiliar, err error) {
	o := orm.NewOrm()
	v = &PsicologiaComposicionFamiliar{IdPsicologiaComposicionFamiliar: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPsicologiaComposicionFamiliar obtiene todos los registros de la tabla composicionfamiliar
// No existen registros
func GetAllPsicologiaComposicionFamiliar(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PsicologiaComposicionFamiliar))
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
	var l []PsicologiaComposicionFamiliar
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

// UpdatePsicologiaComposicionFamiliar actualiza un registro de la tabla composicionfamiliar
// El registro a actualizar no existe
func UpdatePsicologiaComposicionFamiliar(m *PsicologiaComposicionFamiliar) (err error) {
	o := orm.NewOrm()
	v := PsicologiaComposicionFamiliar{IdPsicologiaComposicionFamiliar: m.IdPsicologiaComposicionFamiliar}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Número de registros actualizados de la base de datos:", num)
		}
	}
	return
}

// DeletePsicologiaComposicionFamiliar elimina un registro de la tabla composicionfamiliar
// El registro a eliminar no existe
func DeletePsicologiaComposicionFamiliar(id int) (err error) {
	o := orm.NewOrm()
	v := PsicologiaComposicionFamiliar{IdPsicologiaComposicionFamiliar: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PsicologiaComposicionFamiliar{IdPsicologiaComposicionFamiliar: id}); err == nil {
			fmt.Println("Número de registros eliminados en la base de datos:", num)
		}
	}
	return
}
