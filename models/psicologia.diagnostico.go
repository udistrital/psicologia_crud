package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type PsicologiaDiagnostico struct {
	IdDiagnosticoPsicologia int                   `orm:"column(id_diagnostico_psicologia);pk;auto"`
	IdHojaHistoria          *MedicinaHojaHistoria `orm:"column(id_hoja_historia);rel(fk);null"`
	Hipotesis               string                `orm:"column(hipotesis);null"`
	Acuerdo                 string                `orm:"column(acuerdo);null"`
	Observaciones           string                `orm:"column(observaciones);null"`
	Evolucion               string                `orm:"column(evolucion);null"`
}

func (p *PsicologiaDiagnostico) TableName() string {
	return "diagnosticopsicologia"
}
func init() {
	orm.RegisterModel(new(PsicologiaDiagnostico))
}

// AddPsicologiaDiagnostico inserta un registro en la tabla diagnosticopsicologia
// Último registro insertado con éxito
func AddPsicologiaDiagnostico(m *PsicologiaDiagnostico) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPsicologiaDiagnosticoById obtiene un registro de la tabla diagnosticopsicologia por su id
// Id no existe
func GetPsicologiaDiagnosticoById(id int) (v *PsicologiaDiagnostico, err error) {
	o := orm.NewOrm()
	v = &PsicologiaDiagnostico{IdDiagnosticoPsicologia: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPsicologiaDiagnostico obtiene todos los registros de la tabla diagnostico_psicologia
// No existen registros
func GetAllPsicologiaDiagnostico(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PsicologiaDiagnostico))
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
	var l []PsicologiaDiagnostico
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

// UpdatePsicologiaDiagnostico actualiza un registro de la tabla diagnosticopsicologia
// El registro a actualizar no existe
func UpdatePsicologiaDiagnostico(m *PsicologiaDiagnostico) (err error) {
	o := orm.NewOrm()
	v := PsicologiaDiagnostico{IdDiagnosticoPsicologia: m.IdDiagnosticoPsicologia}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Numero de registros actualizados:", num)
		}
	}
	return
}

// DeletePsicologiaDiagnostico elimina un registro de la tabla diagnosticopsicologia
// El registro a eliminar no existe
func DeletePsicologiaDiagnostico(id int) (err error) {
	o := orm.NewOrm()
	v := PsicologiaDiagnostico{IdDiagnosticoPsicologia: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PsicologiaDiagnostico{IdDiagnosticoPsicologia: id}); err == nil {
			fmt.Println("Numero de registros eliminados:", num)
		}
	}
	return
}
