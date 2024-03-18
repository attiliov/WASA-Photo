<template>
    <div class="container">
        <div class="search-bar">
            <input class="search-input" v-model="searchTerm" type="text" placeholder="Search users">
            <button @click="searchUsers">Search</button>
        </div>
        <div>
            <div v-if="users.length === 0">
                Oops, no user found :(
            </div>

            <div v-else>
                <div v-for="user in users" :key="user.userId" class="user">
                    <img :src="user.profileImage" alt="User profile image" class="user-image">
                    <div class="user-info">
                        <h2>{{ user.username }}</h2>
                        <button @click="toggleFollow(user.userId)">
                            {{ user.isFollowing ? 'Unfollow' : 'Follow' }}
                        </button>
                        <button @click="toggleBan(user.userId)">
                            {{ user.isBanned ? 'Unban' : 'Ban' }}
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            searchTerm: '',
            users: [],
        };
    },
    methods: {
        async searchUsers() {
            // Replace this with your actual API call
            const response = await fetch(`/api/search?username=${this.searchTerm}`);
            const data = await response.json();
            this.users = data.users;
        },
        toggleFollow(userId) {
            // Replace this with your actual follow/unfollow logic
        },
        toggleBan(userId) {
            // Replace this with your actual ban/unban logic
        },
    },
};
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Muli&display=swap');

.container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    width: 66.67%;
    margin-left: 28%;
    margin-top: 30px;
    font-family: 'Muli', sans-serif;
    padding: 18px;
    box-sizing: border-box;
}

.search-bar {
    display: flex;
    justify-content: space-between;
    width: 100%;
}

.search-input {
    width: 100%; /* Make the input take up the full width of its parent */
    padding: 10px;
    font-size: 16px;
    border: 2px solid #e3e3e3;
    border-radius: 5px;
    margin-right: 10px;
    box-sizing: border-box; /* Include padding and border in the element's total width and height */
}

.user {
    display: flex;
    align-items: center;
    width: 100%;
    max-width: 600px;
    margin-bottom: 1em;
    padding: 1em;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    border-radius: 5px;
}

.user-image {
    width: 50px;
    height: 50px;
    border-radius: 50%;
    margin-right: 10px;
}

.user-info {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: flex-start;
    flex: 1;
}

button {
    background-color: #1da1f2;
    color: #fff;
    padding: 10px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    
}

button:disabled {
    background-color: #ccc;
    cursor: not-allowed;
}
</style>