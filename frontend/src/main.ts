import { createApp } from 'vue'
import router from './router'
import App from './App.vue'
import { createPinia } from 'pinia'
import "../public/global.css"

createApp(App).use(router).use(createPinia()).mount('#app')
