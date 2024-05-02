
<template>
  <section class="inmid" style="display: flex">
  <section class="lists ">
  <el-tabs type="card" stretch="@true" :value="now"  @tab-click="tabClicked" class="isblock tomid" >
    <el-tab-pane label="家乡风景" key=0 >
      <el-header class="middle_title" >
        <h1 align="center" style="font-size:26px">家乡风景</h1>
      </el-header>
      <section style="height: 400px">
        <el-carousel :interval="4000" type="card" height="300px" @change="cg2">
          <el-carousel-item v-for="(item,idx) in images2" :key="idx" :label="item.alt" >
            <el-image :src="item.URL"></el-image>
          </el-carousel-item>
        </el-carousel>
        <div class="de">
          <p id="des" >{{describe2[idx2]}}</p>
        </div>
      </section>
<!--        <el-divider/>-->

<!--        <section  >-->
<!--           <el-row>-->
<!--              <el-col span="10"/>-->
<!--             <el-col span="4">-->
<!--               <el-input v-model="txt1" placeholder="搜索内容">-->
<!--                 <template #append>-->
<!--                    <el-button @click="query1" icon="Search"  />-->
<!--                 </template>-->
<!--               </el-input>-->
<!--             </el-col>-->
<!--             <el-col span="10"/>-->

<!--           </el-row>-->
<!--            <section v-if="check1==='ok'">-->
<!--              <el-row  v-for="(item) in q1" :key="item.url" >-->
<!--                <el-col span="12">-->
<!--                  <el-card style="max-width: 480px">-->
<!--                    <template #header>{{item.name}}</template>-->
<!--                    <img-->
<!--                        :src="require('../assets/images/'+item.url)"-->
<!--                        style="width: 100%"-->
<!--                    />-->
<!--                  </el-card>-->
<!--                </el-col>-->
<!--                <el-col span="12">-->
<!--                  <div class="de">-->
<!--                    <p id="des" >{{item.description}}</p>-->
<!--                  </div>-->
<!--                </el-col>-->
<!--              </el-row>-->
<!--            </section>-->
<!--            <section v-else>-->
<!--              <el-result title="404" sub-title="Sorry, request error"/>-->
<!--            </section>-->
<!--        </section>-->
    </el-tab-pane>
    <el-tab-pane label="家乡美食" key=1>
      <div class="middle_title">
        <h1 align="center" style="font-size:26px">家乡美食</h1>
      </div>
      <section style="height:  400px">
        <el-carousel :interval="4000" type="card" height="300px"   @change="cg1">
          <el-carousel-item v-for="(item,idx) in images" :key="idx" :label="item.alt"  >
            <el-image :src="item.URL"></el-image>
          </el-carousel-item>
        </el-carousel>
        <div class="de">
          <p id="des" >{{describe[idx1]}}</p>
        </div>
      </section>
<!--      <el-divider/>-->

<!--      <section  >-->
<!--        <el-row>-->
<!--          <el-col span="10"/>-->
<!--          <el-col span="4">-->
<!--            <el-input v-model="txt2"  placeholder="搜索内容">-->
<!--              <template #append>-->
<!--                <el-button @click="query2" icon="Search" >-->
<!--                </el-button>-->
<!--              </template>-->
<!--            </el-input>-->
<!--          </el-col>-->
<!--          <el-col span="10"/>-->

<!--        </el-row>-->
<!--        <section v-if="check2==='ok'">-->
<!--          <el-row  v-for="(item) in q2" :key="item.url" >-->
<!--            <el-col span="12">-->
<!--              <el-card style="width: 400px">-->
<!--                <template #header>{{item.name}}</template>-->
<!--                <img-->
<!--                    :src="require('../assets/images/'+item.url)"-->
<!--                    style="width: 100%"-->
<!--                />-->
<!--              </el-card>-->
<!--            </el-col>-->
<!--            <el-col span="12">-->
<!--              <div class="de">-->
<!--                <p id="des" >{{item.description}}</p>-->
<!--              </div>-->
<!--            </el-col>-->
<!--          </el-row>-->
<!--        </section>-->
<!--        <section v-else>-->
<!--          <el-result title="404" sub-title="Sorry, request error"/>-->
<!--        </section>-->
<!--      </section>-->
    </el-tab-pane>

  </el-tabs>
    <el-divider/>

    <section  >
      <el-row>
        <el-col span="10"/>
        <el-col span="4">
          <el-input v-model="txt1" placeholder="搜索内容">
            <template #append>
              <el-button @click="query1" icon="Search"  />
            </template>
          </el-input>
        </el-col>
        <el-col span="10"/>

      </el-row>
