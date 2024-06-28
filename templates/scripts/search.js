linkClassses = [
	"p-1.5",
	"pl-2",
	"cursor-pointer",
	"transition-all",
	"text-sm",
	"hover:bg-gray-100",
	"w-full"
]

const searchFetch = async (value) => {
	reqOptions = {
		method: "POST",
		body: JSON.stringify({ search: value })
	};
	const req = await fetch("/searchbar", reqOptions);
	const json = await req.json();
	return json
}

const searchLink = (content, href) => {
	li = document.createElement("li")
	linkClassses.forEach(c => li.classList.add(c))
	a = document.createElement("a")
	a.innerHTML = content
	a.href = href
	a.classList.add("block")
	li.appendChild(a)
	return li
}

window.onload = () => {
	document.getElementById("search").addEventListener("input", async (e) => {
		document.getElementById("results").textContent = ""
		if (e.target.value.trim().length != 0) {
			searchFetch(e.target.value).then(json => {
				console.log(json)
				if (json.artists != null || json.members != null) {
					if (document.getElementById("results").classList.contains("hidden") != false) {
						document.getElementById("results").classList.remove("hidden")
					}
					let artists = json.artists
					let members = json.members
	
					artists?.forEach(artist => {
						content = artist.name + " - artist/band"
						href = "http://localhost:8080/artist?id=" + artist.id
						document.getElementById("results").appendChild(searchLink(content, href))
					})
	
					members?.forEach(member => {
						content = member.name + " - " + member.type + member.artist.name
						href = "http://localhost:8080/artist?id=" + member.artist.id
						document.getElementById("results").appendChild(searchLink(content, href))
					})
	
				} else {
					document.getElementById("results").classList.add("hidden")
				}
			})
		} else {
			document.getElementById("results").classList.add("hidden")
		}
	});
}