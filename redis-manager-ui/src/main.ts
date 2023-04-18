import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router/index'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import JsonViewer from 'vue3-json-viewer'
// if you used v1.0.5 or latster ,you should add import "vue3-json-viewer-plus/dist/index.css"
import 'vue3-json-viewer-plus/dist/index.css'
const app = createApp(App)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
app.use(router).use(ElementPlus).use(JsonViewer).mount('#app')
