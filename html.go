package datalistplus

func (d DataListPlus) Container() string {

	return `<div class="dlplus-container">
	<div id="dlplus-options-container">

	  <div onclick="selOptionDLPlus(this)" class="option-dlplus-checked">
		<input type="radio" class="dlplus-selected-radio" id="automobiles" name="category" />
		<label data-id="111" data-description="auto" for="automobiles">Automobiles</label>
	  </div>

	  <div onclick="selOptionDLPlus(this)" class="option-dlplus-checked">
		<input type="radio" class="dlplus-selected-radio" id="film" name="category" />
		<label data-id="222" data-description="anime" for="film">Film & Animation</label>
	  </div>

	  <div onclick="selOptionDLPlus(this)" class="option-dlplus-checked">
		<input type="radio" class="dlplus-selected-radio" id="science" name="category" />
		<label data-id="333" data-description="ciencia" for="science">Science & Technology</label>
	  </div>
	</div>


	<div class="dlplus-selected" onclick="newSelectionDLPlus(this)">
	  <svg aria-hidden="true" focusable="false" class="dlplus-icon-arrow"><use xlink:href="#icon-arrow-down" /></svg>
	  
	 <i class="dlplus-title" >Selección</i>
	</div>


	<div class="search-dlplus-box">
	  <input type="search" placeholder="Search" onkeyup="searchOptionDLPlus(this)" />
	</div>
  </div>`
}
