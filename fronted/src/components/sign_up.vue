<template>
  <div class="frame">

      <div id="box2">

        <h2 style="color: #fffefe">Register</h2>
        <el-form :rules="rules" :model="form" label-position="left"   label-width="140px" target="sendcode" >
          <el-form-item label="Username" prop="name">
            <el-input v-model="form.name"  placeholder="请输入用户名"  />
          </el-form-item>
          <el-form-item label="password" prop="password">
              <el-input  v-model="form.password" type="password" placeholder="请输入密码"/>
          </el-form-item>
          <el-form-item label="password again"  prop="againpassword" >
            <el-input v-model="form.again_password" type="password" placeholder="再次输入密码"/>
          </el-form-item>
          <el-form-item label="email" prop="email">
            <el-input v-model="form.email" placeholder="请输入邮箱" />
          </el-form-item>
          <el-form-item>
            <el-input  v-model="incode" placeholder="请输入验证码" style="width: 400px ;height: 40px"  >
              <!-- eslint-disable-next-line-->
              <template #append style="box-shadow: none;background: none">
                <el-button @click="sendCode" :disabled="issend || countdown>0"  style="background-color: transparent;width: 150px;margin-top : 0;text-align: center;padding-top:0" >
                  <span style="color:#7FFF00">{{ issend ? '点击发送验证码' : countdown > 0 ? `${countdown}秒后重试` : '发送验证码' }}
                  </span>
                </el-button>
              </template>
            </el-input>

          </el-form-item>
          <button @click="TryRegister" >注册</button><br>
        </el-form>
        <br>
        <p id="switch">拥有账号？<a href="#/login_in" style="color:deeppink" >登录</a> </p>
        <iframe width=0 height=0 frameborder=0 name="sendcode"></iframe>
      </div>

  </div>
</template>

<script setup>
import {reactive, ref} from "vue";
import  checksame from "../assets/js/signup.js";
// eslint-disable-next-line no-unused-vars
import axios from "axios";
// import api from "../api/request.js"
const form = reactive({
  name : '',
  password:'',
  again_password:'',
  email:'',
  code:'',

    }
)
const rules = ref({
  name:[
    { required: true, trigger: 'blur' ,pattern: /[a-zA-Z0-9]{4,20}/, message: '名称只能由英文字母数字及下划线组成,且长度为4-20个字符'}
  ],
  password:[
    {required: true, trigger: 'blur' ,pattern: /[a-zA-Z0-9]{6,30}/,message:"密码只能由英文字母和数字下划线组成，长度在6-30之间",}
  ],
  againpassword:[
    {required: true, trigger: 'blur' ,validator: checksame,}
  ],
  email:[
    {required: true, trigger: 'blur',pattern: /[a-zA-Z0-9][\w\\.-]*[a-zA-Z0-9]@[a-zA-Z0-9][\w\\.-]*[a-zA-Z0-9]\.[a-zA-Z][a-zA-Z\\.]*[a-zA-Z]/, message:"邮箱格式不正确"},
  ]
})
const issend =  ref(false);
const countdown = ref(0);
const incode = ref('');
const reg = /[a-zA-Z0-9][\w\\.-]*[a-zA-Z0-9]@[a-zA-Z0-9][\w\\.-]*[a-zA-Z0-9]\.[a-zA-Z][a-zA-Z\\.]*[a-zA-Z]/

// eslint-disable-next-line no-unused-vars
function startCountdown() {
  countdown.value = 60;
  const timer = setInterval(() => {
    countdown.value--;
    if (countdown.value <= 0) {
      clearInterval(timer);
    }
  }, 1000);
}
function sendCode() {
  if (countdown.value > 0 || issend.value|| !reg.test(form.email) ) {
    return;
  }
  issend.value = true;

  console.log(form.email)
  axios.post("/apis/api/register/sendcode",{
    txt: 'test',
    email:form.email,
  }, {
    headers: {
      'Content-Type': 'application/json'
    }
  }).then(response=>{
    console.log(response)
    issend.value = false;
  }).catch(error=>{
    console.log(error)
  })
  startCountdown();
}
function TryRegister(){
  axios.post("/apis/api/register",{
    name:form.name,
    password:form.password,
    email:form.email,
    againpassword:form.again_password,
    code:incode.value,
  },{
    headers:{
      'Content-Type':'application/json'
    }
  }).then((response)=>{
    console.log(response)
  }).catch((error)=>{
    console.log(error)
  })
}
</script>

<style scoped lang="less">
@import "../assets/style/login_register.less";
#box2 {
  width: 500px;
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