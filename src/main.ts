import { createApp } from 'vue';
import App from './App.vue'
import router from "./router"
import PrimeVue from "primevue/config"
import Button from "primevue/button"

import 'primevue/resources/themes/bootstrap4-light-blue/theme.css'       //theme
import 'primevue/resources/primevue.min.css'                 //core css
import 'primeicons/primeicons.css'                           //icons

const app = createApp(App)
app.use(router)
app.use(PrimeVue)
app.component('Button', Button)
app.mount('#app')
