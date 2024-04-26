<template>
  <div class="frame">

    <form  method="post" onsubmit="return check()" >
      <div id="box2">

        <h2 style="color: #fffefe">Register</h2>
        <el-form :rules="rules" :model="form">
          <el-form-item label="Username" prop="name">
            <el-input v-model="form.name"  placeholder="请输入用户名"  />
          </el-form-item>
          <el-form-item label="password" prop="password">
              <el-input minlength="20" v-model="form.password" type="password" placeholder="请输入密码"/>
          </el-form-item>
          <el-form-item label="password again"  prop="againpassword" >
            <el-input v-model="form.again_password" type="password" placeholder="再次输入密码"/>
          </el-form-item>
          <el-form-item label="email" prop="email">
            <el-input v-model="form.email" placeholder="请输入邮箱" />
          </el-form-item>
        </el-form>
        <button  >注册</button><br>
        <br>
        <p id="switch">拥有账号？<a href="#/login_in" style="color:deeppink" >登录</a> </p>
      </div>

    </form>

  </div>
</template>

<script setup>
import {reactive, ref} from "vue";
import  checksame from "../assets/js/signup.js";
const form = reactive({
  name : '',
  password:'',
  again_password:'',
  email:'',
  interest :'',


    }
)
const rules = ref({
  name:[
    { required: true, trigger: 'blur' ,pattern: /[a-zA-Z0-9]{4,20}/, message: '名称只能由英文字母数字及下划线组成,且长度为4-20个字符'}
  ],
  password:[
    {required: true, trigger: 'blur' ,pattern: /[a-zA-Z]{6,30}/,message:"密码只能由英文字母和数字下划线组成，长度在6-30之间",}
  ],
  againpassword:[
    {required: true, trigger: 'blur' ,validator: checksame,}
  ],
  email:[
    {required: true, trigger: 'blur',pattern: /[a-zA-Z0-9][\w\\.-]*[a-zA-Z0-9]@[a-zA-Z0-9][\w\\.-]*[a-zA-Z0-9]\.[a-zA-Z][a-zA-Z\\.]*[a-zA-Z]/, message:"邮箱格式不正确"},
  ]
})

</script>

<style scoped lang="less">
@import "../assets/style/login_register.less";

</style>