import axios from 'axios'
import router from "./router"
import config from '../config';

var instance = axios.create({
    baseURL: config.env.API_ROOT,
    withCredentials: true,
});

// axios.defaults.baseURL= "http://127.0.0.1:8888";
// axios.defaults.withCredentials = true;


// http request 拦截器
instance.interceptors.request.use(
    config => {
            if (sessionStorage.getItem('user') != "") {
              config.headers.Authorization = sessionStorage.getItem('user');
            }
        return config;
    },
    err => {
        return Promise.reject(err);
});

// http response 拦截器
instance.interceptors.response.use(function(response){
  return response;
},function(error){
    //对返回的错误进行一些处理
    if (error.response) {
        if (error.response.status == 401) {
            sessionStorage.removeItem('user');
            router.push({ path: '/login' })
        }
    }
    return Promise.reject(error.response);
});

export default instance;
