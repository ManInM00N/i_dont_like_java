<template>
  <section class="middle_title" >
    <h1 style=" margin: auto; text-align: center; margin-top: 10px;font-size: 30px;">个人介绍</h1>
  </section>
  <section class="extitile">
  <section class="outside">
    <div class="rectangle-container" >
      <p> 现在在你面前的是 </p>
      <p> Monster Hunter World 1200h+ 全成就白金 </p>
      <p> 黄金奖杯49分58秒记录拥有者 </p>
      <p> 2042大师雌鹿出门30秒撞山记录拥有者 </p>
      <p> 高强度雀魂2000h+ </p>
      <p> 曾经从雀豪1连输26把俯冲到雀杰2的史诗级冲击波选手 </p>
      <h1 style="padding-left:20% ;"><p> ManInM00N</p></h1>
      <br>
      <br>
      <br>
      <br>
      <br>
      <img src="../assets/images/nt1.gif ">
      <div class="trapezoid">
                    <span class="innerP">
                        <br>
                        <p>10人对局9个回合8颗人头7次TK老6整局  </p>
                        <p>🧑‍🍳现场烹饪🧑‍🍳</p>
                        <br>
                        <p>5起沙鹰4局马枪3个背身2人回头1发不中</p>
                        <p>🦀菜就多练🦀 </p>
                        <br>
                        <p> ♥ 伟大无需多言 ♥ </p>
                    </span>
        <br>
        <br>
        <br>
        <br>
        <br>
        <br>
        <br>
        <img src="../assets/images/nt2.gif">
      </div>
    </div>
  </section>
  </section>
</template>

<script setup>

</script>

<style lang = 'less' scoped>


.extitile{
  margin-top: 20px;
  height :calc(100% - 60px)
}
.outside {
  border-radius: 30px;
  width: 750px;
  height: 680px;
  overflow: hidden;
  margin-left: auto;
  margin-right: auto;
  margin-top: 0;
  margin-bottom:0;


}
.rectangle-container {
  width:750px;
  height: 100%;
  float: left;
  position: relative;
  background-color: rgba(12, 12, 12,0.35);
  background-image:url("../assets/images/bb2.jpg");
  & p{
    background-clip: text;
    color: transparent;
    background: linear-gradient(to right,red, blue);
    -webkit-background-clip: text;
    font-weight: bold;
  }
}
.innerP{
  width: 200px;
  height: 100%;
  position: relative;
  & p{

    line-height: 26px;
    background-clip: text;
    color: transparent;
    background: linear-gradient(to right,red, blue);
    -webkit-background-clip: text;
  }
}
.trapezoid {
  position: absolute;
  width:750px;
  height: 700px;
  float: right;
  top: 0;
  text-align: right;
  left: 0;
  clip-path: polygon(60% 0, 100% 0, 100% 100%, 40% 100%);
  background-color:rgba(246, 177, 48, 0.34);/* 调整这里的颜色 */
  background-image: url("../assets/images/bb1.jpg");
  /* border-radius: 30px; */
}
@import "../assets/style/variable";
</style>