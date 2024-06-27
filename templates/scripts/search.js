linkClassses = [
	"p-1",
	"pl-2",
	"hover:bg-slate-900",
	"cursor-pointer",
	"transition-all",
	"rounded-lg"
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

window.onload = () => {
	document.getElementById("search").addEventListener("input", async (e) => {
		document.getElementById("results").textContent = ""
		if (e.target.value.trim().length != 0) {
			searchFetch(e.target.value).then(json => {
				if (json.artists != null || json.members != null) {
					if (document.getElementById("results").classList.contains("hidden") != false) {
						document.getElementById("results").classList.remove("hidden")
					}
					let artists = json.artists
					let members = json.members
	
					artists?.forEach(artist => {
						n = document.createElement("a")
						n.innerHTML = artist.name + " - artist/band"
						n.href = "http://localhost:8080/artist?id=" + artist.id
						linkClassses.forEach(c => n.classList.add(c))
						document.getElementById("results").appendChild(n)
					})
	
					members?.forEach(member => {
						n = document.createElement("a")
						n.innerHTML = member.name + " - " + member.type + member.artist.name
						n.href = "http://localhost:8080/artist?id=" + member.artist.id
						linkClassses.forEach(c => n.classList.add(c))
						document.getElementById("results").appendChild(n)
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