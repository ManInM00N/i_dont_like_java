/* eslint-disable */

<template>
  <el-container
      class="middle_problem"
      style="align:center"
  >
    <section id="tip">
      <div  class="information" >
        <h1 align="center" style="font-size:26px">个人简介</h1>
        <el-row>
          <el-col :span="6"></el-col>
          <el-col :span="6">
            <div align="center">
              <el-image :src="form.url" :fit="fill" style="height: 100px ;width:100px" />
              <el-upload
                action="#"
                ref="up"
                :limit="1"
                :before-upload="beforeUpload"
                :on-exceed="handleExceed"
                :auto-upload="false"
                :http-request="upload"
              >
                <template #trigger>
                  <el-button type="primary">上传头像</el-button>
                </template>
              </el-upload>
              <el-button @click="submitUpload">
                 修改头像
              </el-button>
            </div>
          </el-col>
          <el-col :span="6"/>
          <el-col :span="6"/>
        </el-row>
        <el-row>
          <el-col :span="6"/>
          <el-col :span="4"><div>姓名</div> </el-col>
          <el-col :span="8"><div >{{form.name}}</div> </el-col>
          <el-col :span="6"/>
        </el-row>
        <el-row>
          <el-col :span="6"/>
          <el-col :span="4"><div >求学经历：</div> </el-col>
          <el-col :span="8"><el-input v-model="form.xueli" ></el-input> </el-col>
          <el-col :span="6"/>
        </el-row>
        <el-row>
          <el-col :span="6"/>
          <el-col :span="4"><div >获奖经历：</div> </el-col>
          <el-col :span="8"><el-input el-input v-model="form.awards"> </el-input> </el-col>
          <el-col :span="6"/>
        </el-row>
        <el-row>
          <el-col :span="6"/>
          <el-col :span="4"><div >参与社团：</div> </el-col>
          <el-col :span="8"><el-input v-model="form.group"></el-input> </el-col>
          <el-col :span="6"/>
        </el-row>
        <el-row>
          <el-col :span="6"/>
          <el-col :span="4"><div >兴趣爱好：</div> </el-col>
          <el-col :span="8"><el-input v-model="form.interest" ></el-input> </el-col>
          <el-col :span="6"/>
        </el-row>
        <el-row>
          <el-col :span="6"/>
          <el-col :span="4"><div >座右铭：</div> </el-col>
          <el-col :span="8"><el-input v-model="form.motto"></el-input> </el-col>
          <el-col :span="6"/>
        </el-row >
        <el-row>
          <el-col :span="8"/>
          <el-col :span="3"/>
          <el-col :span="2"><el-button @click="update" color="#626aef" :dark="true" plain>修改信息</el-button> </el-col>
          <el-col :span="3"/>
          <el-col :span="8"/>
        </el-row>
      </div>
    </section>
  </el-container>
</template>
<script lang="js" setup>
import {ref} from "vue";
import axios from "axios";
import {useRoute} from "vue-router";
import {ElMessage} from "element-plus";
import router from "@/router";
const route = useRoute()
const up =  ref()
const form = ref({
  name:'',
  xueli:'',
  awards:'',
  interest:'',
  motto:'',
  group:'',
  url:'',
})
function init(){
  form.value.name = route.params.id
  axios.get("/apis/user/"+form.value.name).then(response=>{

    console.log(response)
    form.value.name = response.data.name
    form.value.awards = response.data.awards
    form.value.group = response.data.groups
    form.value.motto = response.data.motto
    form.value.xueli = response.data.xueli
    form.value.interest= response.data.interest
    form.value.url = require("../assets/images/" + response.data.url)
  }).catch(error=>{
    console.log(error)

    if (error.response.status===404){
      router.replace("/")
    }
  })
}
// function changeHeader(param){
// }
function update(){
  axios.post("/apis/api/update",{
    name:form.value.name,
    xueli:form.value.xueli,
    interest:form.value.interest,
    motto:form.value.motto,
    group:form.value.group,
    awards:form.value.awards,
  },{
    headers:{
      'Content-Type':'application/json'
    }
  }).then((response)=>{
    console.log(response)
    if (response.status===200){
      console.log("goto",form.value.name)
    }
  }).catch((error)=>{
    console.log(error)
  })
}
function handleExceed(){
  ElMessage.warning("超出最大上传文件数量的限制！")
  return false;

}
function beforeUpload(file) {
  const isJPG = file.type === "image/jpeg" || file.type == "image/png";
  const isLt2M = file.size / 1024 / 1024 < 4;
  if (!isJPG) {
    ElMessage.error("上传头像图片只能是 JPG 或 PNG 格式!")
    return false;
  }
  if (!isLt2M) {
    ElMessage.error("上传头像图片大小不能超过 2MB!")
    return false;
  }

  return true;
}
const submitUpload = () => {
  up.value.submit()
}
function upload(file){
  // const fileData = new FormData();
  // fileData.append("avatar", file);
  const formData = new FormData();
  formData.append("file", file.file);
  formData.append("name",form.value.name)
  axios.post("/apis/api/headerchange",formData,{
    headers:{
      'Content-Type':'multipart/form-data'
    }
  }).then(response=>{

    form.value.url = require("../assets/images/"+ response.data.url)
  }).catch(error=>{
    console.log(error)
  })
}
init()

</script>

<style scoped lang="less">
.information {

  width: 650px;
  height: 700px;
  background-color:rgba(255, 255, 255, 0.25);
  margin-top: 5%;
  border-radius: 20px;
  padding: 25px 25px;
  padding-bottom: 0px;
  display: block;
  margin-left: calc(50% - 325px);
  margin-right: calc(50% - 325px);
}
.middle_problem{
  display:flex;
  justify-content: center;
  align-items:center;
  text-align: left;
  width: 100%;
  height: 100%;
  line-height:25px;
  margin-bottom: 10%;
  & #tip{
    align: center;
    display: inline-block;
    border: none;
    font-size: 14px;
    .el-row{
      margin-bottom:10px;
    }
  }
}
</style>