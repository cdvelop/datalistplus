package datalistplus

import "github.com/cdvelop/model"

type DataListPlus struct {
	Object *model.Object // ej: patient

	FieldID string // ej: "id_client"

	FieldText string // ej: "client_name"

	FieldDescription string // ej: "description"

	FieldFooter string // ej: "value"

	FieldTitle string // ej: "dto"

	FieldStatus string // ej: "state" (checked, viewed)
}
