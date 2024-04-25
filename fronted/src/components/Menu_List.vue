<template>
  <el-header
    style="background: transparent"
  >

    <el-menu
        :default-active="$route.path"
        class="vertical-menu"
        mode="vertical"
        style="background-color: rgba(255,255,255,0.2)"
        :router="true"
    >
    <el-container class="top-items" >

      <el-menu-item
          v-for="(item,idx) in items"
          :key="item.key"
          :id="item.id"
          :index="idx+''"
          :route="item.index"
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
          />
        </keep-alive>
        <component v-else :is="Component" class="forall"/>
      </router-view>
    </section>
    <el-divider style="margin-top: 5px;margin-bottom: 5px"/>
<!--    <el-backtop :right="100" :bottom="100" :visibility-height="200" />-->
    <el-container class="footer" align="center"  >
      <p >个人学号：2200303310 姓名：陈忠鹏 班级：22计算机3      <br>
      </p>
      <p >联系方式：15757285081 QQ:571404393</p>
      <p>你是今天第{{num}}个访客</p>
    </el-container>
  </section>
</template>

<script setup>
// eslint-disable-next-line no-unused-vars
import {defineComponent, onMounted, ref} from "vue";
import axios from "axios"
const num = ref(0)
const items = ref([
  {id : 1, iconmsg : "HomeFilled",key:"Search",index:"/"},
  {id : 2, iconmsg : "UserFilled",key:"Login_in",index:"/login_in"},
  {id : 3, iconmsg:"Avatar",key:"MyselfMsg",index:"/MyselfMsg"},
  {id:4,iconmsg: "User",key:"User",index:"/User"},

])
onMounted(()=>{
    axios.get("http://192.168.216.244:7234/").then(response=>(
        num.value=response.data.num
    ))

});

// eslint-disable-next-line no-unused-vars
const activeIndex = ref('/')




// eslint-disable-next-line no-unused-vars
function  handleMenuSelect(index) {
  console.log("ee",this.activeIndex)
}
</script>

<style lang="less" scoped>

@import "src/assets/style/Color.less";
/*@import "src/assets/style/menu.less";*/
.forall{
  height:100%;
}
/deep/
.footer{
  align-items: center;float:left;width:100%;height:10%;
  &p{
    margin: 0 0 0 0;
  }
}

</style>