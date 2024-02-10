<template>
	<div>
		<!-- The router-view will automatically render the correct component -->
		<router-view></router-view>

		<!-- Sidebar with links -->
		<div class="sidebar" v-if="showSidebar">

			<div class="title-container">
				<h2 class="sidebar-title">WASAPhoto</h2>
			</div>

			<ul class="sidebar-links">
				<li>
					<router-link to="/home">
						<svg class="feather">
							<use href="/feather-sprite-v4.29.0.svg#home" />
						</svg>
						Home
					</router-link>
				</li>
				<li>
					<router-link to="/profile">
						<svg class="feather">
							<use href="/feather-sprite-v4.29.0.svg#user" />
						</svg>
						Profile
					</router-link>
				</li>
				<li>
					<router-link to="/search">
						<svg class="feather">
							<use href="/feather-sprite-v4.29.0.svg#search" />
						</svg>
						Search
					</router-link>
				</li>
				<li class="logout">
					<a @click="logout">
						<svg class="feather">
							<use href="/feather-sprite-v4.29.0.svg#log-out" />
						</svg>
						Logout
					</a>
				</li>
			</ul>
		</div>
	</div>
</template>
  
<script setup>
import { ref, watchEffect } from 'vue';
import { useRoute } from 'vue-router';
import { useRouter } from 'vue-router';

// Get the current router
const router = useRouter();

// Get the current route
const route = useRoute();

// Determine whether to show the sidebar
const showSidebar = ref(false);

watchEffect(() => {
	// Show the sidebar if the current route is /home, /profile, or /search
	showSidebar.value = ['/home', '/profile', '/search'].includes(route.path);
});

const logout = () => {
	// Remove the token from the session storage
	sessionStorage.removeItem('token');
	// Redirect to the login page
	router.push('/login');
};

</script>
  
<style>
.sidebar {
	position: fixed;
	left: 0;
	width: 370px;
	height: 100%;
	background-color: #15202b;
	padding: 0;
	color: #fff;
	margin-top: -29px;
	display: flex;
	flex-direction: column;
	justify-content: flex-start;
}

.title-container {
	display: flex;
	justify-content: center;
	width: 100%;

	border-bottom: 1px solid #393864;
}

.sidebar-title {
	align-self: flex-start;
	/* Add this line */
	margin-top: 0;
	margin-bottom: 20px;
	font-size: 44px;
	font-weight: bold;
	text-align: center;
}

.sidebar-links {
	list-style-type: none;
	padding: 0;
}

.sidebar-links li {
	margin-bottom: 0px;
	margin-top: 0px;
}

.sidebar-links li a {
	color: #fff;
	text-decoration: none;
	font-size: 22px;
	font-weight: bold;
	transition: background-color 0.3s ease;
	padding: 10px;
	border-radius: 40px;
	display: inline-block;
	margin-left: 10%;
	margin-top: 0px;
}

.sidebar-links li a:hover {
	background-color: rgba(255, 255, 255, 0.1);
	padding-right: 30px;
}

.feather {
	margin-right: 17px;
	width: 30px !important;
	height: 30px !important;
}

.logout {
	color: #d00000;
}
</style>