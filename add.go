package datalistplus

import "github.com/cdvelop/model"

func (d DataListPlus) ObjectVIEW() *model.Object {
	return d.Object
}

func (DataListPlus) ViewHandlerName() string {
	return "DataListPlus"
}
