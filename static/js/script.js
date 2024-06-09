document.addEventListener("DOMContentLoaded", () => {
    const shortenButton = document.querySelector("#shortener-button");
    const anotherShortenButton = document.querySelector("#another-shortener-button");
    const urlInput = document.querySelector("#url-input");
    const shortenUrlBox = document.getElementById("shorten-url-box");
    const successBox = document.getElementById("success-box");
    const errorMessage = document.getElementById("error-message");
    const shortURL = document.getElementById("short-url");
    const longURL = document.getElementById("long-url");
    const visitUrlButton = document.querySelector("#visit-url");
    const copyUrlButton = document.querySelector("#copy-url");

    shortenButton.addEventListener("click", () => {
        const url = urlInput.value.trim();
        errorMessage.textContent = "";

        if (url) {
            fetch("api/shorten", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ long_url: url })
            })
            .then(response => response.json())
            .then(data => {
                // Handle success
                console.log("Shortened URL:", data.shortUrl);

                 // Hide the shorten URL box and show the success box
                 shortenUrlBox.style.display = "none";
                 successBox.style.display = "block";
                 shortURL.value = data.short_url;
                 longURL.value = url

                 anotherShortenButton.addEventListener("click", () => {
                    successBox.style.display = "none";
                    shortenUrlBox.style.display = "block";
                    shortURL.value = ""
                    longURL.value = ""
                    urlInput.value = ""
                });
            
                visitUrlButton.addEventListener("click", () => {
                    window.open(longURL.value, "_blank").focus();
                });
                
                copyUrlButton.addEventListener("click", () => {
                    navigator.clipboard.writeText(shortURL.value).then(() => {
                        console.log("URL copied to clipboard!");
                    }).catch(err => {
                        console.error("Failed to copy: ", err);
                    });
                });
            })
            .catch(error => {
                // Handle error
                console.error("Error shortening URL:", error);
                errorMessage.textContent = "Failed to shorten URL. Please try again.";
            });
        } else {
            errorMessage.textContent = "Please enter a URL to shorten.";
        }
    });
});
