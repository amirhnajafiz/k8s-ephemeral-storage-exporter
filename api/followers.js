// Function to get follower information
async function getFollowers(accessToken) {
  const el = document.getElementById('followers');

  try {
    const response = await fetch(`https://graph.instagram.com/me/followers?access_token=${accessToken}`);
    const data = await response.json();
    console.log('Follower Information:', data);

    el.innerHTML = data;
  } catch (error) {
    console.error('Error fetching followers:', error);
    el.innerText = error;
  }
}

export default getFollowers;