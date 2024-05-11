
<template>

  <div  class="information" >
    <el-main>
      <div v-for="(item,idx) in dialog" :key="idx" style="display: flex">
        <div class="dialog">
          {{item.name}}

        </div>
        <div class="chatBox-left chatBox">
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
import {onBeforeUnmount, ref} from "vue";
import {ElMessage} from "element-plus";


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
  ws.value = new WebSocket("ws://10.5.51.9:7234/api/ws?name="+form.value.name.toString());

  ws.value.onopen = () => {
    console.log('WebSocket connected');
  };

  ws.value.onmessage = (event) => {
    // res.value = event.data;
    handleMessage(JSON.parse(event.data));
  };
  ws.value.onclose = () => {
    ElMessage.error("远程主机关闭")
    console.log('WebSocket closed');
  };

  ws.value.onerror = (error) => {
    console.error('WebSocket error:', error);
  };
};
const handleMessage = (data) => {
  if (data && data.type === 2 && data.data) {
    dialog.value=data.data;
  }else if (data.type === 3 && data.data){
    dialog.value.push(data.data);
  }
  console.log(data.type, data.data,data);
  console.log(dialog.value);
};
onBeforeUnmount(()=>{
  if (ws.value) {
    ws.value.close();
  }
})
function comment(){
  if (!logined.value){
    centerDialogVisible.value = true;
    return
  }
  if (form.value.name === ''){
    form.value.name ="匿名";
  }
  ws.value.send(
      JSON.stringify({
        type:3,
        data:{
          name:form.value.name,
          inner:inner.value,
        },
      })
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
  width: 60px;
  //background: @sliver;
  margin: 5px 5px 5px 5px;
}
.chatBox-left::before{
  content: '';
  position: absolute;
  width: 0;
  height: 0;
  left: -20px;
  top:5px;
  border: 10px solid;
  border-color: transparent #bc3b4a transparent transparent ;
}
.chatBox{
  position: relative;
  margin:12px;
  padding:5px 8px;
  word-break: break-all;
  background: #ffffff;
  border: 1px solid #989898;
  border-radius: 5px;
  max-width:180px;
}
</style>