<template>
  <div class="frame">

      <div id="box2">

        <h2 style="color: #fffefe">Register</h2>
        <el-form :rules="rules" :model="form" label-position="left"   label-width="140px" target="sendcode" >
          <el-form-item label="Username" prop="name">
            <el-input v-model="form.name"  placeholder="请输入用户名"  />
          </el-form-item>
          <el-form-item label="password" prop="password">
              <el-input  v-model="form.password" type="password" placeholder="请输入密码" show-password="true"/>
          </el-form-item>
          <el-form-item label="password again"  prop="againpassword" >
            <el-input v-model="form.again_password" type="password" placeholder="再次输入密码" show-password="true"/>
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
        <p id="switch">拥有账号？<a href="/login_in" style="color:deeppink" >登录</a> </p>
        <iframe width=0 height=0 frameborder=0 name="sendcode"></iframe>
      </div>

  </div>
</template>

<script setup>
import {reactive, ref} from "vue";
import axios from "axios";
import router from "@/router";
import {ElMessage} from "element-plus";
const form = reactive({
  name : '',
  password:'',
  again_password:'',
  email:'',
  code:'',
    }
)
const checksame = (rule, value, callback) => {
  if (form.again_password === '') {
    callback(new Error('请再次输入密码'))
  } else if (form.again_password !== form.password) {
    callback(new Error('两次输入密码不一致!'))
  } else {
    callback()
  }
};
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
    ElMessage.warning("邮箱可能不合适")
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
  }).catch(error=>{
    console.log(error)
  }).finally(()=>{
    issend.value = false;

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
  }).then(response=>{
    console.log(response)

    if (response.status===200){
      console.log("goto",form.name)
      router.push({path:"/user/"+form.name,params:{
          id:form.name
        }})
    }

  }).catch((error)=>{
    if (error.response.data.code!==200){
      ElMessage.error(error.response.data.message)
    }
    console.log(error)
  })
}
</script>

<style scoped lang="less">
@import "../assets/style/login_register.less";
@import "../assets/style/variable.less";

</style>