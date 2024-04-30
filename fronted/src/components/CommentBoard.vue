<template>
  <section class="middle_title">
    <div class="information">
      <h1 class="Title">留言板</h1>

    </div>
  </section>

</template>
<script setup>

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
  background-color:rgba(255, 255, 255, 0.25);
  margin-top: 5%;
  border-radius: 20px;
  padding: 25px 25px;
  padding-bottom: 0px;
  display: block;
  margin-left: calc(50% - 325px);
  margin-right: calc(50% - 325px);
}
@import "../assets/style/Color.less";
@import "../assets/style/variable.less";
</style>