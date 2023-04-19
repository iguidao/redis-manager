<template>
  <div class="content">
      <el-row :gutter="1">
        <!-- <el-col :span="1">
          <div class="grid-content">
            <span>集群：</span>
          </div>
        </el-col> -->
        <el-col  :span="2">
          <el-select v-model="redisname" placeholder="Redis集群" @change="getRedisName()">
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
        <el-col  :span="5"  v-if="redisname === 'codis'" >
          <el-input v-model="queryname" class="w-10 m-2" v-if="$props.opkey==='query'" placeholder="要查询的内容" />
          <el-input v-model="queryname" class="w-10 m-2" v-else-if="$props.opkey==='del'" placeholder="要查询的内容" />
          <el-select v-model="codisgroup" v-else  placeholder="选择Group">
            <el-option 
              v-for="item in codiscluster" 
              :key="item" 
              :label="item" 
              :value="item" />
          </el-select>
        </el-col>
        <el-col  :offset="8" :span="4">
          <el-button type="primary" @click="operationkey()">查询</el-button>
        </el-col>
      </el-row>

      <el-divider></el-divider>
      <div>
        <JsonViewer :value="jsonData" expand-depth="1" copyable boxed sort/>
      </div>
  </div>

</template>
  
<script lang="ts" setup>

import { onMounted, ref, reactive} from 'vue';
import { listCodis, listCluster } from '../../api/codis'
import { cliRedisOpkey } from '../../api/cli'
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
const codisgroup = ref("")
const queryname = ref("")
const queryfrom = reactive({
  cache_type: '',
  cache_op: '',
  cluster_name: '',
  key_name: '',
  codis_url: '',
  group_name: '',
})
const queryresult = ref<any>()
const jsonData = reactive(queryresult);


// 数据请求
// 操作key
const operationkey = async () => {
  if ( redisname.value == "codis" ) {
    queryfrom.cache_type = redisname.value
    queryfrom.cache_op = props.opkey
    queryfrom.cluster_name = codisname.value
    queryfrom.key_name = queryname.value
    queryfrom.codis_url = codisurl.value
    let result = (await cliRedisOpkey(queryfrom)).data
    if (result.errorCode === 0 ) {
      queryresult.value = result.data
    } else {
      ElMessage.error(result.msg)
    }
  }

}
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
// 选择集群
const getRedisName = async () => {
  if ( redisname.value == "codis" ) {
    let result = (await listCodis()).data
    if (result.errorCode === 0 ) {
      codismanager.value = result.data.lists
    } else {
      ElMessage.error(result.msg)
    }
  }
}
// 启动执行
const load = async () => {
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
.el-col {
  border-radius: 1px;
}

.grid-content {
  border-radius: 1px;
  min-height: 10px;
}
</style>