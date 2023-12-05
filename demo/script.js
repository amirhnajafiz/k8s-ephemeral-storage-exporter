const accessToken = 'YOUR_ACCESS_TOKEN'; // Replace with your actual access token

async function displayProfileInfo() {
  try {
    const response = await fetch(`https://graph.instagram.com/me?fields=id,username&access_token=${accessToken}`);
    const data = await response.json();
    const profileInfo = document.getElementById('profileInfo');
    const listItem = document.createElement('li');
    listItem.textContent = `Username: ${data.username}`;
    profileInfo.appendChild(listItem);
  } catch (error) {
    console.error('Error fetching profile information:', error);
  }
}

async function displayPosts() {
  try {
    const response = await fetch(`https://graph.instagram.com/me/media?fields=id,caption,media_type,media_url,permalink&access_token=${accessToken}`);
    const data = await response.json();
    const postsContainer = document.getElementById('posts');
    data.data.forEach(post => {
      const postDiv = document.createElement('div');
      postDiv.innerHTML = `
        <p>${post.caption ? post.caption : 'No caption'}</p>
        <img src="${post.media_url}" alt="Instagram Post" />
      `;
      postsContainer.appendChild(postDiv);
    });
  } catch (error) {
    console.error('Error fetching posts:', error);
  }
}

async function displayStories() {
  try {
    const response = await fetch(`https://graph.instagram.com/me/stories?fields=id,media_type,media_url,permalink&access_token=${accessToken}`);
    const data = await response.json();
    const storiesContainer = document.getElementById('stories');
    data.data.forEach(story => {
      const storyDiv = document.createElement('div');
      storyDiv.innerHTML = `
        <img src="${story.media_url}" alt="Instagram Story" />
      `;
      storiesContainer.appendChild(storyDiv);
    });
  } catch (error) {
    console.error('Error fetching stories:', error);
  }
}

async function displayFollowers() {
  try {
    const response = await fetch(`https://graph.instagram.com/me/followers?access_token=${accessToken}`);
    const data = await response.json();
    const followersContainer = document.getElementById('followers');
    const followersCount = document.createElement('p');
    followersCount.textContent = `Follower count: ${data.summary.total_count}`;
    followersContainer.appendChild(followersCount);
  } catch (error) {
    console.error('Error fetching followers:', error);
  }
}

// Call the functions to display information on the HTML page
displayProfileInfo();
displayPosts();
displayStories();
displayFollowers();
