/* eslint-disable */

<template>
  <el-container
      class="middle_problem"
      style="align:center"
  >
    <section id="tip">
      <div  class="information" >
        <h1 align="center" style="font-size:26px">个人简介</h1>
        <!--      <el-image :src='require("@../assets/images/tx.jpg")' alt="" align="center" />-->
        <el-row>
          <el-col :span="6"></el-col>
          <el-col :span="6">
            <div align="center">
              <img src="../assets/images/tx.jpg" alt="" />
            </div>

          </el-col>
          <el-col :span="6"></el-col>
          <el-col :span="6"></el-col>
        </el-row>
        <el-row>
          <el-col :span="6"><div ></div></el-col>
          <el-col :span="6"><div>姓名</div> </el-col>
          <el-col :span="6"><div >{{form.name}}</div> </el-col>
          <el-col :span="6"><div ></div> </el-col>
        </el-row>
        <el-row>
          <el-col :span="6"><div></div></el-col>
          <el-col :span="6"><div >求学经历：</div> </el-col>
          <el-col :span="6"><el-input v-model="form.xueli" ></el-input> </el-col>
          <el-col :span="6"><div ></div> </el-col>
        </el-row>
        <el-row>
          <el-col :span="6"><div ></div></el-col>
          <el-col :span="6"><div >获奖经历：</div> </el-col>
          <el-col :span="6"><el-input el-input v-model="form.awards"> </el-input> </el-col>
          <el-col :span="6"><div ></div> </el-col>
        </el-row>
        <el-row>
          <el-col :span="6"><div ></div></el-col>
          <el-col :span="6"><div >参与社团：</div> </el-col>
          <el-col :span="6"><el-input v-model="form.group"></el-input> </el-col>
          <el-col :span="6"><div ></div> </el-col>
        </el-row>
        <el-row>
          <el-col :span="6"><div ></div></el-col>
          <el-col :span="6"><div >兴趣爱好：</div> </el-col>
          <el-col :span="6"><el-input v-model="form.interest" ></el-input> </el-col>
          <el-col :span="6"><div ></div> </el-col>
        </el-row>
        <el-row>
          <el-col :span="6"><div ></div></el-col>
          <el-col :span="6"><div >座右铭：</div> </el-col>
          <el-col :span="6"><el-input v-model="form.motto"></el-input> </el-col>
          <el-col :span="6"><div ></div> </el-col>
        </el-row >
        <el-row>
          <el-col :span="8"><div ></div></el-col>
          <el-col :span="3"> </el-col>
          <el-col :span="2"><el-button @click="update" color="#626aef" :dark="true" plain>修改信息</el-button> </el-col>

          <el-col :span="3"> </el-col>

          <el-col :span="8"> </el-col>
        </el-row>
      </div>
    </section>
  </el-container>
</template>
<script setup>
import {reactive} from "vue";
import axios from "axios";
// eslint-disable-next-line no-unused-vars
import {useRoute} from "vue-router";
const route = useRoute()
const form = reactive({
  name:'',
  xueli:'',
  awards:'',
  interest:'',
  motto:'',
  group:'',

})
function init(){
  form.name = route.params.id
  axios.get("/apis/user/"+form.name).then(response=>{
    console.log(response)
    form.name = response.data.name
    form.awards = response.data.awards
    form.group = response.data.groups
    form.motto = response.data.motto
    form.xueli = response.data.xueli
    form.interest= response.data.interest
  }).catch(error=>{
    console.log(error)
  })
}
function update(){
  axios.post("/apis/api/update",{
    name:form.name,
    xueli:form.xueli,
    interest:form.interest,
    motto:form.motto,
    group:form.group,
    awards:form.awards,
  },{
    headers:{
      'Content-Type':'application/json'
    }
  }).then((response)=>{
    console.log(response)
    if (response.status===200){
      console.log("goto",form.name)
    }
  }).catch((error)=>{
    console.log(error)
  })
}

init()

</script>

<style scoped lang="less">
.information {

  width: 650px;
  height: 600px;
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
  }
}
</style>