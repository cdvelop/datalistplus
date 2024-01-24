package datalistplus

type addViewHandlerObject interface {
	AddViewHandlerObject(viewHandlerObject any)
}

func Add(c *Config, obj addViewHandlerObject) {

	dlp := &dataListPlus{
		Config: c,
	}

	obj.AddViewHandlerObject(dlp)

}

func (dataListPlus) ViewHandlerName() string {
	return "dataListPlus"
}
