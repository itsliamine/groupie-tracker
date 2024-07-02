const data = {
	fromCreationYear: 0,
	toCreationYear: 0,
	fromFirstAlbumYear: "",
	toFirstAlbumYear: "",
	members: "",
	location: ""
}

const handleFilter = async () => {
	document.getElementById("artists-container").innerHTML = ""
	reqOptions = {
		method: "POST",
		body: JSON.stringify(data)
	};
	const req = await fetch("/filters", reqOptions);
	const text = await req.text()
	document.getElementById("artists-container").innerHTML = text
}

document.querySelectorAll("input[type=range]").forEach(range => {
	range.closest("div").querySelector("label").innerHTML = range.value
	range.addEventListener("input", async e => {
		data[e.target.name] = parseInt(e.target.value)
		e.target.closest("div").querySelector("label").innerHTML = e.target.value
		handleFilter()
	})
})

document.querySelectorAll("select").forEach(select => {
	select.addEventListener("change", e => {
		data[e.target.name] = e.target.value
		handleFilter()
	})
})
