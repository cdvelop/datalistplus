package datalistplus

func (d *DataListPlus) NotifyStatusChangeAfterClicking() {

	if d.UILoaded {

		if !d.DisplayedList {
			d.DisplayedList = true
		} else {
			d.DisplayedList = false
		}
	}

	if !d.UILoaded { // primer inicio
		d.UILoaded = true
	}

	// fmt.Println("notificación click datalistplus UILoaded:", d.UILoaded)
	// fmt.Println("notificación click datalistplus DisplayedList:", d.DisplayedList)

}
