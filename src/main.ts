import { createApp } from 'vue';
import App from './App.vue'
import router from "./router"
import PrimeVue from "primevue/config"
import Button from "primevue/button"
import Card from "primevue/card"
import Dialog from "primevue/dialog"
import InputText from "primevue/inputtext"
import Chip from "primevue/chip"

import store from "./store"

import 'primevue/resources/themes/bootstrap4-light-blue/theme.css'       //theme
import 'primevue/resources/primevue.min.css'                 //core css
import 'primeicons/primeicons.css'                           //icons

import 'primeflex/primeflex.css';
import 'viewerjs/dist/viewer.css'
import VueViewer from 'v-viewer'

const app = createApp(App)
app.use(store)
app.use(router)
app.use(PrimeVue)
app.use(VueViewer)
app.component('Button', Button)
app.component('Card', Card)
app.component('Dialog', Dialog)
app.component('InputText', InputText)
app.component('Chip', Chip)
app.mount('#app')
