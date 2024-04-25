<template>
  <div class="frame">

    <form  method="post" onsubmit="return check()" >
      <div id="box">

        <h2 style="color: #fffefe">Register</h2>
        <el-form :rules="rules" :model="form">
          <el-form-item label="User name" prop="name">
            <el-input v-model="form.name"  placeholder="请输入用户名"  />
          </el-form-item>
          <el-form-item label="">

          </el-form-item>
        </el-form>
        <div class="clear"></div>
        <div class = "input_box">
          <input type="password" name="passwordfirst" id="passwordfirst" placeholder="请输入想设定的密码" ><span id="passwordfir"></span>
        </div>
        <div class="clear"></div>

        <div class = "input_box">
          <input type="password" name="passwordsecond" id="passwordsecond" placeholder="再次输入密码"><span id="passwordsec"></span>
        </div>
        <div class="clear"></div>


        <div class = "input_box">
          <input type="text" name="email" id="email" placeholder="请输入邮箱地址"><span id="email_msg"></span>
        </div>
        <div class="clear"></div>
        <button  >注册</button><br>
        <br>
        <p id="switch">拥有账号？<a href="#/login_in" style="color:deeppink" >登录</a> </p>
      </div>

    </form>

  </div>
</template>

<script setup>
import {reactive, ref} from "vue";

const form = reactive({
  name : ''



    }
)
const rules = ref({
  name:[
    { required: true, trigger: 'blur' },
    { min: 4, max: 20, message: '长度在 4 到 20 个字符', trigger: 'blur' },
    {pattern: '/^[a-zA-Z0-9]+$/', message: '名称只能由英文字母数字及下划线组成'}
  ]
})
</script>

<style scoped lang="less">
@import "../assets/style/login_register.less";

</style>