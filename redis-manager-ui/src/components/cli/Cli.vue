<template>
  <div class="content">
    <div>
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
          <el-select v-model="txregion" v-else-if="redisname === 'txredis'" placeholder="选择腾讯Region" @change="getTxRedisCluster()">
            <el-option 
              v-for="item in txredisregion" 
              :key="item.RegionName" 
              :label="item.RegionName" 
              :value="item.Region" />
          </el-select>
          <el-select v-model="clusterid" v-else-if="redisname === 'cluster'" placeholder="选择集群"  @change="getClusterMaster()">
            <el-option 
              v-for="item in clusterlist" 
              :key="item.Name" 
              :label="item.Name" 
              :value="item.ID" />
          </el-select>
        </el-col>
        <el-col :span="2">
          <el-select v-model="codisname" v-if="redisname === 'codis'" placeholder="选择集群" @change="getCodisGroup()">
            <el-option 
              v-for="item in codiscluster" 
              :key="item" 
              :label="item" 
              :value="item" />
          </el-select>
          <el-select v-model="txredisid" v-else-if="redisname === 'txredis'" placeholder="选择集群">
            <el-option 
              v-for="item in txrediscluster" 
              :key="item.InstanceName" 
              :label="item.InstanceName" 
              :value="item.InstanceId" />
          </el-select>
          <div  v-if="redisname === 'cluster'">
            <el-input v-model="queryname" class="w-10 m-2" v-if="$props.opkey==='query'" placeholder="请输入key名称" />
            <el-input v-model="queryname" class="w-10 m-2" v-else-if="$props.opkey==='del'" placeholder="要删除的key名称" />
            <el-select v-model="masterid" v-else  placeholder="选择Master">
              <el-option 
                v-for="item in clustermasterlist" 
                :key="item.Ip+':'+item.Port" 
                :label="item.Flags + ': '+ item.SlotRange" 
                :value="item.NodeId" />
            </el-select>
          </div>
        </el-col>
        <el-col  :span="5">
          <div  v-if="redisname === 'codis'">
            <el-input v-model="queryname" class="w-10 m-2" v-if="$props.opkey==='query'" placeholder="请输入key名称" />
            <el-input v-model="queryname" class="w-10 m-2" v-else-if="$props.opkey==='del'" placeholder="要删除的key名称" />
            <el-select v-model="groupname" v-else  placeholder="选择Group">
              <el-option 
                v-for="item in codisgroup" 
                :key="item" 
                :label="item" 
                :value="item" />
            </el-select>
          </div>
          <div  v-else-if="redisname === 'txredis'">
            <el-input v-model="queryname" class="w-10 m-2" v-if="$props.opkey==='query'" placeholder="请输入key名称" />
            <el-input v-model="queryname" class="w-10 m-2" v-else-if="$props.opkey==='del'" placeholder="要删除的key名称" />
          </div>
        </el-col>
        <el-col  :offset="8" :span="4">
          <el-popconfirm
              v-if="$props.opkey==='del'" 
              title="确定要删除吗？" 
              confirm-button-text="确认" 
              cancel-button-text="取消" 
              confirm-button-type="danger" 
              cancel-button-type="primary" 
              @confirm="operationkey()">
              <template #reference>
                <el-button  type="danger">删除</el-button>
              </template>
            </el-popconfirm>
          <el-button type="primary" v-else @click="operationkey()" >查询</el-button>
        </el-col>
      </el-row>
    </div>
      <el-divider></el-divider>
      <div>
        <JsonViewer :value="jsonData" expand-depth="2" copyable boxed sort/>
      </div>
  </div>

</template>
  
<script lang="ts" setup>

import { onMounted, ref, reactive} from 'vue';
import { listCodis, listCodisCluster, listCodisGroup } from '../../api/codis'
import { listCloudRegion, listCloudRedis } from '../../api/cloud'
import { listCluster, listClusterMaster } from '../../api/cluster'
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
    label: '自建cluster集群',
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
// codis
const codisurl = ref("")
const codisname = ref("")
const codismanager = ref<any[]>([])
const codiscluster = ref<any[]>([])
const codisgroup = ref<any[]>([])
const fromcodis = reactive({
  curl: '',
  cname: '',
  cluster_name: '',
})

