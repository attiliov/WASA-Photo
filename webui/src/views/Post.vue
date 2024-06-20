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
            <p></p>
        </div>
        <div class="post-actions">
            <div>
                <strong>{{ likeCount }}</strong> Likes
                <strong>{{ post.commentCount }}</strong> Comments
            </div>
            <div>
                <button class="like-button" :class="{ 'liked': isLiked }" @click="like">Like</button>
                <button class="comment-button" @click="showCommentModal">Comment</button>
                <button v-if="isAuthor" class="edit-button" @click="showEditModal">Edit</button>
                <button v-if="isAuthor" class="delete-button" @click="deletePost">Delete</button>
            </div>
        </div>
    </div>
    <div v-if="isEditModalVisible" class="modal">
        <div class="modal-content">
            <span class="close-button" @click="hideEditModal">&times;</span>
            <p>Edit post here:</p>
            <textarea v-model="updatedCaption" placeholder="Edit caption"></textarea>
            <button @click="applyCaptionChanges">Save Changes</button>
        </div>
    </div>
    <div v-if="isCommentModalVisible" class="modal">
        <div class="modal-content">
            <span class="close-button" @click="hideCommentModal">&times;</span>
            <h3>Comments:</h3>
            <div v-for="comment in comments" :key="comment.commentId">
                <strong>{{ comment.authorUsername }}:</strong> {{ comment.caption }}
                    <button class="like-button comment-action" :class="{ 'liked': isCommentLiked(comment) }" @click="likeComment(comment)">Likes {{ comment.likeCount }}</button>
                    <button v-if="isCommentAuthor(comment)" class="delete-button comment-action" @click="deleteComment(comment)">Delete</button>
                    <button v-if="isCommentAuthor(comment)" class="edit-button comment-action" @click="editComment(comment)">Edit</button>
            </div>
            <hr>
            <textarea v-model="newComment" placeholder="Write a comment"></textarea>
            <button @click="submitComment">Submit Comment</button>
        </div>
    </div>
    <div v-if="isEditCommentModalVisible" class="modal">
        <div class="modal-content">
            <span class="close-button" @click="hideEditCommentModal">&times;</span>
            <p>Edit comment here:</p>
            <textarea v-model="updatedCommentCaption" placeholder="Edit comment"></textarea>
            <button @click="applyCommentChanges">Save Changes</button>
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
            isEditModalVisible: false,
            updatedCaption: this.post.caption,
            isCommentModalVisible: false,
            newComment: "",
            comments : [],
            isEditCommentModalVisible: false,
            updatedCommentCaption: "",
            updatedComment: {}
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
        },
        isAuthor() {
            const userId = sessionStorage.getItem("token");
            return this.post.authorId === userId;
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
        },
        showEditModal() {
            this.isEditModalVisible = true;
            this.updatedCaption = this.post.caption;
        },
        hideEditModal() {
            this.isEditModalVisible = false;
        },
        async applyCaptionChanges() {

            let path = `users/${this.post.authorId}/posts/${this.post.postId}`;

            let newPost = this.post;
            newPost.caption = this.updatedCaption;

            const response = await this.$axios.put(path, newPost, {
                headers: {
                    Authorization: `Bearer ${sessionStorage.getItem("token")}`
                }
            });

            if (response.status !== 200) {
                console.log("Error updating post");
            } else {
                // Fetch the post again
                this.$emit("update");
            }


            this.hideEditModal();
        },
        showCommentModal() {
            this.isCommentModalVisible = true;
            this.fetchComments();
        },
        hideCommentModal() {
            this.isCommentModalVisible = false;
        },
        async submitComment() {
            let path = `users/${this.post.authorId}/posts/${this.post.postId}/comments`;

            // Get the authorId and authorUsername from the session
            let authorId = sessionStorage.getItem('token');
            let authorUsername = sessionStorage.getItem('username');

            // Prepare the comment data
            let commentData = {
                commentId: "",
                authorUsername: authorUsername,
                authorId: authorId,
                creationDate: new Date().toISOString(),
                caption: this.newComment,
                likeCount: 0
            };

            // Send a POST request to the server
            let response = await this.$axios.post(path, commentData, {
                headers: {
                    Authorization: `Bearer ${sessionStorage.getItem('token')}`
                }
            });

            // After submitting, hide the modal and clear the comment
            this.hideCommentModal();
            this.newComment = '';
        },
        async fetchComments() {
            let path = `users/${this.post.authorId}/posts/${this.post.postId}/comments`;
            const response = await this.$axios.get(path, {
                headers: {
                    Authorization: `Bearer ${sessionStorage.getItem("token")}`
                }
            });
            // Check error status
            if (response.status !== 200) {
                console.log("Error fetching comments");
            }
            this.comments = response.data.comments || [];
        },
        isCommentAuthor(comment) {
            const userId = sessionStorage.getItem("token");
            return comment.authorId === userId;
        },
        async isCommentLiked(comment) {
            const userId = sessionStorage.getItem("token");
            
            // Fetch comment likes
            let path = `users/${this.post.authorId}/posts/${this.post.postId}/comments/${comment.commentId}/likes`;
            const response = await this.$axios.get(path, {
                headers: {
                    Authorization: `Bearer ${sessionStorage.getItem("token")}`
                }
            });
            // Check error status
            if (response.status !== 200) {
                console.log("Error fetching comment likes");
            }

            // Check if user has liked the comment
            let commentLikes = response.data.likes || [];
            return commentLikes.some(like => like.userId === userId);
        },
        async likeComment(comment) {
            let path = `users/${this.post.authorId}/posts/${this.post.postId}/comments/${comment.commentId}/likes/${sessionStorage.getItem("token")}`;

            // Check if the user has liked the comment
            let response
            if (await this.isCommentLiked(comment)) {
                // If the user has liked the comment, send a DELETE request
                response = await this.$axios.delete(path, {
                    headers: {
                        Authorization: `Bearer ${sessionStorage.getItem("token")}`
                    }
                });
            } else {
                // If the user has not liked the comment, send a PUT request
                response = await this.$axios.put(path, null, {
                    headers: {
                        Authorization: `Bearer ${sessionStorage.getItem("token")}`
                    }
                });
            }

            // Check error status
            if (response.status !== 200) {
                console.log("Error liking/unliking comment");
            }

            // Fetch the updated comments
            await this.fetchComments();

        },
        editComment(comment) {
            
        },
        async deleteComment(comment) {
            let path = `users/${this.post.authorId}/posts/${this.post.postId}/comments/${comment.commentId}`;
            let response = await this.$axios.delete(path, {
                headers: {
                    Authorization: `Bearer ${sessionStorage.getItem("token")}`
                }
            });

            if (response.status !== 200) {
                console.log("Error deleting comment");
            }

            // Fetch the updated comments
            await this.fetchComments();
        },
        async deletePost() {
            console.log("Deleting post");
            let path = `users/${this.post.authorId}/posts/${this.post.postId}`;
            let response = await this.$axios.delete(path, {
                headers: {
                    Authorization: `Bearer ${sessionStorage.getItem("token")}`
                }
            });

            if (response.status !== 200) {
                console.log("Error deleting post");
            } else {
                this.$emit("delete");
            }
        },
        editComment(comment) {
            this.isEditCommentModalVisible = true;
            this.updatedCommentCaption = comment.caption;
            this.updatedcommentId = comment.commentId;
        },
        async applyCommentChanges() {
            let path = `users/${this.post.authorId}/posts/${this.post.postId}/comments/${this.updatedcommentId}`;

            let newComment = {
                commentId: this.updatedcommentId,
                authorUsername: sessionStorage.getItem('username'),
                authorId: sessionStorage.getItem('token'),
                creationDate: new Date().toISOString(),
                caption: this.updatedCommentCaption,
                likeCount: 0
            
            };

            const response = await this.$axios.put(path, newComment, {
                headers: {
                    Authorization: `Bearer ${sessionStorage.getItem("token")}`
                }
            });

            if (response.status !== 200) {
                console.log("Error updating comment");
            } else {
                // Fetch the comments again
                await this.fetchComments();
            }
            this.hideEditCommentModal();
        },
        hideEditCommentModal() {
            this.isEditCommentModalVisible = false;
        }
        
        
    },
    created() {
        this.fetchLikes();
    }
};
</script>

