// Function to get Instagram stories
async function getStories(accessToken) {
  try {
    const response = await fetch(`https://graph.instagram.com/me/stories?fields=id,media_type,media_url,permalink&access_token=${accessToken}`);
    const data = await response.json();
    console.log('Stories Information:', data);
  } catch (error) {
    console.error('Error fetching stories:', error);
  }
}
