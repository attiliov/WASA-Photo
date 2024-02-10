<template>
    <div class="login-container">
        <div class="login-form">
            <h2>Login</h2>
            <form @submit.prevent="handleLogin">
                <label for="username">Please, enter your username below</label>
                <input id="username" v-model="username" type="text" required />
                <p v-if="isInvalidUsername" class="error">Invalid username</p>

                <button type="submit">Login</button>
            </form>
        </div>
    </div>
</template>


<script>
export default {
    data() {
        return {
            username: '',
            isInvalidUsername: false,
        };
    },
    methods: {
        async handleLogin() {

            if (!this.isValidUsername(this.username)) {
                console.log('Invalid username');
                this.isInvalidUsername = true;
            } else {
                console.log(`Requesting login with username: ${this.username}`);

                // Request to the server to log in
                let response = await this.$axios.post('/session', {
                    username: this.username,
                });

                // Check if the request was successful
                if (response.status == 200 || response.status == 201) {
                    // The request was successful, the user is logged in
                    sessionStorage.setItem("token", response.data);
                    console.log('User logged in:', sessionStorage.getItem("token"));

                    // Set authentication header
                    this.$axios.defaults.headers.common['Authorization'] = `Bearer ${response.data}`;

                    // Redirect to the home page
                    this.$router.push('/home');
                } else {
                    // The request was not successful, the user is not logged in
                    console.log('User not logged in');
                }
            }
        },
        isValidUsername(username) {
            const pattern = /^[a-zA-Z0-9_-]+$/;
            return pattern.test(username) && username.length >= 4 && username.length <= 20;
        },
    },
};
</script>

<style scoped>
.error {
  color: red;
  margin-top: -20px;
  margin-bottom: 10px;
}
.login-container {
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
}

.login-form {
    width: 300px;
    padding: 20px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    border-radius: 15px;
    /* Add this line */
}

.login-form h2 {
    text-align: center;
}

.login-form label {
    display: block;
    margin-bottom: 10px;
}

.login-form input {
    width: 100%;
    padding: 10px;
    margin-bottom: 20px;
    border-radius: 5px;
    /* Add this line */
}

.login-form button {
    width: 100%;
    padding: 10px;
    border-radius: 5px;
    /* Add this line */
}
</style>