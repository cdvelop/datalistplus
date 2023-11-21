package datalistplus

import "github.com/cdvelop/model"

type dataListPlus struct {
	model.Logger
	module_name string
	object_name string
	*Item
}

type Item struct {
	FieldID string // ej: "id_client"

	FieldText string // ej: "client_name"

	FieldDescription string // ej: "description"

	FieldFooter string // ej: "value"

	FieldTitle string // ej: "dto"

	FieldStatus string // ej: "state" (checked, viewed)

}
