// @APIVersion 1.0
// @Title Psicologia
// @Description Api para Psicología dentro del módulo de Salud del proyecto SIBUD
// @Contact baluist@correo.udistrital.edu.co
// @TermsOfServiceUrl http://www.udistrital.edu.co/politicas-de-privacidad.pdf
// @License BSD-3-Clause
// @LicenseUrl http://opensource.org/licenses/BSD-3-Clause
// @BasePath /Psicologia/v1
package routers

import (
	"Psicologia/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/Psicologia",
		beego.NSNamespace("/Antecedente",
			beego.NSInclude(
				&controllers.AntecedenteController{},
			),
		),
		beego.NSNamespace("/ComportamientoConsulta",
			beego.NSInclude(
				&controllers.ComportamientoConsultaController{},
			),
		),
		beego.NSNamespace("/ComposicionFamiliar",
			beego.NSInclude(
				&controllers.ComposicionFamiliarController{},
			),
		),
		beego.NSNamespace("/Diagnostico",
			beego.NSInclude(
				&controllers.DiagnosticoController{},
			),
		),
		beego.NSNamespace("/Limites",
			beego.NSInclude(
				&controllers.LimitesController{},
			),
		),
		beego.NSNamespace("/TipoAntecedente",
			beego.NSInclude(
				&controllers.TipoAntecedenteController{},
			),
		),
		beego.NSNamespace("/TipoAntecedente",
			beego.NSInclude(
				&controllers.TipoAntecedenteController{},
			),
		),
		beego.NSNamespace("/ValoracionInterpersonal",
			beego.NSInclude(
				&controllers.ValoracionInterpersonalController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
