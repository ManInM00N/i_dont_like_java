<template>
  <div class = "frame">
    <form action="login" method="post">
      <div id ="box">
        <h2 style="color: #fffefe">Login</h2>
        <br>
        <!--                                <br>-->
        <div class ="input_box">
          <el-input type="text" name="username" id="username" v-model="account" placeholder="请输入用户名"/>
        </div>
        <br>
        <div class ="input_box">
          <el-input type="password" name="password" id="password" v-model="password" placeholder="请输入密码"/>
        </div>
        <br>
        <button>登录</button><br>
        <p id="switch">没有账号？<a href="#/sign_up" style="color:deeppink" >注册</a> </p>
      </div>
    </form>
  </div>
</template>


<script setup>
import {ref} from "vue";

const account = ref('')
const password = ref('')
// function ifnull(txt){
//   if(txt.length==0){
//     return true;
//   }
//   var $reg=/\s+/;
//   return $reg.test(txt);
// };
// import sign_up from "sign-up.vue";
</script>

<style scoped lang="less">
@import "../assets/style/login_register.less";
</style>