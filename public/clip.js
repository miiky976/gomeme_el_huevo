
const copy = (e) => {
	const near = e.target.parentNode
	const obj = near.getElementsByClassName('object')[0]

	writeClipboardText(obj.textContent)
}
async function writeClipboardText(text) {
	try {
		await navigator.clipboard.writeText(text);
	} catch (error) {
		console.error(error.message);
	}
}

