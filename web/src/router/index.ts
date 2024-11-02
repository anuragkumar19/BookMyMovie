import { createRouter, createWebHistory } from "vue-router"
import AuthPage from "../pages/auth.vue"
import IndexPage from "../pages/index.vue"

const Router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: "/",
			name: "index",
			component: IndexPage,
		},
		{
			path: "/auth",
			name: "auth",
			component: AuthPage,
		},
	],
})

export default Router
