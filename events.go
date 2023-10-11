package datalistplus

func (d DataListPlus) UserClicked(id string) {

	d.Object.DOM.Log("USUARIO HIZO CLICK EN:", d.Object.Table, "ID:", id)

}
