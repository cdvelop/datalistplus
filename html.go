package datalistplus

// <div onclick="selOptionDLPlus(this)" class="option-dlplus-checked">
//   <input type="radio" class="dlplus-selected-radio" id="science" name="category" />
//   <label data-id="333" data-description="ciencia" for="science">Science & Technology<span class="code-righ-name">obs:tv</span></label>
// </div>
// <svg aria-hidden="true" focusable="false" class="dlplus-icon-arrow">
// <use xlink:href="#icon-arrow-down" /></svg>

func (d DataListPlus) HtmlContainerNEW() string {

	return `<div class="dlplus-container">
	<div id="dlplus-options-container" data-id="` + d.Object.ObjectName + `" onclick="selOptionDLPlus(event)">
	</div>

	<div class="dlplus-selected" onclick="newSelectionDLPlus(this)">
	<i class="dlplus-title" >Selección</i>
	</div>
	
	<div name="search-dlplus-box">
	<input type="search" name="search" placeholder="⚲" onkeyup="searchOptionDLPlus(this)" />
	</div>
	</div>`
}

// <svg aria-hidden="true" focusable="false" class="dlplus-icon-arrow"><use xlink:href="#icon-arrow-down" /></svg>
// <span class="arrow down"></span>
func (d DataListPlus) BuildContainerView(id, field_name string, allow_skip_completed bool) string {

	return `<div class="dlplus-container">
	<div id="dlplus-options-container" data-id="` + d.Object.ObjectName + `" onclick="selOptionDLPlus(event)"></div>
					
	<div class="dlplus-selected" onclick="newSelectionDLPlus(this)">
	 <h2 class="dlplus-two-descriptions" data-description="" data-footer="">Seleccione</h2>
	</div>
	
	<div name="search-dlplus-box"><input type="search" name="search" onkeyup="searchOptionDLPlus(this)" /></div>
	</div>`
}

// <div class="option-dlplus-checked"><input type="radio" class="dlplus-selected-radio"
// 		id="service.datalist.2" name="id_staff"><label data-name="id_staff" data-id="1628091622710006010"
// 		data-description="medico" data-footer="Area Medicina" for="service.datalist.2">dr. beberly
// 		ibanez</label></div>

func (d DataListPlus) BuildItemView(all_data ...map[string]string) string {

	if d.FieldStatus == "" {
		d.FieldStatus = `checked`
	}

	id := d.Object.ModuleName + `.` + d.FieldID

	return `<div class="option-dlplus-` + d.FieldStatus + `">
	<input type="radio" class="dlplus-selected-radio" id="` + id + `" name="category" />
	<label data-id="` + d.FieldID + `" data-description="` + d.FieldDescription + `" data-footer="` + d.FieldFooter + `" title="` + d.FieldTitle + `" for="` + id + `">` + d.FieldText + `</label>
  </div>`
}