// txredis
const txregion = ref("")
const txredisid = ref("")
const txredisregion = ref<any[]>([])
const txrediscluster = ref<any[]>([])
const fromtxredis = reactive({
  cloud: '',
  region: '',
})
//cluster
const clusterid = ref("")
const masterid = ref("")
const clusterlist = ref<any[]>([])
const clustermasterlist  = ref<any[]>([])
const fromcluster = reactive({
  cluster_id: '',
})
// all
const groupname = ref("")
const queryname = ref("")
const queryfrom = reactive({
  cache_type: '',
  cache_op: '',
  cluster_name: '',
  key_name: '',
  codis_url: '',
  group_name: '',
  region: '',
	instance_id: '',
  cluster_id: '',
  node_id: '',
})
const queryresult = ref<any>()
const jsonData = reactive(queryresult);


// 数据请求
// 操作key
const operationkey = async () => {
  queryfrom.cache_type = redisname.value
  queryfrom.cache_op = props.opkey
  queryfrom.cluster_name = codisname.value
  queryfrom.key_name = queryname.value
  queryfrom.codis_url = codisurl.value
  queryfrom.group_name = groupname.value
  queryfrom.region = txregion.value
  queryfrom.instance_id = txredisid.value
  queryfrom.cluster_id = String(clusterid.value)
  queryfrom.node_id = masterid.value
  let result = (await cliRedisOpkey(queryfrom)).data
  if (result.errorCode === 0 ) {
    queryresult.value = result.data
  } else {
    ElMessage.error(result.msg)
  }

}
// codis请求
const getCodisCluster = async () => {
  codisname.value = ""
  fromcodis.curl = codisurl.value
  let result = (await listCodisCluster(fromcodis)).data
  if (result.errorCode === 0 ) {
    codiscluster.value = result.data
  } else {
    ElMessage.error(result.msg)
  }
}
const getCodisGroup = async () => {
  fromcodis.curl = codisurl.value
  fromcodis.cluster_name = codisname.value
  let result = (await listCodisGroup(fromcodis)).data
  if (result.errorCode === 0 ) {
    codisgroup.value = result.data
  } else {
    ElMessage.error(result.msg)
  }
}

// txredis 请求
const getTxRedisCluster = async () => {
  fromtxredis.cloud = redisname.value
  fromtxredis.region = txregion.value
  let result = (await listCloudRedis(fromtxredis)).data
  if (result.errorCode === 0 ) {
    txrediscluster.value = result.data.redis_list
  } else {
    ElMessage.error(result.msg)
  }
}
//cluster请求
const getClusterMaster = async () => {
  fromcluster.cluster_id = clusterid.value
  let result = (await listClusterMaster(fromcluster)).data
  if (result.errorCode === 0 ) {
    clustermasterlist.value = result.data
  } else {
    ElMessage.error(result.msg)
  }
}
// 选择codis/cluster/txredis/aliredis集群
const getRedisName = async () => {
  if ( redisname.value == "codis" ) {
    let result = (await listCodis()).data
    if (result.errorCode === 0 ) {
      codismanager.value = result.data.lists
    } else {
      ElMessage.error(result.msg)
    }
  } else if ( redisname.value == "txredis" ) {
    fromtxredis.cloud = redisname.value
    let result = (await listCloudRegion(fromtxredis)).data
    if (result.errorCode === 0 ) {
      txredisregion.value = result.data.region_list
    } else {
      ElMessage.error(result.msg)
    }  
  } else if ( redisname.value == "cluster" ) {
    let result =  (await listCluster()).data
    if (result.errorCode === 0 ) {
      clusterlist.value = result.data
    } else {
      ElMessage.error(result.msg)
    }  
  }
}
//腾讯云redis请求

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