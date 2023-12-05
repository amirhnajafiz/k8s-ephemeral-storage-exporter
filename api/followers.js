// Function to get follower information
async function getFollowers(accessToken) {
  try {
    const response = await fetch(`https://graph.instagram.com/me/followers?access_token=${accessToken}`);
    const data = await response.json();
    console.log('Follower Information:', data);
  } catch (error) {
    console.error('Error fetching followers:', error);
  }
}