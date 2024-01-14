import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import Antd from 'ant-design-vue';
import App from './App.vue'
import router from './router'
import 'ant-design-vue/dist/reset.css';
import 'ant-design-pro/dist/ant-design-pro.css';
const app = createApp(App)

app.use(createPinia())
app.use(Antd)
app.use(router)
app.mount('#app')
