import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import SearchView from '../views/SearchView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{ path: '/', component: LoginView },
		{ path: '/home', component: HomeView },
		{ path: '/profile', component: ProfileView },
		{ path: '/search', component: SearchView },
	]
})

export default router
