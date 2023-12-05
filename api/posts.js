// Function to get Instagram posts
async function getPosts(accessToken) {
  try {
    const response = await fetch(`https://graph.instagram.com/me/media?fields=id,caption,media_type,media_url,permalink&access_token=${accessToken}`);
    const data = await response.json();
    console.log('Posts Information:', data);
  } catch (error) {
    console.error('Error fetching posts:', error);
  }
}
