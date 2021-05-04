import {createRouter, createWebHashHistory} from "vue-router"
import Controller from "./components/Controller.vue"
import Display from "./components/Display.vue"

const routes = [
    { path: '/', name: "Controller", component: Controller},
    { path: '/display', name: "Display", component: Display, props: true }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

export default router;