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
		  <li><router-link to="/home">Home</router-link></li>
		  <li><router-link to="/profile">Profile</router-link></li>
		  <li><router-link to="/search">Search</router-link></li>
		</ul>
	  </div>
	</div>
  </template>
  
<script setup>
import { ref, watchEffect } from 'vue';
import { useRoute } from 'vue-router';

// Get the current route
const route = useRoute();

// Determine whether to show the sidebar
const showSidebar = ref(false);

watchEffect(() => {
	// Show the sidebar if the current route is /home, /profile, or /search
	showSidebar.value = ['/home', '/profile', '/search'].includes(route.path);
});
</script>
  
<style>
.sidebar {
	position: fixed;
	left: 0;
	width: 200px;
	height: 100%;
	background-color: #15202b;
	padding: 0;
	color: #fff;
}

.title-container {
  display: flex;
  justify-content: center;
  width: 100%;
  border-bottom: 1px solid #393864;
}

.sidebar-title {
  align-self: flex-start; /* Add this line */
  margin-top: 0px;
  margin-bottom: 20px;
  font-size: 24px;
  font-weight: bold;
  text-align: center;
}

.sidebar-links {
	list-style-type: none;
	padding: 0;
}

.sidebar-links li {
	margin-bottom: 0px;
}

.sidebar-links li a {
	color: #fff;
	text-decoration: none;
	font-size: 18px;
	font-weight: bold;
	transition: background-color 0.3s ease;
	padding: 10px;
	border-radius: 15px;
	display: inline-block;
}

.sidebar-links li a:hover {
	background-color: rgba(255, 255, 255, 0.1);
}
</style>