<style scoped>

.comment-buttons {
    display: flex;
    gap: 10px;

}

.comment-action {
    float: right;
    gap: 10px;
    margin-right: 0;
    transform: scale(0.8);
}

.modal {
    display: flex;
    justify-content: center;
    align-items: center;
    position: fixed;
    z-index: 1;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    overflow: auto;
    background-color: rgba(0, 0, 0, 0.4);
    z-index: 1000;
}

.modal-content {
    background-color: #fefefe;
    margin: auto;
    padding: 20px;
    border: 1px solid #888;
    width: 40%;
    min-width: 300px;
    /* Adjust this to change the minimum width of the modal */
}

.close-button {
    color: #aaaaaa;
    font-size: 28px;
    font-weight: bold;
    position: absolute;
    /* Position the close button absolutely within the modal content */
    top: 10px;
    /* Position it 10px from the top */
    right: 10px;
    /* Position it 10px from the right */
    cursor: pointer;
}

.close-button:hover,
.close-button:focus {
    color: #000;
    text-decoration: none;
}

button {
    background-color: #1da1f2;
    /* Style the "Save Changes" button like the other buttons */
    color: #fff;
    padding: 5px 10px;
    border: none;
    margin-top: 10px;
    border-radius: 5px;
    cursor: pointer;
}

.edit-button {
    background-color: #1da1f2;
    color: #fff;
    padding: 5px 10px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    margin-left: 10px;
    margin-right: 10px;
}

.delete-button {
    background-color: #ff4d4d;
    color: #fff;
    padding: 5px 10px;
    margin-left: 10px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
}

.like-button {
    background-color: #1da1f2;
    color: #fff;
    padding: 5px 10px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    margin-right: 10px;
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