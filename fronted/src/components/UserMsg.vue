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
          <el-col :span="6"><div >{{form.xueli}}</div> </el-col>
          <el-col :span="6"><div ></div> </el-col>
        </el-row>
        <el-row>
          <el-col :span="6"><div ></div></el-col>
          <el-col :span="6"><div >获奖经历：</div> </el-col>
          <el-col :span="6"><div >{{form.awards}}</div> </el-col>
          <el-col :span="6"><div ></div> </el-col>
        </el-row>
        <el-row>
          <el-col :span="6"><div ></div></el-col>
          <el-col :span="6"><div >参与社团：</div> </el-col>
          <el-col :span="6"><div >{{form.group}}</div> </el-col>
          <el-col :span="6"><div ></div> </el-col>
        </el-row>
        <el-row>
          <el-col :span="6"><div ></div></el-col>
          <el-col :span="6"><div >兴趣爱好：</div> </el-col>
          <el-col :span="6"><div >{{form.interest}}</div> </el-col>
          <el-col :span="6"><div ></div> </el-col>
        </el-row>
        <el-row>
          <el-col :span="6"><div ></div></el-col>
          <el-col :span="6"><div >座右铭：</div> </el-col>
          <el-col :span="6"><div >{{form.motto}}</div> </el-col>
          <el-col :span="6"><div ></div> </el-col>
        </el-row >
      </div>
    </section>
  </el-container>
</template>
<script setup>
import {reactive} from "vue";
import axios from "axios";
const form = reactive({
  name:'',
  xueli:'',
  awards:'',
  interest:'',
  motto:'',
  group:'',

})
function init(){
  form.name = this.$route.params.id
  axios.get("/apis/").then(response=>{
    console.log(response)
    form.name = response.data.name
    form.awards = response.data.awards
    form.group = response.data.awards
    form.motto = response.data.motto
    form.xueli = response.data.xueli
    form.interest= response.data.interest
  }).catch(error=>{
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