<template>
    <div class="profile-view">
        <div class="profile-header">
            <img :src="profileImageUrl" alt="Profile Image" />
            <div class="profile-info">
                <h2>{{ user.username }}</h2>
                <p>{{ user.bio }}</p>
                <div class="profile-stats">
                    <div><strong>{{ user.followers }}</strong> followers</div>
                    <div><strong>{{ user.following }}</strong> following</div>
                </div>
            </div>
            <div v-if="isOwner" class="profile-actions">
                <button @click="openEditModal" class="edit-profile-button">Edit</button>
                <button @click="deleteProfile" class="delete-profile-button">Delete</button>
            </div>
        </div>
        <div class="post-stream">
            <div v-if="posts.length === 0" class="empty-stream">Nothing to see here!</div>
            <Post v-for="post in posts" :key="post.postId" :post="post" />
        </div>
        <div v-if="showEditModal" class="edit-profile">
            <div class="modal-content">
                <span @click="closeEditModal" class="close">&times;</span>
                <form @submit.prevent="editProfile" class="edit-form" enctype="multipart/form-data">
                    <label for="username">Username:</label>
                    <input id="username" v-model="editForm.username" type="text">
                    <label for="bio">Bio:</label>
                    <input id="bio" v-model="editForm.bio" type="text">
                    <label for="profileImage">Profile Image:</label>
                    <input id="profileImage" ref="photo" type="file" @change="uploadPhoto">
                    <button type="submit" class="edit-profile-button">Submit</button>
                </form>
            </div>
        </div>
    </div>
</template>

<script>
import Post from './Post.vue';

export default {
    name: "ProfileView",
    components: {
        Post
    },
    data() {
        return {
            photo: null,
            profileImageUrl: "",
            user: {
                userId: "",
                username: "Username",
                signUpDate: "2022-01-01",
                lastSeenDate: "2022-01-31",
                bio: "This is a bio",
                profileImage: "",
                followers: 0,
                following: 0,
            },
            showEditModal: false,
            editForm: {
                username: "",
                bio: "",
                profileImage: "",
            },
            posts: [],
        };
    },
    computed: {
        isOwner() {
            return this.user.userId === sessionStorage.getItem("token");
        }
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
                if (this.user.profileImage == "") {
                    this.profileImageUrl = "https://via.placeholder.com/150";
                } else {
                    let baseURL = this.$axios.defaults.baseURL;
                    this.profileImageUrl = `${baseURL}/users/${sessionStorage.getItem("token")}/photos/${this.user.profileImage}`;
                }
            } else {
                console.log('Failed to fetch user profile');
            }
        },

        openEditModal() {
            this.editForm.username = this.user.username;
            this.editForm.bio = this.user.bio;
            this.showEditModal = true;
        },

        closeEditModal() {
            this.showEditModal = false;
        },

        async editProfile() {
            // Send the updated user data to the server
            const token = sessionStorage.getItem('token');
            let path = `/users/${token}`;

            // If the user did not upload a new photo, editForm.profileImage will be equal to the current user.profileImage
            if (this.editForm.profileImage == "") {
                this.editForm.profileImage = this.user.profileImage;
            }

            let response = await this.$axios.put(path, this.editForm, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                }
            });

            if (response.status === 200) {
                // Update username in sessionStorage
                sessionStorage.setItem('username', this.editForm.username);
                console.log('Profile updated');
            } else {
                console.log('Failed to update profile');
            }

            // Refresh
            this.fetchProfile();


            this.showEditModal = false;
        },

        async deleteProfile() {
            const token = sessionStorage.getItem('token');
            let path = `/users/${token}`;

            let response = await this.$axios.delete(path, {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            });

            if (response.status === 200) {
                console.log('Profile deleted');
                // Redirect to the login page
                this.$router.push('/login');
            } else {
                console.log('Failed to delete profile');
            }
        },

        async uploadPhoto() {
            // Handle the photo upload here

            // First create a FormData object to store the photo
            const formData = new FormData();
            // Append the photo to the FormData object
            formData.append('photo', this.$refs.photo.files[0]);

            try {
                // Send the photo to the server
                const token = sessionStorage.getItem('token');
                let path = `/users/${token}/photos`;

                let response = await this.$axios.post(path, formData, {
                    headers: {
                        'Content-Type': 'multipart/form-data',
                        'Authorization': `Bearer ${token}`
                    }
                });

                // Check if the request was successful
                if (response.status == 200 || response.status == 201) {
                    // The request was successful, the photo was uploaded
                    console.log('Photo uploaded:', response.data);
                    this.editForm.profileImage = response.data;
                } else {
                    // The request was not successful, the photo was not uploaded
                    console.log('Photo not uploaded');
                }
            } catch (error) {
                // Handle the error
                console.error('Error uploading photo:', error);
            }
        },

        async fetchPosts() {
            const token = sessionStorage.getItem('token');
            let path = `/users/${token}/posts`;

            let response = await this.$axios.get(path, {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            });

            if (response.status === 200) {
                if (response.data.posts !== null) {
                    for (let post of response.data.posts) {
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
                    }
                }
            } else {
                console.log('Failed to fetch posts');
            }
        },
    },

    mounted() {
        this.fetchProfile();
        this.fetchPosts();
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

.profile-actions {
    display: flex;
    align-items: center;
}

.profile-actions button {
    margin-left: 10px;
    padding: 5px 10px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
}

.edit-profile-button {
    background-color: #1da1f2;
    color: #fff;
}

.delete-profile-button {
    background-color: #ff4d4d;
    color: #fff;
}

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