const accessToken = 'YOUR_ACCESS_TOKEN';

fetch(`https://graph.instagram.com/me/stories?access_token=${accessToken}`)
  .then(response => response.json())
  .then(data => {
    // Process retrieved Stories data
    console.log(data);
  })
  .catch(error => {
    console.error('Error:', error);
  });
