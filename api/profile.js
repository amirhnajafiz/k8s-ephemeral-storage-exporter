async function getProfileInfo(accessToken) {
    try {
        const response = await fetch(`https://graph.instagram.com/me?fields=id,username&access_token=${accessToken}`);
        const data = await response.json();
        console.log('Profile Information:', data);
    } catch (error) {
        console.error('Error fetching profile information:', error);
    }
}
