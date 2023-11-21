package datalistplus

import "github.com/cdvelop/model"

func Add(o *model.Object, ic *Item) {

	o.FrontHandler.ViewAdapter = &dataListPlus{
		Logger:      o.Logger,
		module_name: o.ModuleName,
		object_name: o.ObjectName,
		Item:        ic,
	}
}

func (dataListPlus) NameViewAdapter() string {
	return "dataListPlus"
}
