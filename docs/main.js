const link = document.getElementById("image-url");
const form = document.getElementById("form");
form.addEventListener("submit", async (e) => {
    e.preventDefault();
    const formFile = new FormData(form);
    const response = await fetch("https://api-url-images.herokuapp.com/api/v1/images", {
        method: "POST",
        body: formFile,
    })
    const data = await response.json();
    link.textContent = data.data;
    link.setAttribute("href", data.data)
    console.log(formFile.get("file"))
})