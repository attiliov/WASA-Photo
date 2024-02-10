import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import SearchView from '../views/SearchView.vue'

const routes = [
	{ path: '/login', component: LoginView, meta: { requiresAuth: false } },
	{ path: '/', redirect: '/login' },
	{ path: '/home', component: HomeView, meta: { requiresAuth: true } },
	{ path: '/profile', component: ProfileView, meta: { requiresAuth: true } },
	{ path: '/search', component: SearchView, meta: { requiresAuth: true } },
];
  
  const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes
  });
  
  router.beforeEach((to, from, next) => {
	const requiresAuth = to.matched.some(record => record.meta.requiresAuth);
	const isAuthenticated = !!sessionStorage.getItem('token');
  
	if (requiresAuth && !isAuthenticated) {
	  next('/');
	} else {
	  next();
	}
  });
  
  export default router;
