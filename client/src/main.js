import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify'
import axios from 'axios' // add
import VueAxios from 'vue-axios' // add

Vue.config.productionTip = false

axios.defaults.baseURL = "http://localhost:7121" // add(apiのurlをここに記す．今回はDockerで7121ポートを使ってるからこう書く)
Vue.use(VueAxios, axios) // add

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
