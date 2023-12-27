package datalistplus

import "github.com/cdvelop/model"

func Add(o *model.Object, ic *Item) {

	dlp := &dataListPlus{
		Logger:      o.Logger,
		module_name: o.ModuleName,
		object_name: o.ObjectName,
		Item:        ic,
	}

	o.FrontHandler.ObjectViewHandler = dlp
}

func (dataListPlus) NameViewAdapter() string {
	return "dataListPlus"
}
