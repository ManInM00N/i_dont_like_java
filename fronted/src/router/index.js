import {createRouter, createWebHashHistory} from "vue-router";
import Home_Main from "../components/Home_Main.vue"
import login_in  from "../components/login_in"
import MyselfMsg from "../components/MyselfMsg"
import HomeTown from "../components/HomeTown.vue";
import CommentBoard from "@/components/CommentBoard.vue";
import sign_up from "@/components/sign_up.vue";
import UserMsg from "@/components/UserMsg.vue";
import ReRefresh from "@/components/ReRefresh.vue";
const router = createRouter({
    history: createWebHashHistory(),
    routes:[
        {
            path: '/',
            component: Home_Main,
            meta :{
                keepAlive :false,
                refresh: true
            },

        },
        {
            path: '/login_in',
            component: login_in,
            meta :{
                keepAlive :false,
                refresh: true
            },

        },
        {
            path: '/MyselfMsg',
            component: MyselfMsg,
            meta:{
                keepAlive :false,
                refresh: true
            }
        },
        {
            path: '/HomeTown',
            component: HomeTown,
            meta:{
                keepAlive :false,
                refresh: true,
            },
            props: (route) => ({ query: route.query.search })
        },
        {
            path: '/CommentBoard',
            component: CommentBoard,
            meta:{
                keepAlive :false,
                refresh: true,
            }
        },{
            path: '/sign_up',
            component: sign_up,
            meta:{
                keepAlive :false,
                refresh: true,
            }
        },{
            path:'/user/:id' ,
            component : UserMsg,
            meta:{
                keepAlive :true,
                refresh: true,
            }
        },{
            path: '/refresh',
            component:ReRefresh,
            meta: {
                title: '用于同路由刷新',
                refresh: true,
            }
        }





    ]



})
export default router