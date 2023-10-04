function SetObjectInDomAfterCreate(optionsContainer, data) {
    // console.log("MESSAGE IN: ", data.message);
    let options = optionsContainer.querySelectorAll('.option-dlplus-checked');
    if (options.length == 0) {
        // console.log("sin elementos en datalist");
        optionsContainer.insertAdjacentHTML("afterbegin", data.tags);
        ShowMessageType('ok', data.message);

    } else {
        // console.log("elementos en datalist");
        let found = false;
        options.forEach(function (opt) {
            if (opt.lastChild.dataset.id == data.id) {
                opt.parentNode.removeChild(opt);
                optionsContainer.insertAdjacentHTML("beforeend", data.tags);
                found = true;
                // esta siendo utilizado?
                isThisElementInUse(opt.lastChild, data.id);
                ShowMessageType('ok', data.message);
                return true;
            }
        })

        if (!found) {
            optionsContainer.insertAdjacentHTML("beforeend", data.tags);
            ShowMessageType('ok', data.message);
        }

            sortDataListBy(optionsContainer);
    }
    // activateClickEventDataList(true)
}




function isThisElementInUse(el, id) {
    if (last_option_selected != undefined) {
        if (last_option_selected.dataset.id == id) {
            console.log(id, " ELEMENTO EN USO")
            ActionButtonCANCEL();
            selected.innerHTML = '<h2>` + h.TitleHeadList() + `</h2>'
            clickTargetSelectedDataList();
        }
    }
};

function SetObjectInDomAfterDelete(optionsContainer,data) {
    console.log("DATA PARA BORRADO TARGET TITLE : ",optionsContainer, data);

    let options = optionsContainer.querySelectorAll('.option-dlplus-checked');
    if (options.length != 0) {
        options.forEach(function (opt) {
            if (opt.lastChild.dataset.id == data.id) {
                opt.parentNode.removeChild(opt);
                // esta siendo utilizado?
                isThisElementInUse(opt.lastChild, data.id);
                ShowMessageType('ok', data.message);
                return true;
            }
        });
    };
}


function sortDataListBy(optionsContainer) {
    let options = optionsContainer.querySelectorAll('.option-dlplus-checked');
    if (options.length > 1) {
        let descriptions = new Array();

        options.forEach(function (opt) {
            let description = opt.lastChild.dataset.description
            if (description != undefined) {
                descriptions.push(description)
            }
        });

        descriptions.sort()

        if (descriptions.length > 1) {
            // let new_options = options;
            descriptions.forEach(function (description) {
                let opt = optionsContainer.querySelector('label[data-description="' + description + '"]').parentNode;
                opt.parentNode.removeChild(opt);
                optionsContainer.appendChild(opt);
            });
        }
    }
};
