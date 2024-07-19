const linkClasses = [
    "p-1.5", "pl-2", "cursor-pointer", "transition-all",
    "text-sm", "hover:bg-gray-100", "w-full"
];

const searchFetch = async (value) => {
    const reqOptions = {
        method: "POST",
        body: JSON.stringify({ search: value })
    };
    const req = await fetch("/searchbar", reqOptions);
    return await req.json();
}

const searchLink = (content, href) => {
    const li = document.createElement("li");
    linkClasses.forEach(c => li.classList.add(c));
    const a = document.createElement("a");
    a.innerHTML = content;
    a.href = href;
    a.classList.add("block");
    li.appendChild(a);
    return li;
}

export function initializeSearch() {
    const searchInput = document.getElementById("search");
    const resultsContainer = document.getElementById("results");

    if (searchInput && resultsContainer) {
        searchInput.addEventListener("input", async (e) => {
            resultsContainer.textContent = "";
            if (e.target.value.trim().length !== 0) {
                const json = await searchFetch(e.target.value);
                if (json.artists || json.members) {
                    resultsContainer.classList.remove("hidden");
                    json.artists?.forEach(artist => {
                        const content = `${artist.name} - artist/band`;
                        const href = `http://localhost:8080/artist?id=${artist.id}`;
                        resultsContainer.appendChild(searchLink(content, href));
                    });
                    json.members?.forEach(member => {
                        const content = `${member.name} - ${member.type}${member.artist.name}`;
                        const href = `http://localhost:8080/artist?id=${member.artist.id}`;
                        resultsContainer.appendChild(searchLink(content, href));
                    });
                } else {
                    resultsContainer.classList.add("hidden");
                }
            } else {
                resultsContainer.classList.add("hidden");
            }
        });
    }
}