<!--      <section v-if="check1.value==='ok'">-->
        <el-row  v-for="(item) in q1" :key="item.url" >
          <el-col span="12">
            <el-card style="max-width: 480px">
              <template #header>{{item.name}}</template>
              <img
                  :src="require('../assets/images/'+item.url)"
                  style="width: 100%"
              />
            </el-card>
          </el-col>
          <el-col span="12">
            <div class="de">
              <p id="des" >{{item.description}}</p>
            </div>
          </el-col>
        </el-row>
<!--      </section>-->
<!--      <section v-else>-->
<!--        <el-result title="404" sub-title="Sorry, request error"/>-->
<!--      </section>-->
    </section>
  </section>
  <section width="100%">
    <div id="vide" width="100%">
      <video  width="400" height="auto"  :src="vid[now]" type="video/mp4" controls="controls">
      </video>
    </div>
  </section>

  </section>

</template>
<script setup>
import {ref, watch} from "vue";
import axios from "axios";
import {useRoute, useRouter} from "vue-router";
// const route = useRouter()
const idx1 = ref(0);
const idx2 = ref(0);
const now = ref(0);
const vid = ref([
    require("@/assets/images/xuanchuan.mp4"),
    require("@/assets/images/meishi.mp4"),
]);
const q1 = ref([])
const check1 = ref("ok")
// const check2  = ref("ok")
// let q2 = reactive([])
const txt1= ref('')
// const txt2 = ref('')
const route = useRoute();
// eslint-disable-next-line no-unused-vars
const router = useRouter();
const searchParam = route.query.search;
// const param = reactive({
//   search : txt1.toString(),
// })
const search = async () => {
  console.log(txt1.value );
  try {
    const response = await axios.get(
        "/apis/HomeTown?search="+
        route.query.search
    );
    check1.value=response.data.message
    console.log(response.data,check1.value==='ok')
    q1.value = response.data.res
  } catch (error) {
    console.error('Error fetching search results:', error);
  }
};

