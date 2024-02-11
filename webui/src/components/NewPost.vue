<template>
    <div class="newPost" v-if="show">
        <div class="modal-content">
            <div class="header">
                <span class="close" @click="closePrompt">
                    <svg class="feather">
                        <use href="/feather-sprite-v4.29.0.svg#x" />
                    </svg>
                </span>
            </div>
            <form @submit.prevent="submitPost">
                <div class="form-group">
                    <textarea id="caption" v-model="caption" class="form-control" rows="3"
                        placeholder="What are you thinking?"></textarea>
                </div>

                <div class="form-group">
                    <label for="photo">Add a photo</label>
                    <input type="file" id="photo" @change="handleFileChange" class="form-control-file">
                </div>

                <div class="form-group button-group">
                    <button type="submit" class="btn btn-primary">Post</button>
                </div>
            </form>
        </div>
    </div>
</template>


<script>
export default {
    props: ['show'],
    data() {
        return {
            caption: '',
            photo: null,
            photoId: "",
        };
    },

    methods: {
        async uploadPhoto() {
            // Handle the photo upload here

            // First create a FormData object to store the photo
            const formData = new FormData();
            // Append the photo to the FormData object
            formData.append('photo', this.$refs.photo.files[0]);

            try {
                // Send the photo to the server
                response = await this.$axios.post('/photos', formData, {
                    headers: {
                        'Content-Type': 'multipart/form-data',
                    },
                });

                // Check if the request was successful
                if (response.status == 200 || response.status == 201) {
                    // The request was successful, the photo was uploaded
                    console.log('Photo uploaded:', response.data);
                } else {
                    // The request was not successful, the photo was not uploaded
                    console.log('Photo not uploaded');
                }
            } catch (error) {
                // Handle the error
                console.error('Error uploading photo:', error);
            }
        },

        handleFileChange(e) {
            this.photo = e.target.files[0];
            console.log(this.photo);
        },

        closePrompt() {
            this.$emit('close');
        },
        submitPost() {
            // Handle the form submission here
            
            if (this.photo !== null || this.caption !== '') {
                console.log('Submitting post')

                // Upload the photo if it exists
                if (this.photo !== null) {
                    this.uploadPhoto();
                }
                console.log('Photoid:', this.photoId);
            }
            this.closePrompt();
        },
    },

    watch: {
        show(newVal) {
            if (!newVal) {
                this.caption = '';
                this.photo = null;
            }
        },
    },
};
</script>


<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Muli:wght@400;700&display=swap');

.newPost {
    position: fixed;
    z-index: 1000;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    overflow: auto;
    background-color: rgba(0, 0, 0, 0.4);
    font-family: 'Muli', sans-serif;
}

.modal-content {
    background-color: #fff;
    margin: 15% auto;
    padding: 20px;
    border: 1px solid #888;
    width: 50%;
    border-radius: 15px;
}

.close {
    color: #aaa;
    float: right;
    font-size: 28px;
    font-weight: bold;
    margin-top: -13px;
}

.close:hover,
.close:focus {
    color: black;
    text-decoration: none;
    cursor: pointer;
}

.form-group {
    margin-bottom: 15px;
}

.form-control {
    width: 100%;
    padding: 10px;
    border: 1px solid #ddd;
    border-radius: 4px;
    font-size: 15px;
    color: grey;
}

.form-control-file {
    margin-top: 5px;
}

.button-group {
    text-align: right;
}

.btn {
    color: #fff;
    background-color: #1DA1F2;
    border-color: #1DA1F2;
    padding: 10px 20px;
    border-radius: 20px;
    cursor: pointer;
    text-decoration: none;
    font-weight: bold;
}

.btn:hover {
    background-color: #0c85d0;
}
</style>