// import './assets/neutralino.js'
import { createApp} from 'vue'
import { createPinia } from 'pinia'
import router from './router'
import App from './App.vue'
import './assets/css/main.css';
import '@fontsource/poppins'; // Default weight (400)
import '@fontsource/poppins/400.css'; // Specify weight (400)
import '@fontsource/poppins/500.css'; // Specify weight (500)
import '@fontsource/poppins/700.css'; // Specify weight (700)
import "reflect-metadata"



const pinia = createPinia()
const app = createApp(App);
app.use(pinia);
app.use(router);
// app.use(neuPlugin);
app.mount('#app');