if (searchParam) {
  // txt1.value = searchParam;
  search()
}
watch(route, (to, from) => {
  console.log(txt1.value)

  if (to.query.search !== from.query.search) {
    txt1.value = to.query.search || '';
    // 执行重新加载数据或刷新页面的操作
    console.log(txt1.value)
    search();
  }
});
function query1(){

  router.push("/HomeTown?search="+txt1.value.toString());
  // console.log(txt1.value)
  router.replace  ('/refresh')
  // router.replace("/HomeTown?search="+txt1.value.toString());
}
const describe2 = [
  "湖州太湖旅游度假区是湖州滨湖大城市建设重点打造的滨湖新区，集旅游、购物、休闲、度假、居住为一体的国家级旅游区。在太湖月亮湾上面有一个标志性建筑太湖明珠，就是月亮酒店，是湖州地标性建筑。度假区主要景点有太湖温泉水世界、渔人码头、月亮广场、奥特莱斯、发现岛主题乐园、黄金湖岸、长田漾湿地公园等。",
  "安吉竹博园是一家集竹海观光、竹文化主题体验及科普教育为一体的竹类大观园，是湖州著名的旅游景点之一。竹博园总占地1200余亩，其中竹子分类观赏园区600亩，园内收集有散生、混生、丛生竹种300余种，其中安吉乡土竹种50种，国外引种20余种；主要有热带雨林、玉带桥熊猫欢乐世界、中国竹子博物馆、安吉大小熊猫馆、雨雾广场、玉带桥、古代艺术品展览馆、文明竹迹广场等旅游景点。",
  "莫干山，因春秋末年，吴王阖闾派干将、莫邪在此铸成举世无双的雌雄双剑而得名，与北戴河、庐山、鸡公山并称为中国四大避暑胜地，以竹、云、泉“三胜”和清、静、绿、凉“四优”而驰名中外，以绿荫如海的修竹、清澈不竭的山泉、星罗棋布的别墅、四季各异的迷人风光称秀于江南，享有“江南第一山”之美誉。莫干山还曾入选“湖州十大名片”名单，是湖州响当当的城市名片，也是湖州著名景点之一。",
  "在阳光明媚的春天，带上家人，来中南百草园踏青吧！中南百草园就是个氧吧，大片的绿色，大口大口地呼吸新鲜空气，清清脑，洗洗肺。在小镇上走走停停看看，幸福的时光在不经意中悄悄溜走了。",
  "规划面积1000余亩。园内餐饮、品茶居临湖而建，环境优雅，主推农家特色菜肴。多功能会议厅与仿古商街、客房设施齐全。"
];
var images2 = [
  {URL : require('@/assets/images/feature1.jpg'),alt:'湖州太湖旅游度假区'},
  {URL : require('@/assets/images/feature2.jpg'),alt:'安吉竹博园'},
  {URL : require('@/assets/images/feature3.jpg'),alt:'莫干山'},
  {URL : require('@/assets/images/feature4.jpg'),alt:'中南百草园'},
  {URL : require('@/assets/images/feature5.jpg'),alt:'移沿山生态农庄'}

];
var describe = [
  "湖州粽子是浙江湖州的一道著名传统小吃，和嘉兴粽子齐名，是公认的“粽子之王”。这种粽子是长条形的，形似枕头，又叫枕头粽，馅料有很多种。例如：蛋黄粽，豆沙粽，鲜肉粽，满天星粽、蜜枣粽等，吃着鲜香软糯，清爽不腻，令人无法抗拒，就连大S、金庸等人都对它念念不忘。",
  "千张包子是浙江湖州的一道传统小吃，用千张代替面皮，来包裹馅料，最后制作成一种三角形的包子，里边的馅料包含有虾仁、瘦肉、干贝等，一口咬下去，特别有层次感，能吃到外皮的豆香味和肉馅的饱满细腻，鲜味很足。",
  "湖州大馄饨是浙江湖州的一道传统小吃，这种馄饨皮都是特制的，呈长方形，色泽漂亮，煮熟后吃着有嚼劲。而馄饨的汤用的是白汤，汤底要清澈，搭配上紫菜、葱花、虾皮、蛋丝等配料，吃到嘴里脆、鲜、香，口感富有层次。大馄饨煮熟之后，外在看着白嫩细腻，光润晶莹，有着“水晶元宝”的美称。",
  "长兴干挑面是浙江省湖州长兴县的一道著名面食，它属于干面，有猪油、肥肠、大排等各种特色浇头，再搭配上素鸡、肉圆、咸菜、鸡蛋、爆鱼等配菜。当用筷子挑起面条时，汤汁的滋味裹匀每根面条，吃着干而不涩，滑而不腻，实在是美味，很多外地游客尝过之后，都会对它念念不忘。",
  "烂糊鳝丝是浙江湖州的一道传统名菜，也是浙江十大经典名菜之一，以前还曾是宫廷菜，它主要是用鳝鱼爆炒，再搭配上“五丝”，分别是青椒丝、火腿丝、姜丝、葱丝、虾仁去点缀，最中间凹陷处放上蒜泥，最后用热猪油浇制而成。鳝鱼的鲜美被激发出来，吃的时候用筷子拌一拌，鳝丝和佐料的味道互相融合，入口咸香、滑爽、鲜嫩、香醇，不同层次的味觉在舌尖进行碰撞，吃着令人欲罢不能。",
  "来到练市，怎能不尝一尝当地的特色红烧羊肉呢？这道红烧羊肉腥膻味尽褪，在香料的处理后，看着色泽红亮，入口香醇，汁浓肉香，肥而不腻，酥而不烂，是当地人待客的必备佳肴。"
]
var images = [
  {URL : require('../assets/images/f1.jpg'),alt:'湖州粽子'},
  {URL : require('../assets/images/f2.jpg'),alt:'千张包子'},
  {URL : require('../assets/images/f3.jpg'),alt:'湖州大馄饨'},
  {URL : require('../assets/images/f4.jpg'),alt:'长兴干挑面'},
  {URL : require('../assets/images/f5.jpg'),alt:'烂糊鳝丝'},
  {URL : require('../assets/images/f6.jpg'),alt:'练市红烧羊肉'}
]

function tabClicked(tab) {
  // console.log(tab);
  // console.log(tab.name);
  now.value = tab.index;
  // console.log(now.value);
  // console.log(tab.index);
}
function cg2(now){
  idx2.value = now;
  // console.log(now,prev);
}
function cg1(now){
  idx1.value = now;
  // console.log(now,prev);
}


</script>

<style scoped lang="less">
.el-carousel__item h3 {
  color: #475669;
  opacity: 0.75;
  line-height: 200px;
  margin: 0;
  text-align: center;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}
.lists{
  width:80%;
  height: 100%;
  margin-left:0 ;
}
#des{
  background : linear-gradient(to right,rgb(175, 253, 175) ,yellow);
  -webkit-background-clip: text;
  color: transparent;
  margin-top: 0px;
}
@import "../assets/style/variable.less";
</style>