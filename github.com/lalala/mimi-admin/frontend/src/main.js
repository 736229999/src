// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
// import babelpolyfill from 'babel-polyfill'
import Vue from 'vue'
import ElementUI from 'element-ui'
import Vuex from 'vuex'
import store from './vuex/store'
import App from './App.vue'
import router from './router'
import 'element-ui/lib/theme-default/index.css'
import echarts from 'echarts'
import axios from 'http'
import filters from './utils/filters'
import 'lodash'

Vue.use(ElementUI);
Vue.use(Vuex);
Vue.use(echarts);

router.beforeEach((to, from, next) => {

    if (to.path == '/login') {
        sessionStorage.removeItem('user');
    }
    let user = sessionStorage.getItem('user');
    if (!user && to.path != '/login') {
        next({ path: '/login' })
    } else {
        next()
    }
});


/* eslint-disable no-new */
new Vue({
    //el: '#app',
    //template: '<App/>',
    router,
    store,
    axios,
    // components: { App },
    render: h => h(App)
}).$mount('#app');

// 全局过滤器
for (var filter in filters) {
    if (filters.hasOwnProperty(filter)) {
        Vue.filter(filter, filters[filter])
    }
}


