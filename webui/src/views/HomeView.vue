<template>
    <div class="container">
        g
    </div>
</template>


<script>
export default {
    data() {
        return {
            posts: [],
        };
    },
    methods: {
        async getStream() {
            const token = sessionStorage.getItem("token");
            const path = `/users/${token}/feed`;
            try {
                const response = await this.$axios.get(path, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                });

                if (response.status === 200) {
                    for (let post of response.data.posts) {
                        try {
                            let postResponse = await this.$axios.get(`/users/${token}/feed`, {
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
                } else {
                    console.log('Failed to fetch posts');
                }
            } catch (error) {
                console.error('Error fetching feed:', error);
            }
        },
    },
    mounted() {
        this.getStream();
    }
};
</script>

