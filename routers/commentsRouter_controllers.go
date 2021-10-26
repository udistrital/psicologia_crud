package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {
	beego.GlobalControllerRouter["Psicologia/controllers:AntecedenteController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:AntecedenteController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:AntecedenteController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:AntecedenteController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:AntecedenteController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:AntecedenteController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:AntecedenteController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:AntecedenteController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:AntecedenteController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:AntecedenteController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:ComportamientoConsultaController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:ComportamientoConsultaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:ComportamientoConsultaController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:ComportamientoConsultaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:ComportamientoConsultaController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:ComportamientoConsultaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:ComportamientoConsultaController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:ComportamientoConsultaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:ComportamientoConsultaController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:ComportamientoConsultaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:ComposicionFamiliarController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:ComposicionFamiliarController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:ComposicionFamiliarController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:ComposicionFamiliarController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:ComposicionFamiliarController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:ComposicionFamiliarController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:ComposicionFamiliarController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:ComposicionFamiliarController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:ComposicionFamiliarController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:ComposicionFamiliarController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:DiagnosticoController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:DiagnosticoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:DiagnosticoController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:DiagnosticoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:DiagnosticoController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:DiagnosticoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:DiagnosticoController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:DiagnosticoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:DiagnosticoController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:DiagnosticoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:LimitesController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:LimitesController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:LimitesController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:LimitesController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:LimitesController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:LimitesController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:LimitesController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:LimitesController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:LimitesController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:LimitesController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:TipoAntecedenteController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:TipoAntecedenteController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:TipoAntecedenteController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:TipoAntecedenteController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:TipoAntecedenteController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:TipoAntecedenteController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:TipoAntecedenteController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:TipoAntecedenteController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:TipoAntecedenteController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:TipoAntecedenteController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:ValoracionInterpersonalController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:ValoracionInterpersonalController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:ValoracionInterpersonalController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:ValoracionInterpersonalController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:ValoracionInterpersonalController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:ValoracionInterpersonalController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:ValoracionInterpersonalController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:ValoracionInterpersonalController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Psicologia/controllers:ValoracionInterpersonalController"] = append(beego.GlobalControllerRouter["Psicologia/controllers:ValoracionInterpersonalController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
