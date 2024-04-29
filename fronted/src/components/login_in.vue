<template>
  <div class = "frame">
      <div id ="box">
        <h2 style="color: #fffefe">Login</h2>
        <el-form v-model="form" @onsubmit="subform" label-position="left"   label-width="120px" target="sendcode">
        <el-form-item label="Username" class ="input_box">
          <el-input type="text"  v-model="form.account" placeholder="请输入用户名"/>
        </el-form-item>
        <el-form-item label="password" class ="input_box">
          <el-input type="password" name="password" id="password" v-model="password" placeholder="请输入密码"/>
        </el-form-item>
        <button>登录</button><br>
        <p id="switch">没有账号？<a href="#/sign_up" style="color:deeppink" >注册</a> </p>
        </el-form>
        <iframe width=0 height=0 frameborder=0 name="sendcode"></iframe>

      </div>
  </div>
</template>


<script setup>
// eslint-disable-next-line no-unused-vars
import {reactive, ref} from "vue";
import router from "@/router";
import axios from "axios";
const form = reactive({
    account: "",
    password:"",
})
function subform(){
    axios.get("/apis/login",{
      name:form.account,
      password:form.password,
    },{
      headers:{
        'Content-Type':'application/json'
      }
    }).then((response)=>{
      console.log(response)
      if (response.status===200){
        localStorage.clear()
        localStorage.setItem("")
        router.push("/user/"+form.name)
      }
    }).catch((error)=>{
      console.log(error)
    })
}



</script>

<style scoped lang="less">
@import "../assets/style/login_register.less";

#box {
  width: 420px;
  height: 600px;
  background-color: #00000060;
  margin: auto;
  margin-top: 5%;
  text-align: center;
  border-radius: 10px;
  padding: 50px 50px;
  padding-bottom: 0px;
}

</style>