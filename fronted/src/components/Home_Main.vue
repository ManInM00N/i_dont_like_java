<template>
  <el-main
    class="middle_problem"
  >
    <section id="tip">
    <div  class="information" >
      <h1 align="center">个人简介</h1>
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
        <el-col :span="6"><div >陈忠鹏</div> </el-col>
        <el-col :span="6"><div ></div> </el-col>
      </el-row>
      <el-row>
        <el-col :span="6"><div></div></el-col>
        <el-col :span="6"><div >求学经历：</div> </el-col>
        <el-col :span="6"><div >	中国计量大学本科</div> </el-col>
        <el-col :span="6"><div ></div> </el-col>
      </el-row>
      <el-row>
        <el-col :span="6"><div ></div></el-col>
        <el-col :span="6"><div >获奖经历：</div> </el-col>
        <el-col :span="6"><div >蓝桥杯省二、天梯赛国二+国一、天梯赛个人银、浙江省省赛银</div> </el-col>
        <el-col :span="6"><div ></div> </el-col>
      </el-row>
      <el-row>
        <el-col :span="6"><div ></div></el-col>
        <el-col :span="6"><div >参与社团：</div> </el-col>
        <el-col :span="6"><div >	acm</div> </el-col>
        <el-col :span="6"><div ></div> </el-col>
      </el-row>
      <el-row>
        <el-col :span="6"><div ></div></el-col>
        <el-col :span="6"><div >兴趣爱好：</div> </el-col>
        <el-col :span="6"><div >音游</div> </el-col>
        <el-col :span="6"><div ></div> </el-col>
      </el-row>
      <el-row>
        <el-col :span="6"><div ></div></el-col>
        <el-col :span="6"><div >座右铭：</div> </el-col>
        <el-col :span="6"><div >一得阁拉米，阿米诺斯</div> </el-col>
        <el-col :span="6"><div ></div> </el-col>
      </el-row >
      <el-row>
        <el-col :span="8"></el-col>
        <el-col :span="8" align="center">

          <p id="word" >
            活跃网站:
            <a href="https://codeforces.com/profile/ManInM00N" style="color: palevioletred;">CodeForces</a>|
            <a href="https://www.luogu.com.cn/user/448777" style="color: palevioletred;">LuoGu</a>|
            <a href="https://atcoder.jp/users/ManInM00N" style="color: palevioletred;">Atcoder</a>|
            <a href="https://Github.com/" style="color: palevioletred;">Github</a>
          </p>
        </el-col>
      </el-row>
    </div>
    </section>
  </el-main>

</template>

<script setup>
// const url = '../assets/images/tx.jpg';


</script>

<style scoped>
.information {

  width: 650px;
  height: 600px;
  background-color:rgba(255, 255, 255, 0.25);
  /*margin: auto;*/
  margin-top: 5%;
  /*text-align: center;*/
  border-radius: 20px;
  padding: 25px 25px;
  padding-bottom: 0px;
  display: block;
  margin-left: calc(50% - 325px);
  margin-right: calc(50% - 325px);
}
.middle_problem{
  float:left;
  text-align: left;
  margin-left: 2%;
  margin-right: 2%;
  width: 90%;
  height: 700px;
  line-height:25px;
  left :10%;
  margin-bottom: 10%;
}
.middle_problem table{
  float: left;
  width: 100%;

  /*left: 10%;*/

}
.middle_problem #word{
  float: left;
  border: none;
  /* left: 10%; */
  margin-left: 25%;
}
.middle_problem #tip{
  //color: azure;
  border: none;
  font-size: 12px;
}
</style>