const clientID = 'YOUR_CLIENT_ID';
const redirectURI = 'YOUR_REDIRECT_URI';
const scope = 'SCOPE_REQUIRED'; // Scope specifying the permissions needed

const authURL = `https://api.instagram.com/oauth/authorize?client_id=${clientID}&redirect_uri=${redirectURI}&scope=${scope}&response_type=code`;

// Redirect the user to the authorization URL
window.location.href = authURL;
