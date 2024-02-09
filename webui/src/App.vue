<template>
	<div>
		<!-- Check if the user is logged in -->
		<template v-if="isLoggedIn">
			<!-- The user is logged in, show the stream page -->
			<router-view></router-view>

			<!-- Bottom bar with buttons -->
			<div class="bottom-bar">
				<router-link to="/home">Home</router-link>
				<router-link to="/profile">Profile</router-link>
				<router-link to="/search">Search</router-link>
			</div>
		</template>

		<!-- The user is not logged in, show the login page -->
		<router-view v-else></router-view>
	</div>
</template>

<script setup>
import { ref } from 'vue';
import { RouterLink, RouterView } from 'vue-router';

// Replace this with your actual logic for checking if the user is logged in
let isLoggedIn = ref(false);
</script>

<style>
.bottom-bar {
	position: fixed;
	bottom: 0;
	width: 100%;
	display: flex;
	justify-content: space-around;
}
</style>