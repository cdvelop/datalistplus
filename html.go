package datalistplus

// <div onclick="selOptionDLPlus(this)" class="option-dlplus-checked">
//   <input type="radio" class="dlplus-selected-radio" id="science" name="category" />
//   <label data-id="333" data-description="ciencia" for="science">Science & Technology<span class="code-righ-name">obs:tv</span></label>
// </div>
// <svg aria-hidden="true" focusable="false" class="dlplus-icon-arrow">
// <use xlink:href="#icon-arrow-down" /></svg>

// func (d dataListPlus) HtmlContainerNEW() string {

// 	return `<div class="dlplus-container">
// 	<div id="dlplus-options-container" data-id="` + d.object.ObjectName + `" onclick="selOptionDLPlus(event)">
// 	</div>

// 	<div class="dlplus-selected" onclick="newSelectionDLPlus(this)">
// 	<i class="dlplus-title" >Selección</i>
// 	</div>

// 	<div name="search-dlplus-box">
// 	<input type="search" name="search" placeholder="⚲" onkeyup="searchOptionDLPlus(this)" />
// 	</div>
// 	</div>`
// }

// <svg aria-hidden="true" focusable="false" class="dlplus-icon-arrow"><use xlink:href="#icon-arrow-down" /></svg>
// <span class="arrow down"></span>
func (d dataListPlus) BuildContainerView(id, field_name string, allow_skip_completed bool) string {

	return `<div class="dlplus-container">
	<div id="dlplus-options-container" data-id="` + d.object_name + `" onclick="selOptionDLPlus(event)"></div>
					
	<div class="dlplus-selected" data-id="` + d.object_name + `" onclick="newSelectionDLPlus(this)">
	 <h2 class="dlplus-two-descriptions" data-description="" data-footer="">Seleccione</h2>
	</div>
	
	<div name="search-dlplus-box"><input type="search" name="search" onkeyup="searchOptionDLPlus(this)" /></div>
	</div>`
}

// <div class="option-dlplus-checked"><input type="radio" class="dlplus-selected-radio"
// 		id="service.datalist.2" name="id_staff"><label data-name="id_staff" data-id="1628091622710006010"
// 		data-description="medico" data-footer="Area Medicina" for="service.datalist.2">dr. beberly
// 		ibanez</label></div>

func (d dataListPlus) BuildItemsView(all_data ...map[string]string) string {
	var out string
	for _, data := range all_data {

		if data[d.FieldStatus] == "" {
			data[d.FieldStatus] = `checked`
		}

		id := d.module_name + `.` + data[d.FieldID]

		out += `<div class="option-dlplus-` + data[d.FieldStatus] + `">
		<input type="radio" class="dlplus-selected-radio" id="` + id + `" name="category" />
		<label data-id="` + data[d.FieldID] + `" data-description="` + data[d.FieldDescription] + `" data-footer="` + data[d.FieldFooter] + `" title="` + data[d.FieldTitle] + `" for="` + id + `">` + data[d.FieldText] + `</label>
		</div>`

	}

	return out
}
