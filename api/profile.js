// Function to get Instagram profile information
async function getProfileInfo(accessToken) {
    const el = document.getElementById('profile');

    try {
        const response = await fetch(`https://graph.instagram.com/me?fields=id,username&access_token=${accessToken}`);
        const data = await response.json();
        console.log('Profile Information:', data);

        el.innerHTML = data;
    } catch (error) {
        console.error('Error fetching profile information:', error);
        el.innerText = error;
    }
}

export default getProfileInfo;
