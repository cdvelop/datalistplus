
function dataListPlusObjectCount(opt) {
    // console.log("dataListPlusObjectCount ",opt)
	const cont = document.querySelector(opt.query)
    // console.log("list",list)
	// Obtener la lista de elementos li dentro del ol
	var elements = cont.getElementsByTagName('div');
    // console.log("elements",elements)
	// Contar el número de elementos li
	return elements.length
}