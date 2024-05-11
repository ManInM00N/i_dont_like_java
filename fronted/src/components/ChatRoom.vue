
<template>

  <div  class="information" >
    <el-main>
      <div v-for="(item,idx) in dialog" :key="idx">
        {{item.name}}
        <div class="dialog">
          {{item.inner}}
        </div>
      </div>
      <el-divider/>
      <el-card>
        <template #header>
              <el-button @click="comment" tag="发送">
                发送
              </el-button>
        </template>
        <el-input type="textarea" v-model="inner"  placeholder="输入..." >
        </el-input>

      </el-card>
      <el-dialog
          v-model="centerDialogVisible"
          title="用户身份"
          width="500"
          align-center
      >
        <el-form :model="form" :rules="rules">
          <el-form-item label="昵称" label-width="140px" prop="name">
            <el-input v-model="form.name" autocomplete="off" />
          </el-form-item>
        </el-form>
        <template #footer>
          <div class="dialog-footer">
            <el-button @click="centerDialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="centerDialogVisible = false,logined=true,startWebSocket()">
              Confirm
            </el-button>
          </div>
        </template>
      </el-dialog>
    </el-main>
  </div>
</template>
<script setup>
import {ref} from "vue";


const inner = ref('')
const centerDialogVisible = ref(true)
const logined = ref(false)
let ws = WebSocket;
const dialog = ref([])
const form =  ref({
  name:'',
})
const rules = ref({
  name:[
    { required: true, trigger: 'blur' ,pattern: /[a-zA-Z0-9]{4,20}/, message: '名称只能由英文字母数字及下划线组成,且长度为4-20个字符'}
  ],
})
const startWebSocket = () => {
  ws.value = new WebSocket(`/apis/api/ws?name=${form.value.name}`);

  ws.value.onopen = () => {
    console.log('WebSocket connected');
  };

  ws.value.onmessage = (event) => {
    // res.value = event.data;
    dialog.value = event.data;
  };
  ws.value.onclose = () => {
    console.log('WebSocket closed');
  };

  ws.value.onerror = (error) => {
    console.error('WebSocket error:', error);
  };
};
function comment(){
  if (!logined.value){
    centerDialogVisible.value = true;
    return
  }
  if (form.value.name === ''){
    form.value.name ="匿名";
  }
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
@import "../assets/style/Color.less";
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
.dialog{
  display:inline-block;
  padding:5px;
  background: @sliver
}
</style>