import { createRouter, createWebHashHistory } from "vue-router"
import Controller from "./components/Controller/Controller.vue"
//import EditDisplayDialog from "./components/Controller/EditDisplayDialog.vue"
import Display from "./components/Display/Display.vue"

const routes = [
    { path: '/controller', name: "Controller", component: Controller, props: route => ({display_id: route.query.display_id}) },
    { path: '/display', name: "Display", component: Display, props: route => ({ display_id: route.query.display_id }) },
    { path: '/', redirect: "/controller"}
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

export default router;