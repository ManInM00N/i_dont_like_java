<template>
  <el-section
      class="inline "
    style="min-height : 100vh;display: flex;
  flex-direction: column;"
  >

  <el-header
    style="background: transparent;text-align:center"
    class=""
  >

    <el-menu
        :default-active="$route.path"
        class="vertical-menu"
        mode="vertical"
        style="background-color: rgba(255,255,255,0.2);"
        :router="true"
    >
    <el-container class="top-items" >
      <el-menu-item
          v-for="(item,idx) in items"
          :key="item.key"
          :id="item.id"
          :index="idx+''"
          :route="item.index"
          :title="item.key"
          class="menu_item"
          @select="handleMenuSelect"
      >
        <el-container
            class="item_body"
        >
          <el-icon size="30" class="item_icon">
            <component
                :is="item.iconmsg"
                style="color: #32CD99"
            />
          </el-icon>
        </el-container>
      </el-menu-item>
    </el-container>

    </el-menu>

  </el-header>
  <section class="forall">
    <section style="height:90%">

      <router-view v-slot="{ Component,route }">

        <keep-alive v-if="route.meta.keepAlive ">
          <component
              :is="Component"
              class="forall"
          />
        </keep-alive>
        <component v-else :is="Component" class="forall" />
      </router-view>
    </section>
    <el-divider  style="margin-top: 5px;margin-bottom: 5px"/>
    <el-footer  class="footer" align="center"  >
      <p >个人学号：2200303310 姓名：李瑟钰 班级：22计算机3</p>
      <p >联系方式： QQ:571404393</p>
      <p>{{dateFormat(date)}}  你是今天第{{num}}个访客</p>
    </el-footer>
  </section>
  </el-section>

</template>

<script setup>
import {onMounted, ref} from "vue";
import {dateFormat} from "../assets/js/signup.js"
import axios from "axios"
const num = ref(0)
const date = ref()
const timer = ref()
const items = ref([
  {id : 1, iconmsg : "HomeFilled",key:"Home",index:"/"},
  {id : 2, iconmsg : "ChatDotRound",key:"CommentBoard",index:"/CommentBoard"},
  {id : 3, iconmsg:"Avatar",key:"Myself",index:"/MyselfMsg"},
  {id:4,iconmsg: "Promotion",key:"HomeTown",index:"/HomeTown"},
  {id:5,iconmsg:"ChatDotSquare",key:"chat",index:"/chat"},
  {id:6,iconmsg: "User",key:"Sign in",index:"/login_in"},

])


timer.value = setInterval(()=>{
  date.value = new Date()
})
onMounted(()=>{
axios.get("/apis/").then(response=>{
console.log(response)
num.value=response.data.num
}).catch(error=>{
console.log(error)
})

});
</script>

<style lang="less" scoped>

@import "src/assets/style/Color.less";
@import "src/assets/style/menu.less";
.forall{
  height:100%;
  margin-bottom: 0;
}
/deep/
.footer{
  align-items: center;float:left;width:100%;
  position: relative;
  &p{
    margin: 0 0 0 0;
  }
}

</style>