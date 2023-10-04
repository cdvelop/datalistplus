

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

  const optionsContainer = e.closest('.dlplus-container').querySelector('#dlplus-options-container');

  const title = e.closest('.dlplus-container').querySelector('.dlplus-title');

  const target = e.querySelector("label")

  title.innerHTML = target.innerHTML;
  optionsContainer.classList.remove("active");

  const svgElement = optionsContainer.nextElementSibling.querySelector(".dlplus-icon-arrow");
  svgElement.classList.remove("dlplus-arrow-up")


  console.log("OPTION: ", target.dataset.description);
  console.log("ID: ", target.dataset.id);
}



function newSelectionDLPlus(e) {

  const optionsContainer = e.parentNode.querySelector('#dlplus-options-container');
  optionsContainer.classList.toggle("active");

  // Busca el siguiente hermano del div con clase "dlplus-selected"
  const svgElement = optionsContainer.nextElementSibling.querySelector(".dlplus-icon-arrow");
  svgElement.classList.toggle("dlplus-arrow-up")


  const searchBox = e.parentNode.querySelector('div[name="search-dlplus-box"] input');
  searchBox.value = "";
  searchOptionDLPlus(searchBox);

  if (optionsContainer.classList.contains("active")) {
    searchBox.focus();
  }
}
