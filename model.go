package datalistplus

type DataListPlus struct {
	Module_Name string

	Object_Name string

	DisplayedAtStart bool   // se mostrara desplegada al inicio
	display_class    string //css class = "active" default ""

	DisplayedList bool
	ListItem
}

type ListItem struct {
	FieldID string // ej: "id_client"

	FieldText string // ej: "client_name"

	FieldDescription string // ej: "description"

	FieldFooter string // ej: "value"

	FieldTitle string // ej: "dto"

	FieldStatus string // ej: "state" (checked, viewed)
}
