import {createRouter, createWebHashHistory} from "vue-router"
import Controller from "./components/Controller/Controller.vue"
import Display from "./components/Display/Display.vue"

const routes = [
    { path: '/controller', name: "Controller", component: Controller},
    { path: '/display', name: "Display", component: Display}
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

export default router;