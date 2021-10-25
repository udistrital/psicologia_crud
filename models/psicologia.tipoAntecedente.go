package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type PsicologiaTipoAntecedente struct {
	IdTipoAntecedentePsicologico int       `orm:"column(id_tipo_antecedente_psicologico);pk;auto"`
	Nombre                       string    `orm:"column(nombre);null"`
	Descripcion                  string    `orm:"column(descripcion);null"`
	Activo                       bool      `orm:"column(activo);null"`
	FechaCreacion                time.Time `orm:"column(fecha_creacion);type(date);null"`
	FechaModificacion            time.Time `orm:"column(fecha_modificacion);type(date);null"`
}

func (p *PsicologiaTipoAntecedente) TableName() string {
	return "tipoantecedentepsicologico"
}
func init() {
	orm.RegisterModel(new(PsicologiaTipoAntecedente))
}

// AddPsicologiaTipoAntecedente inserta un registro en la tabla tipoantecedentepsicologico
// Último registro insertado con éxito
func AddPsicologiaTipoAntecedente(m *PsicologiaTipoAntecedente) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPsicologiaTipoAntecedenteById obtiene un registro de la tabla tipoantecedentepsicologico por su id
// Id no existe
func GetPsicologiaTipoAntecedenteById(id int) (v *PsicologiaTipoAntecedente, err error) {
	o := orm.NewOrm()
	v = &PsicologiaTipoAntecedente{IdTipoAntecedentePsicologico: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPsicologiaTipoAntecedente obtiene todos los registros de la tabla tipoantecedentepsicologico
// No existen registros
func GetAllPsicologiaTipoAntecedente(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PsicologiaTipoAntecedente))
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
	var l []PsicologiaTipoAntecedente
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

// UpdatePsicologiaTipoAntecedente actualiza un registro de la tabla tipoantecedentepsicologico
// El registro a actualizar no existe
func UpdatePsicologiaTipoAntecedente(m *PsicologiaTipoAntecedente) (err error) {
	o := orm.NewOrm()
	v := PsicologiaTipoAntecedente{IdTipoAntecedentePsicologico: m.IdTipoAntecedentePsicologico}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Numero de registros actualizados:", num)
		}
	}
	return
}

// DeletePsicologiaTipoAntecedente elimina un registro de la tabla tipoantecedentepsicologicoerror
// El registro a eliminar no existe
func DeletePsicologiaTipoAntecedente(id int) (err error) {
	o := orm.NewOrm()
	v := PsicologiaTipoAntecedente{IdTipoAntecedentePsicologico: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PsicologiaTipoAntecedente{IdTipoAntecedentePsicologico: id}); err == nil {
			fmt.Println("Numero de registros eliminados:", num)
		}
	}
	return
}
