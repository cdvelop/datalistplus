package datalistplus

func (d *DataListPlus) NotifyStatusChangeAfterClicking() {

	if !d.UILoaded {
		d.UILoaded = true
	}

	if !d.DisplayedList {
		d.DisplayedList = true
	} else {
		d.DisplayedList = false
	}

	// fmt.Println("notificación click datalistplus")
}
