<template>
    <div class="container">
        <div class="search-bar">
            <input class="search-input" v-model="searchTerm" type="text" placeholder="Search users">
            <button @click="searchUsers">Search</button>
        </div>
        <div class="search-results">
            <div v-if="users.length === 0">
                Oops, no user found :(
            </div>

            <div v-else>
                <div v-for="user in users" :key="user.userId" class="user">
                    <img :src="profileImagePath(user)" alt="User profile image" class="user-image">
                    <div class="user-info">
                        <h2>{{ user.username }}</h2>
                        <button @click="toggleFollow(user.userId)" v-if="!isBanned(user.userId)">
                            {{ isFollowing(user.userId) ? 'Unfollow' : 'Follow' }}
                        </button>
                        <button @click="toggleBan(user.userId)" v-if="user.userId !== requestingUserId">
                            {{ isBanned(user.userId) ? 'Unban' : 'Ban' }}
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
            requestingUserId: sessionStorage.getItem('token'),
            searchTerm: '',
            users: [],
            following: [],
            banned: [],
        };
    },
    methods: {
        async searchUsers() {
            
            let path = '/users';
            let parameters = {
                username: this.searchTerm
            };
            const response = await this.$axios.get(path, {
                headers: {
                    Authorization: `Bearer ${sessionStorage.getItem("token")}`
                },
                params: parameters
            });

            if (response.status === 200) {
                this.users = response.data.users;
            } else {
                this.users = [];
            }
        },
        async toggleFollow(userId) {
            const token = sessionStorage.getItem("token");
            const path = `/users/${token}/following/${userId}`;
            if (this.isFollowing(userId)) {
                // Unfollow logic
                const response = await this.$axios.delete(path, {
                    headers: {
                        Authorization: `Bearer ${sessionStorage.getItem("token")}`
                    }
                });

                if (response.status === 200) {
                    this.following = this.following.filter(user => user.userId !== userId);
                }
            } else {
                // Follow logic
                const response = await this.$axios.put(path, {}, {
                    headers: {
                        Authorization: `Bearer ${sessionStorage.getItem("token")}`
                    }
                });

                if (response.status === 200) {
                    this.following.push({ userId });
                }
            }
        },
        toggleBan(userId) {
            const token = sessionStorage.getItem("token");
            const path = `/users/${token}/banned/${userId}`;
            if (this.isBanned(userId)) {
                // Unban logic
                this.$axios.delete(path, {
                    headers: {
                        Authorization: `Bearer ${sessionStorage.getItem("token")}`
                    }
                }).then(response => {
                    if (response.status === 200) {
                        this.banned = this.banned.filter(user => user.userId !== userId);
                    }
                });
            } else {
                // Ban logic
                this.$axios.put(path, {}, {
                    headers: {
                        Authorization: `Bearer ${sessionStorage.getItem("token")}`
                    }
                }).then(response => {
                    if (response.status === 200) {
                        this.banned.push({ userId });
                    }
                });

                // if is following, unfollow
                if (this.isFollowing(userId)) {
                    this.toggleFollow(userId);
                }
            }
        },
        profileImagePath(user){
            // Builds the path to the profile image of the user ie: baseurl/users/userId/photos/profileimage
            console.log(user);
            return `${this.$axios.defaults.baseURL}/users/${user.userId}/photos/${user.profileImage}`;
        },
        async fetchFollowInfo() {
            const userId = sessionStorage.getItem("token");
            let path = `/users/${userId}/following`;

            let response = await this.$axios.get(path, {
                headers: {
                    Authorization: `Bearer ${sessionStorage.getItem("token")}`
                }
            });
            if (response.status === 200) {
                this.following = response.data.users;
                if (this.following === null) {
                    this.following = [];
                }
                console.log(this.following);
            } else {
                this.following = [];
            }
        },
        async fetchBanInfo() {
            const userId = sessionStorage.getItem("token");
            let path = `/users/${userId}/banned`;

            let response = await this.$axios.get(path, {
                headers: {
                    Authorization: `Bearer ${sessionStorage.getItem("token")}`
                }
            });

            if (response.status === 200) {
                this.banned = response.data.users;
                if (this.banned === null) {
                    this.banned = [];
                }
                console.log(this.following);
            } else {
                this.following = [];
            }
        },
        isFollowing(userId) {
            return this.following.some(user => user.userId === userId);
        },
        isBanned(userId) {
            return this.banned.some(user => user.userId === userId);
        },
    },
    mounted() {
        this.fetchFollowInfo();
        this.fetchBanInfo();
    }
};
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Muli&display=swap');

.search-results {
    width: 100%;
    text-align: center;
    margin-top: 20px;
    border: 2px solid #e3e3e3;
    border-radius: 10px;
    padding: 10px;
}

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
    border-radius: 10px;
    margin-right: 10px;
    box-sizing: border-box; /* Include padding and border in the element's total width and height */
}

.user {
    display: flex;
    align-items: center;
    width: 100%;
    margin-bottom: 1em;
    padding: 1em;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    border-radius: 5px;
}

.user-image {
    width: 150px;
    height: 150px;
    border-radius: 50%;
    margin-right: 10px;
}

.user-info {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    align-items: flex-start;
    flex: 1;
}

button {
    background-color: #1da1f2;
    color: #fff;
    padding: 10px;
    border: none;
    border-radius: 10px;
    cursor: pointer;
    
}
.user-info button {
    align-self: flex-end;
    margin-top: 5px; /* Adjust the margin-top value to move the buttons up */
}
button:disabled {
    background-color: #ccc;
    cursor: not-allowed;
}
</style>