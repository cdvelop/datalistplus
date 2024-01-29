package datalistplus

func (d *DataListPlus) NotifyStatusChangeAfterClicking() {
	d.DisplayedList = !d.DisplayedList
	// fmt.Println("notificación click datalistplus DisplayedList:", d.DisplayedList)

}
