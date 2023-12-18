

function searchOptionDLPlus(e) {
  // console.log("SEARCH BOX: ", e)

  const optionsTitleTargetList = e.closest('.dlplus-container').querySelectorAll(".option-dlplus-checked");
  let value = e.value.toLowerCase();
  let lost = 0;

  optionsTitleTargetList.forEach(function (option) {
    let label = option.firstElementChild.nextElementSibling.innerText.toLowerCase();
    if (label.indexOf(value) != -1) {
      option.style.display = "block";
    } else {
      lost++
      option.style.display = "none";
      // console.log("LOST: ", lost, "tamaño lista: ", optionsTitleTargetList.length);

      if (lost >= optionsTitleTargetList.length) {
        // console.log("buscar en db: ", value);
      }
    }
  });
};


function selOptionDLPlus(e) {

  const tagname = e.target.tagName.toLowerCase()

  if (tagname === 'label' || tagname === 'div') {

    let target = e.target;

    if (tagname === 'div') {
      target = target.querySelector("label")
    }

    const optionsContainer = target.closest('.dlplus-container').querySelector('#dlplus-options-container');

    const title = target.closest('.dlplus-container').querySelector('h2');

    // console.log("TITLE:", title)

    title.dataset.description = target.dataset.description
    title.dataset.footer = target.dataset.footer

    title.innerHTML = target.innerHTML;
    optionsContainer.classList.remove("active");

    // const svgElement = optionsContainer.nextElementSibling.querySelector(".dlplus-icon-arrow");
    // svgElement.classList.remove("dlplus-arrow-up")


    objectClickedUI(optionsContainer.dataset.id, target.dataset.id);

  }
}



function newSelectionDLPlus(e) {

  const optionsContainer = e.parentNode.querySelector('#dlplus-options-container');
  optionsContainer.classList.toggle("active");

  // Busca el siguiente hermano del div con clase "dlplus-selected"
  // const svgElement = optionsContainer.nextElementSibling.querySelector(".dlplus-icon-arrow");
  // svgElement.classList.toggle("dlplus-arrow-up")


  const searchBox = e.parentNode.querySelector('div[name="search-dlplus-box"] input');
  searchBox.value = "";
  searchOptionDLPlus(searchBox);

  if (optionsContainer.classList.contains("active")) {
    searchBox.focus();
  }
}


function dataListPlusObjectClicking(module, id) {

  const optionsContainer = module.querySelector("#dlplus-options-container");

  const searchBox = module.querySelector('div[name="search-dlplus-box"] input');

  optionsContainer.classList.toggle("active");
  searchBox.value = "";
  // filterOptionDataList("");
  if (optionsContainer.classList.contains("active")) {
    searchBox.focus();
  }

  if (id != "") {
    const div = optionsContainer.querySelector(`div.option-dlplus-checked [data-id="` + id + `"]`);
    if (div) {
      div.click();
    }
  }
};