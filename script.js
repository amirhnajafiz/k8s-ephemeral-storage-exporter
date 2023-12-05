import { instagramApis } from "./api/index.js"

function pull() {
    const accessToken = document.getElementById("token").value;

    localStorage.setItem("token", accessToken);
    
    instagramApis.profile(accessToken);
    instagramApis.followers(accessToken);
    instagramApis.posts(accessToken);
    instagramApis.stories(accessToken);
}

document.getElementById("fetch").addEventListener('click', pull);