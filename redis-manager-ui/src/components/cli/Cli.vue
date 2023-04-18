<template>
  <div class="content">
      <el-row class="row-bg">
        <el-col :span="1">
            <span>集群：</span>
        </el-col>
        <el-col :span="2">
          <el-select v-model="redisname" placeholder="Redis集群">
            <el-option 
              v-for="item in redisdic" 
              :key="item.value" 
              :label="item.label" 
              :value="item.value" />
          </el-select>
        </el-col>
        <el-col :span="2">
          <el-select v-model="codisurl" v-if="redisname === 'codis'" placeholder="选择codis平台" @change="getCodisCluster()">
            <el-option 
              v-for="item in codismanager" 
              :key="item.Cname" 
              :label="item.Cname" 
              :value="item.Curl" />
          </el-select>
        </el-col>
        <el-col :span="2">
          <el-select v-model="codisname" v-if="redisname === 'codis'" placeholder="选择集群">
            <el-option 
              v-for="item in codiscluster" 
              :key="item" 
              :label="item" 
              :value="item" />
          </el-select>
        </el-col>
        <el-col :offset="15" :span="2" style="min-width: 120px">
          <el-button  type="primary">查询</el-button>
        </el-col>  
      </el-row>
  </div>
    <div class="box">
        <h4>普通</h4>
            <JsonViewer :value="jsonData" copyable boxed sort theme="light"  @onKeyClick="keyClick"/>
        <h4>暗黑</h4>
            <JsonViewer :value="jsonData" copyable boxed sort theme="dark"  @onKeyClick="keyClick"/>
    </div>
</template>
  
<script lang="ts" setup>
// import JsonViewer from 'vue3-json-viewer'
// import 'vue3-json-viewer/dist/index.css';
import { onMounted, defineProps, ref, reactive} from 'vue';
import { listCodis, listCluster } from '../../api/codis'
// import moment from 'moment';
import { ElMessage } from 'element-plus';

// 配置列表
const redisname = ref("")
const redisdic = [
  {
    value: 'codis',
    label: '自建codis集群',
  },
  {
    value: 'cluster',
    label: '自建redis集群',
  },
  {
    value: 'aliredis',
    label: '阿里redis集群',
  },
  {
    value: 'txredis',
    label: '腾讯redis集群',
  }
]
const codismanager = ref<any[]>([])
const codisurl = ref("")
const codisname = ref("")
const fromcodis = reactive({
  curl: '',
  cname: '',
})
const codiscluster = ref<any[]>([])
let obj = {
  name: "qiu",//字符串
  age: 18,//数组
  isMan:false,//布尔值
  date:new Date(),
  fn:()=>{},
  arr:[1,2,5],
  reg:/ab+c/i
};
const jsonData = reactive(obj);
const keyClick = ()=>{
  console.log("被点击了")
}

// 数据请求
// codis请求
const getCodisCluster = async () => {
  codisname.value = ""
  fromcodis.curl = codisurl.value
  let result = (await listCluster(fromcodis)).data
  if (result.errorCode === 0 ) {
    codiscluster.value = result.data
  } else {
    ElMessage.error(result.msg)
  }
}

// 启动执行
const load = async () => {
  let result = (await listCodis()).data
  if (result.errorCode === 0 ) {
    codismanager.value = result.data.lists
  } else {
    ElMessage.error(result.msg)
  }
}
onMounted(async () => {
  await load()
})
// 接收父参
const props = defineProps({
    opkey: {
        type: String,
        default: String,
        required: true,
    }
})
</script>

<style lange="scss" scoped>
.content {
margin: 20px 8px;
}

.el-button--text {
 margin-right: 15px;
}
.el-select {
 width: 130px;
}
.el-input {
 width: 300px;
}
.cli_commend {
  width: 100%;
  height: 70vh;
  /* height: auto; */
}

.box{
  margin-top: 1rem;
}
</style>