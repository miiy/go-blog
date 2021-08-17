// https://github.com/axios/axios
import axios from "axios";

const instance = axios.create({
    // baseURL: 'http://localhost:8051'
})

// const token = store.getters.token
// if (token.access_token) {
// instance.defaults.headers.common['Authorization'] = AUTH_TOKEN;
//     config.headers.common['Authorization'] = token.token_type + ' ' + token.access_token
// }
// const user = store.getters.user
// if (token.access_token && !user.name) {
//   console.log('get user')
//   store.dispatch('getUser').then(resp => {
//     console.log('set user' + resp.data.name)
//   }).catch(err => {
//     console.log(err)
//     router.push({ name: 'Login' })
//   })
// }



// Add a request interceptor
instance.interceptors.request.use(function (config) {
    // Do something before request is sent
    return config;
}, function (error) {
    // Do something with request error
    return Promise.reject(error);
});

// Add a response interceptor
instance.interceptors.response.use(function (response) {
    // Any status code that lie within the range of 2xx cause this function to trigger
    // Do something with response data
    return response;
}, function (error) {
    // Any status codes that falls outside the range of 2xx cause this function to trigger
    // Do something with response error
    return Promise.reject(error);
});

export default instance