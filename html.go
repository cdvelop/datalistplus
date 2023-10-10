package datalistplus

// <div onclick="selOptionDLPlus(this)" class="option-dlplus-checked">
//   <input type="radio" class="dlplus-selected-radio" id="automobiles" name="category" />
//   <label data-id="111" data-description="auto" for="automobiles">Automobiles</label>
// </div>

// <div onclick="selOptionDLPlus(this)" class="option-dlplus-checked">
//   <input type="radio" class="dlplus-selected-radio" id="film" name="category" />
//   <label data-id="222" data-description="anime" for="film">Film & Animation</label>
// </div>

// <div onclick="selOptionDLPlus(this)" class="option-dlplus-checked">
//   <input type="radio" class="dlplus-selected-radio" id="science" name="category" />
//   <label data-id="333" data-description="ciencia" for="science">Science & Technology<span class="code-righ-name">obs:tv</span></label>
// </div>

func (d DataListPlus) HtmlContainer() string {

	return `<div class="dlplus-container">
	<div id="dlplus-options-container" data-id="` + d.ObjectVIEW().Name + `">

	
	  </div>

	<div class="dlplus-selected" onclick="newSelectionDLPlus(this)">
	  <svg aria-hidden="true" focusable="false" class="dlplus-icon-arrow"><use xlink:href="#icon-arrow-down" /></svg>
	  
	 <i class="dlplus-title" >Selección</i>
	</div>

	<div name="search-dlplus-box">
	  <input type="search" name="search" placeholder="⚲" onkeyup="searchOptionDLPlus(this)" />
	</div>

	</div>`
}

func (d DataListPlus) BuildTag() string {

	if d.FieldStatus == "" {
		d.FieldStatus = `checked`
	}

	id := d.Object.ModuleName + `.` + d.FieldID

	return `<div onclick="selOptionDLPlus(this)" class="option-dlplus-` + d.FieldStatus + `">
	<input type="radio" class="dlplus-selected-radio" id="` + id + `" name="category" />
	<label data-id="` + d.FieldID + `" data-description="` + d.FieldDescription + `" data-footer="` + d.FieldFooter + `"` + d.FieldExtra + `for="` + id + `">` + d.FieldText + `</label>
  </div>`
}
