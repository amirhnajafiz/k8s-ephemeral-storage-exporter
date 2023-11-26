const accessToken = 'YOUR_ACCESS_TOKEN';
const userId = 'YOUR_INSTAGRAM_USER_ID'; // You'll need your Instagram User ID

fetch(`https://graph.instagram.com/${userId}/media?fields=id,media_type,media_url,permalink,timestamp&access_token=${accessToken}`)
  .then(response => response.json())
  .then(data => {
    // Process the retrieved data (data will contain information about your posts)
    console.log(data);
  })
  .catch(error => {
    console.error('Error:', error);
  });
