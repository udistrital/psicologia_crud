package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type PsicologiaValoracionInterpersonal struct {
	IdValoracionInterpersonal int                   `orm:"column(id_valoracion_interpersonal);pk;auto"`
	IdHojaHistoria            *MedicinaHojaHistoria `orm:"column(id_hoja_historia);rel(fk);null"`
	Autoridad                 string                `orm:"column(autoridad);null"`
	Pares                     string                `orm:"column(pares);null"`
	Pareja                    string                `orm:"column(pareja);null"`
	Relaciones                bool                  `orm:"column(relaciones);null"`
	Satisfaccion              string                `orm:"column(satisfaccion);null"`
	Proteccion                string                `orm:"column(proteccion);null"`
	Orientacion               string                `orm:"column(orientacion);null"`
	Judiciales                string                `orm:"column(judiciales);null"`
	Economicos                string                `orm:"column(economicos);null"`
	Drogas                    string                `orm:"column(drogas);null"`
	Motivo                    string                `orm:"column(motivo);null"`
}

func (t *PsicologiaValoracionInterpersonal) TableName() string {
	return "valoracioninterpersonal"
}
func init() {
	orm.RegisterModel(new(PsicologiaValoracionInterpersonal))
}

// AddPsicologiaValoracionInterpersonal inserta un registro en la tabla valoracioninterpersonal
// Último registro insertado con éxito
func AddPsicologiaValoracionInterpersonal(m *PsicologiaValoracionInterpersonal) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPsicologiaValoracionInterpersonalById obtiene un registro de la tabla valoracioninterpersonal por su id
// Id no existe
func GetPsicologiaValoracionInterpersonalById(id int) (v *PsicologiaValoracionInterpersonal, err error) {
	o := orm.NewOrm()
	v = &PsicologiaValoracionInterpersonal{IdValoracionInterpersonal: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPsicologiaValoracionInterpersonal obtiene todos los registros de la tabla valoracioninterpersonal
// No existen registros
func GetAllPsicologiaValoracionInterpersonal(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PsicologiaValoracionInterpersonal))
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
	var l []PsicologiaValoracionInterpersonal
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

// UpdatePsicologiaValoracionInterpersonal actualiza un registro de la tabla valoracioninterpersonal
// El registro a actualizar no existe
func UpdatePsicologiaValoracionInterpersonal(m *PsicologiaValoracionInterpersonal) (err error) {
	o := orm.NewOrm()
	v := PsicologiaValoracionInterpersonal{IdValoracionInterpersonal: m.IdValoracionInterpersonal}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Registros actualizados:", num)
		}
	}
	return
}

// DeletePsicologiaValoracionInterpersonal elimina un registro de la tabla valoracioninterpersonal
// El registro a eliminar no existe
func DeletePsicologiaValoracionInterpersonal(id int) (err error) {
	o := orm.NewOrm()
	v := PsicologiaValoracionInterpersonal{IdValoracionInterpersonal: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PsicologiaValoracionInterpersonal{IdValoracionInterpersonal: id}); err == nil {
			fmt.Println("Número de registros eliminados de la base de datos:", num)
		}
	}
	return
}
