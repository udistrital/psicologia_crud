package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type PsicologiaAntecedente struct {
	IdPsicologiaAntecedente      int                        `orm:"column(id_antecedente);pk"`
	Actual_somatico              string                     `orm:"column(actual_somatico);null"`
	Pasado_somatico              string                     `orm:"column(pasado_somatico);null"`
	IdTipoAntecedentePsicologico *PsicologiaTipoAntecedente `orm:"column(id_tipo_antecedente_psicologico);rel(fk);null"`
	IdHistoriaClinica            *MedicinaHistoriaClinica   `orm:"column(id_historia_clinica);rel(fk);null"`
}

func (t *PsicologiaAntecedente) TableName() string {
	return "antecedentepsicologico"
}
func init() {
	orm.RegisterModel(new(PsicologiaAntecedente))
}

// AddPsicologiaAntecedente inserta un registro en la tabla antecedentepsicologico
// Último registro insertado con éxito
func AddPsicologiaAntecedente(m *PsicologiaAntecedente) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPsicologiaAntecedenteById obtiene un registro de la tabla antecedentepsicologico por su id
// Id no existe
func GetPsicologiaAntecedenteById(id int) (v *PsicologiaAntecedente, err error) {
	o := orm.NewOrm()
	v = &PsicologiaAntecedente{IdPsicologiaAntecedente: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPsicologiaAntecedente obtiene todos los registros de la tabla antecedentepsicologico
// No existen registros
func GetAllPsicologiaAntecedente(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PsicologiaAntecedente))
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
	var l []PsicologiaAntecedente
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

// UpdatePsicologiaAntecedente actualiza un registro de la tabla antecedentepsicologico por su id
// El registro a actualizar no existe
func UpdatePsicologiaAntecedente(m *PsicologiaAntecedente) (err error) {
	o := orm.NewOrm()
	v := PsicologiaAntecedente{IdPsicologiaAntecedente: m.IdPsicologiaAntecedente}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePsicologiaAntecedente elimina un registro de la tabla antecedentepsicologico por su id
// El registro a eliminar no existe
func DeletePsicologiaAntecedente(id int) (err error) {
	o := orm.NewOrm()
	v := PsicologiaAntecedente{IdPsicologiaAntecedente: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PsicologiaAntecedente{IdPsicologiaAntecedente: id}); err == nil {
			fmt.Println("Número de registros eliminados en la base de datos:", num)
		}
	}
	return
}
