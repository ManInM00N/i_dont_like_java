import {createRouter, createWebHashHistory} from "vue-router";
import Home_Main from "../components/Home_Main.vue"
import login_in  from "../components/login_in"
import MyselfMsg from "../components/MyselfMsg"
const router = createRouter({
    history: createWebHashHistory(),
    routes:[
        {
            path: '/',
            component: Home_Main,
            meta :{
                keepAlive :false,
                refresh: false
            },

        },
        {
            path: '/login_in',
            component: login_in,
            meta :{
                keepAlive :false,
                refresh: false
            },

        },
        {
            path: '/MyselfMsg',
            component: MyselfMsg,
            meta:{
                keepAlive :false,
                refresh: false
            }
        }



    ]



})
export default router