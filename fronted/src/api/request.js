import axios from "axios";
//创建一个axios对象
const instance = axios.create({
    withCredentials: true, // 是否允许带cookie这些
    baseURL: "http://localhost:7234", //这里配置的是后端服务提供的接口
    timeout: 2000
})

//全局请求拦截
instance.interceptors.request.use(
    function (config) {
        console.group('全局请求拦截')
        console.log(config)
        console.groupEnd()
        config.headers.token = '12345'//例如可以在请求头里面设置token值，token变量名前后端约定
        return config;
    },
    function (err) {
        return Promise.reject(err);
    },

)

instance.interceptors.response.use(function (response) {
    console.log('全局响应拦截');
    console.log(response)
    console.groupEnd()

    return response
}),function (err) {
    return Promise.reject(err)

}


export function get (url,params) {
    return instance.get(url,{
        params
    });
}

export function post (url,data) {
    return instance.post(url,data);
}

export function del (url) {
    return instance.delete(url);
}

export function put(url,data) {
    return instance.put(url,data);
}