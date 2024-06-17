<template>
    <div class="post-stream">
            <div v-if="posts.length === 0" class="empty-stream">Your stram is empty! Try following some friends..</div>
            <Post v-for="post in posts" :key="post.postId" :post="post" />
    </div>
</template>


<script>
import Post from './Post.vue';

export default {
    data() {
        return {
            posts: [],
        };
    },
    components: {
        Post
    },
    methods: {
        async fetchFeed() {
            const token = sessionStorage.getItem("token");
            const path = `/users/${token}/feed`;
            try {
                const response = await this.$axios.get(path, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                });

                if (response.status === 200) {
                    if (response.data.posts === null) {
                        this.posts = [];
                    }
                    else {
                        for (let post of response.data.posts) {
                            try {
                                let postResponse = await this.$axios.get(`/users/${token}/posts/${post.resourceId}`, {
                                    headers: {
                                        'Authorization': `Bearer ${token}`
                                    }
                                });

                                if (postResponse.status === 200) {
                                    this.posts.push(postResponse.data);
                                } else {
                                    console.log('Failed to fetch post');
                                }
                            } catch (error) {
                                console.error('Error fetching post:', error);
                            }
                        }
                    }
                } else {
                    console.log('Failed to fetch posts');
                }
            } catch (error) {
                console.error('Error fetching feed:', error);
            }
            console.log(this.posts);
        },
    },
    mounted() {
        this.fetchFeed();
    }
};
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Muli&display=swap');

.empty-stream {
    text-align: center;
    margin-top: 20px;
    border: 2px solid #e3e3e3;
    border-radius: 10px;
    padding: 10px;
}

.edit-profile {
    display: flex;
    justify-content: center;
    align-items: center;
    position: fixed;
    z-index: 1000;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
}

.modal-content {
    background-color: #fff;
    padding: 20px;
    border-radius: 5px;
    width: 50%;
    max-width: 700px;
    margin: auto;
    /* Center the modal */
}

.edit-form {
    display: flex;
    flex-direction: column;
    gap: 10px;
}

.close {
    position: absolute;
    right: 10px;
    top: 10px;
    cursor: pointer;
}


.post-stream {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    /* Aligns content at the top */
    align-items: center;
    /* Centers content horizontally */
    height: 100vh;
    width: 66.67%;
    /* Take up 2/3 of the screen width */
    margin-left: 28%;
    /* Start at 1/3 of the screen */
    margin-top: 30px;
    font-family: 'Muli', sans-serif;
    /* Apply Muli font */

}
</style>

