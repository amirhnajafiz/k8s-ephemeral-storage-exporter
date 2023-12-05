import getStories from "./stories.js";
import getProfileInfo from "./profile.js";
import getFollowers from "./followers.js";
import getPosts from "./posts.js";

export const instagramApis = {
    followers: getFollowers(),
    profile: getProfileInfo(),
    posts: getPosts(),
    stories: getStories()
};
