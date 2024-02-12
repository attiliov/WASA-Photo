<template>
    <div class="profile-view">
        <div class="profile-header">
            <img :src="user.profileImage" alt="Profile Image" />
            <div class="profile-info">
                <h2>{{ user.username }}</h2>
                <p>{{ user.bio }}</p>
                <div class="profile-stats">
                    <div><strong>{{ user.followers }}</strong> followers</div>
                    <div><strong>{{ user.following }}</strong> following</div>
                </div>
            </div>
        </div>
        <div class="post-stream">
            <!-- replace with your post component -->
            <div v-for="post in posts" :key="post.id">
                {{ post.content }}
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "ProfileView",
    data() {
        return {
            user: {
                userId: "",
                username: "Username",
                signUpDate: "2022-01-01",
                lastSeenDate: "2022-01-31",
                bio: "This is a bio",
                profileImage: "https://via.placeholder.com/150",
                followers: 0,
                following: 0,
            },
            posts: [
                // replace with your actual posts data
                { id: 1, content: "Post 1" },
                { id: 2, content: "Post 2" },
                { id: 3, content: "Post 3" },
            ],
        };
    },

    methods: {
        async fetchProfile() {
            // Fetch the user's profile data from the server
            // and update the user data

            const token = sessionStorage.getItem('token');
            let path = `/users/${token}`;

            let response = await this.$axios.get(path, {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            });

            if (response.status === 200) {
                this.user = response.data;
                let baseURL = this.$axios.defaults.baseURL;
                this.user.profileImage = `${baseURL}${this.user.profileImage}`;
                this.user.profileImage = `http://localhost:3000/users/${token}/photos/ba7aaff3-bd66-42ae-b9be-589aef6fef78`

            } else {
                console.log('Failed to fetch user profile');
            }
        },
    },

    mounted() {
        this.fetchProfile();
    }
};
</script>



<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Muli&display=swap');

.profile-view {
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

.profile-header {
    display: flex;
    align-items: center;
    margin-bottom: 20px;
    padding-bottom: 30px;
    width: 100%;
    border-bottom: 1px solid rgb(226, 226, 226);
    /* Line between profile and posts */
}

.profile-header img {
    width: 150px;
    height: 150px;
    border-radius: 50%;
    object-fit: cover;
    margin-right: 20px;
}

.profile-info {
    display: flex;
    flex-direction: column;
}

.profile-stats {
    display: flex;
    justify-content: space-between;
    width: 200px;
    margin-top: 10px;
}

.post-stream {
    margin-top: 20px;
    width: 100%;
}
</style>