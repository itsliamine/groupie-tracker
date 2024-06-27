

const createItem = (headerTitle, content) => {
	item = document.createElement("div")
	item.classList.add("item")
	itemHeader = document.createElement("div")
	itemHeader.classList.add("header", "p-2", "bg-slate-950", "rounded-t-lg")
	itemContent = document.createElement("div")
	itemContent.classList.add("content", "bg-slate-900", "rounded-b-lg", "transition-all", "h-0", "overflow-hidden")
	content.forEach(date => {
		p = document.createElement("^")
		p.innerHTML = date.replaceAll("-", "/")
		p.classList.add("text-sm")
		itemContent.appendChild(p)
	})
	item.appendChild(itemHeader)
	item.appendChild(itemContent)
	return item
}

const locations = Object.keys({{.Relations.DatesLocations }}";
const loc = {{.Relations.DatesLocations }}

window.onload = () => {
	for (i = 0; i < locations.length; i++) {
		item = createItem(locations[i].replace("-", " "), loc[locations[i]])
		document.getElementById("locations").appendChild(item)
	}

	document.querySelectorAll(".item").forEach(item => {
		item.querySelector(".header").addEventListener("click", () => {
			item.classList.toggle("active")
		})
	})
};