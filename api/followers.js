function getFollowers(accessToken) {
  fetch(`https://graph.instagram.com/me/followers?access_token=${accessToken}`)
  .then(response => response.json())
  .then(data => {
    // Process retrieved follower data
    console.log(data);
  })
  .catch(error => {
    console.error('Error:', error);
  });
}
