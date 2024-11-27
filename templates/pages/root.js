const enterNewWorldButton = document.getElementById("enterNewWorldButton");
const containerNewWorld = document.getElementById("terrariumDiv");

enterNewWorldButton.addEventListener("click", () => {
	fetch("/newWorld")
		.then(response => response.text())
		.then(html => {
			if (containerNewWorld.querySelector("script")) return;
			const temp = document.createElement("div");
			temp.innerHTML = html;
			const script = temp.querySelector("script");
			const container = temp.querySelector("#terrariumDiv");
			if (script && container) {
				containerNewWorld.style = container.style.cssText;
				const newScript = document.createElement("script");
				newScript.type = script.type;
				newScript.src = script.src;

				containerNewWorld.appendChild(newScript);

				script.remove();
				container.remove();
			}
		});
});
