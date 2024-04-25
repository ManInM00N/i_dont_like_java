<template >

  <el-main class="Main" style="margin: 0px 0px 0px 0px" v-loading.fullscreen.lock="load">
    <el-header>
<!--      <div class="School">-->
        <p align="center"><img src="./assets/images/xh.png" alt="">中国计量大学-Web前端设计与开发</p>
<!--      </div>-->
    </el-header>
      <Menu_List/>

  </el-main>

</template>

<script setup>
import Menu_List from "./components/Menu_List.vue"
import {ref} from "vue";
// import axios from 'axios';
const load = ref(true);
// const num = ref(0);
function  done(){
  load.value=false;
}

window.addEventListener('load',done)
// axios.get("http://127.0.0.1:7234/")

</script>

<style lang='less' scoped>
@import "assets/style/variable.less";
@import "assets/style/load.less";
</style>
