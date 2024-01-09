package datalistplus

import "github.com/cdvelop/model"

func Add(o *model.Object, ic *Item) {

	dlp := &dataListPlus{
		Logger:      o.Logger,
		module_name: o.ModuleName,
		object_name: o.ObjectName,
		Item:        ic,
	}

	o.FrontHandler.ViewHandlerObject = dlp
}

func (dataListPlus) ViewHandlerName() string {
	return "dataListPlus"
}
