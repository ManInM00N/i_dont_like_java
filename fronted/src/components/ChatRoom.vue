
<template>
  <div  class="information" >
    <el-main>
      <el-divider/>
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
    </el-main>
  </div>
</template>
<script setup>
import {onMounted, ref} from "vue";


const inner = ref('')

const ws = ref(null);
const res = ref('')
const startWebSocket = () => {
  ws.value = new WebSocket('/apis/api/ws');

  ws.value.onopen = () => {
    console.log('WebSocket connected');
  };

  ws.value.onmessage = (event) => {
    res.value = event.data;
  };

  ws.value.onclose = () => {
    console.log('WebSocket closed');
  };

  ws.value.onerror = (error) => {
    console.error('WebSocket error:', error);
  };
};
onMounted(()=>{
  startWebSocket()
})
function comment(){
  ws.value.send(
      {
        type:3,
        message:inner.value,
      }
  )
  inner.value='';
}

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

</style>