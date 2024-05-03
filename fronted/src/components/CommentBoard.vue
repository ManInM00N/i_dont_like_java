<template>
  <section class="middle_title">
    <div class="information">
      <h1 class="Title">留言板</h1>
      <el-divider border-style="dotted"/>
      <div id="messageSub" >
        <div v-for="(item) in q" :key="item.id">
          <el-card class="box-card">
            <template #header>{{item.name}} :</template>
              {{ item.inner}} <br>
          </el-card>
        </div>
      </div>
      <el-divider></el-divider>
      <el-card>
        <template #header>
          <el-input v-model="name" >
            <template #prepend>
              留言人：
            </template>
            <template #append>
              <el-button @click="comment" tag="发送留言">
                发送留言
              </el-button>
            </template>
          </el-input>

        </template>
        <el-input type="textarea" v-model="inner"  placeholder="输入评论..." >
        </el-input>

      </el-card>
    </div>
  </section>

</template>
<script setup>
import {h, onMounted, ref} from "vue";
import axios from "axios";
import {ElMessage} from "element-plus";
// eslint-disable-next-line no-unused-vars
const name = ref('')
const inner = ref('')
let q = ref([
])
onMounted(()=>{
  axios.get("/apis/api/comment").then((response)=>{
    q.value=response.data.res
    console.log(q.value)
  }).catch((error)=>{
    console.log(error)
  })
})
function comment(){
  if (inner.value.length==0){
    ElMessage.error('评论内容为空哦')
    return
  }
  axios.post("/apis/api/comment",{
    name:name.value,
    inner:inner.value ,
  },{
    headers:{
      'Content-Type':'application/json'
    }
  }).then((response)=>{
    ElMessage({
      message: h('p', { style: 'line-height: 1; font-size: 14px' }, [
        h('i', { style: 'color: teal' }, '评论发送成功'),
      ]),
    })
    q.value.push(response.data.res)
  }).catch((error)=>{
    console.log(error)
  })
}
</script>

<style scoped lang="less">
.Title{
  margin: auto;
  text-align: center;
  margin-top: 10px;
  font-size: 30px;
  color: @light-pink;
  text-shadow:2px 2px #777;
}
.information {

  width: 650px;
  height: 100%;
  min-height:650px;
  background-color:rgba(255, 255, 255, 0.25);
  margin-top: 5%;
  border-radius: 20px;
  padding: 25px 25px;
  padding-bottom: 0px;
  display: flex;
  flex-direction: column;
  margin-left: calc(50% - 325px);
  margin-right: calc(50% - 325px);
}
.box-card {
  margin-bottom: 10px;
}
@import "../assets/style/Color.less";
@import "../assets/style/variable.less";
</style>