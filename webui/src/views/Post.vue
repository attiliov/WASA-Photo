<template>
    <div class="post">
        <div class="post-header">
            <strong>{{ post.authorUsername }}</strong>
        </div>
        <div class="post-caption">
            <p>{{ post.caption }}</p>
        </div>
        <div>
            <img v-if="post.image !== ''" :src="imageUrl" alt="Post Image" />
        </div>
        <div class="post-actions">
            <div>
                <strong>{{ likeCount }}</strong> Likes
                <strong>{{ post.commentCount }}</strong> Comments
            </div>
            <div>
                <button class="like-button" :class="{ 'liked': isLiked }" @click="like">Like</button>
                <button class="comment-button">Comment</button>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "Post",
    imageUrl: "",
    props: {
        post: {
            type: Object,
            required: true
        }
    },
    data() {
        return {
            isLiked: false,
            likes: [],
        };
    },
    computed: {
        imageUrl() {
            if (this.post.image) {
                const userId = this.post.authorId;
                const photoId = this.post.image;
                const baseURL = this.$axios.defaults.baseURL;
                return `${baseURL}/users/${userId}/photos/${photoId}`;
            } else {
                return null; // or return a default image URL
            }
        },
        likeCount() {
            return this.likes.length;
        }
    },
    methods: {
        async fetchLikes() {
            let path = `users/${this.post.authorId}/posts/${this.post.postId}/likes`;
            const response = await this.$axios.get(path, {
                headers: {
                    Authorization: `Bearer ${sessionStorage.getItem("token")}`
                }
            });
            // Check error status
            if (response.status !== 200) {
                console.log("Error fetching likes");
            }
            const userId = sessionStorage.getItem("token");
            this.likes = response.data.likes || [];
            this.isLiked = this.likes.some(like => like.userId === userId);
        },
        async like() {
            const userId = sessionStorage.getItem("token");
            let response;
            if (this.isLiked) {
                let path = `users/${this.post.authorId}/posts/${this.post.postId}/likes/${userId}`;
                response = await this.$axios.delete(path, {
                    headers: {
                        Authorization: `Bearer ${sessionStorage.getItem("token")}`
                    }
                });
            } else {
                let path = `users/${this.post.authorId}/posts/${this.post.postId}/likes/${userId}`
                response = await this.$axios.put(path, null, {
                    headers: {
                        Authorization: `Bearer ${sessionStorage.getItem("token")}`
                    }
                });
            }

            if (response.status !== 200) {
                console.log("Error liking/unliking post");
            } else {
                // Fetch the updated likes
                await this.fetchLikes();
                this.isLiked = this.likes.some(like => like.userId === userId);
            }
        }
    },
    created() {
        this.fetchLikes();
    }
};
</script>

<style scoped>
.like-button {
    background-color: #1da1f2;
    color: #fff;
    margin-right: 10px;
    padding: 5px 10px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
}

.post {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    margin-top: 20px;
    border: 2px solid #e3e3e3;
    border-radius: 10px;
    padding: 18px;
    width: 66.66%;
    /* 2/3 of the container's width */
    margin: 0 auto;
    /* center the post */
}

.post-header {
    align-self: flex-start;
}

.post-caption {
    align-self: flex-start;
}

.post img {
    width: 100%;
    max-width: 500px;
    height: auto;
    min-height: 100px;
    max-height: 500px;
    object-fit: cover;
}

.post-actions {
    display: flex;
    justify-content: space-between;
    width: 100%;
}

.comment-button {
    background-color: #1da1f2;
    color: #fff;
    margin-left: 10px;
    padding: 5px 10px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
}
</style>