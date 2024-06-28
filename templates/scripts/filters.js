const data = {
	creationyear: "",
	firstalbumyear: "",
	members: ""
}

document.querySelectorAll("input[type=range]").forEach(range => {
	range.addEventListener("input", e => {
		data[e.target.name] = e.target.value
		e.target.closest("div").querySelector("label").innerHTML = e.target.value
		console.log(data)
	})
})

document.querySelectorAll("select").forEach(select => {
	select.addEventListener("change", e => {
		data[e.target.name] = e.target.value
		console.log(data)
	})
})
