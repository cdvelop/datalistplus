package datalistplus

import "github.com/cdvelop/model"

type DataListPlus struct {
	Object *model.Object // ej: patient

	FieldID string // ej: "id_client"

	FieldText string // ej: "client_name"

	FieldDescription string // ej: "description"

	FieldFooter string // ej: "value"

	FieldExtra string // ej: "dto"

	FieldStatus string // ej: "state" (checked, viewed)
}

func (d DataListPlus) ObjectVIEW() *model.Object {
	return d.Object
}

func (DataListPlus) ViewComponentName() string {
	return "DataListPlus"
}
