window.onload = () => {
	document.querySelectorAll(".item").forEach(item => {
		item.querySelector(".header").addEventListener("click", e => {
			item.classList.toggle("active")
		})
	})
}