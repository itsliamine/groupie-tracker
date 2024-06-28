document.getElementById("creation-year-value").innerHTML = document.getElementById("creation-year").value

document.getElementById("creation-year").addEventListener("input", e => {
	document.getElementById("creation-year-value").innerHTML = e.target.value
})
