<template>
  <div class="content">
   <el-card shadow="never">
     <el-row>
       <el-col :span="3">
         <span>腾讯Redis集群</span>
       </el-col>
       <el-col :span="2">
        <el-select v-model="txregion" placeholder="选择腾讯Region" @change="getTxRedisCluster()">
            <el-option 
              v-for="item in txredisregion" 
              :key="item.RegionName" 
              :label="item.RegionName" 
              :value="item.Region" />
          </el-select>
        </el-col>
       <!-- <el-col :offset="15"  :span="3" style="min-width: 120px">
         <el-button size="small" type="primary" @click="dialogFormVisible = true">新建集群</el-button>
       </el-col> -->
     </el-row>
     <el-divider></el-divider>
     <el-table :data="txrediscluster" v-loading="loading" stripe style="width: 100%">
       <el-table-column prop="InstanceId" label="ID" width="140"/>
       <el-table-column prop="InstanceName" label="名称" width="200" />
       <el-table-column prop="PrivateIp" label="内网IP" width="120" />
       <el-table-column prop="Port" label="端口" width="65" />
       <el-table-column prop="Size" label="容量" width="80" />
       <el-table-column prop="InstanceStatus" label="状态" width="80" />
       <el-table-column prop="RedisShardSize" label="分片大小" width="80" />
       <el-table-column prop="RedisShardNum" label="分片数量" width="80" />
       <el-table-column prop="RedisReplicasNum" label="副本数" width="80" />
       <el-table-column prop="NoAuth" label="鉴权" width="80">
        <template #default="scope">
          <span v-if="scope.row.NoAuth">否</span>
          <span v-if="!scope.row.NoAuth">是</span>
        </template>
       </el-table-column>
       <el-table-column prop="Password" label="密码" width="120">
        <template #default="scope">
          <div v-if="!scope.row.NoAuth">
            <span v-if="scope.row.Password===''">未设置密码</span>
            <span v-else>已设置密码</span>
          </div>
          <div v-else>
            <span v-if="scope.row.Password===''">免密登陆</span>
          </div>
        </template>
       </el-table-column>
       <el-table-column prop="PublicIp" label="外网IP" width="120" />
       <el-table-column prop="Createtime" :formatter="dateFormat" label="创建时间" width="150"/>
       <el-table-column label="操作" width="300">
         <template #default="scope">
           <el-button size="small" @click="handleEditPw(scope.$index, scope.row)">改密</el-button>
           <el-button size="small" type="primary" @click="handleEditShard(scope.$index, scope.row)">扩缩</el-button>
           <el-popconfirm
             title="确定要删除Redis吗？" 
             confirm-button-text="确认" 
             cancel-button-text="取消" 
             confirm-button-type="danger" 
             cancel-button-type="primary" 
             @confirm="handleDelete(scope.row)" >
             <template #reference>
               <el-button size="small" type="danger">删除</el-button>
             </template>
           </el-popconfirm>
         </template>
       </el-table-column>
     </el-table>
   </el-card>
   <el-dialog v-model="dialogFormVisiblePassword" title="更改密码" width="30%" align-center>
        <el-form :model="formpw">
          <el-form-item label="名称" :label-width="formLabelWidth">
              <el-select v-model="formpw.instanceid" placeholder="请选择要更改的实例">
                <el-option 
                  v-for="item in txrediscluster" 
                  :key="item.InstanceName" 
                  :label="item.InstanceName" 
                  :value="item.InstanceId" />
              </el-select>
            </el-form-item>
         <el-form-item label="密码" :label-width="formLabelWidth">
             <el-input v-model="formpw.password" autocomplete="off" />
         </el-form-item>
       </el-form>
       <template #footer>
       <span class="dialog-footer">
           <el-button @click="handleCancel()">取消</el-button>
           <el-button type="primary" @click="handleChangePw()">确定</el-button>
       </span>
       </template>
   </el-dialog>
   <el-dialog v-model="dialogFormVisible" title="配置变更" width="30%" align-center>
       <el-form :model="formshard">
         <el-form-item label="配置项" :label-width="formLabelWidth">
             <el-select v-model="formshard.key" placeholder="请选择需要变更的配置">
               <el-option 
                 v-for="item in options" 
                 :key="item.value" 
                 :label="item.label" 
                 :value="item.value" />
             </el-select>
         </el-form-item>
         <el-form-item label="配置值" :label-width="formLabelWidth">
             <el-input v-model="formshard.value" autocomplete="off" />
         </el-form-item>
       </el-form>
       <template #footer>
       <span class="dialog-footer">
           <el-button @click="handleCancel()">取消</el-button>
           <el-button type="primary" @click="handleChangeShard()">确定</el-button>
       </span>
       </template>
   </el-dialog>
 </div>
</template>

<script lang="ts" setup>
import { onMounted, ref, reactive} from 'vue';
import { listCloudRegion, listCloudRedis, changeCloudRedisPw } from '../../../api/cloud'
import { listCfg, listDefaultCfg, delCfg, updateCfg } from '../../../api/setting';
import moment from 'moment';
import { ElMessage } from 'element-plus';

// 配置列表
const loading = ref(false)
const dateFormat = (row:any, column:any) => {
 const date = row[column.property]
 if (date===undefined) {
   return ''
 }
 return moment(date).utcOffset(8).format('YYYY-MM-DD HH:mm')
}
//腾讯操作
const txregion = ref("")
const txredisregion = ref<any[]>([])
const txrediscluster = ref<any[]>([])
const fromtxredis = reactive({
  cloud: 'txredis',
  region: '',
})
// 弹窗
//密码弹窗
const dialogFormVisiblePassword= ref(false)
const formpw = reactive({
  cloud: 'txredis',
  instanceid: '',
  password: '',
})


//old
const props = {
 expandTrigger: 'hover' as const,
}
const options = ref<any[]>([])
const dialogFormVisible = ref(false)

const formLabelWidth = '100px'
const formshard = reactive({
 key: '',
 value: '',
})

// 添加删除

// 数据请求
// 腾讯请求
const getTxRedisCluster = async () => {
  loading.value = true
  fromtxredis.region = txregion.value
  let result = (await listCloudRedis(fromtxredis)).data
  if (result.errorCode === 0 ) {
    loading.value = false
    txrediscluster.value = result.data.redis_list
  } else {
    loading.value = false
    ElMessage.error(result.msg)
  }
}
const load = async () => {
  let result = (await listCloudRegion(fromtxredis)).data
  if (result.errorCode === 0 ) {
    txredisregion.value = result.data.region_list
  } else {
    ElMessage.error(result.msg)
  }  
}

const handleEditPw = async (id:number, val:any) => {
  dialogFormVisiblePassword.value = true
  formpw.instanceid = val.InstanceId
  formpw.password = val.Password
}

const handleChangePw  = async () => {
  if (!formpw.instanceid) {
   ElMessage.error("未获取到的变更集群")
 } else if (!formpw.password) {
   ElMessage.error("未获取到的变更密码")
 } else {
    const res = await changeCloudRedisPw(formpw)
    if (res.data.errorCode === 0) {
      ElMessage.success("密码保存成功")
      formpw.instanceid = ""
      formpw.password = ""
      dialogFormVisiblePassword.value = false
      await getTxRedisCluster()
    } else {
      ElMessage.error(res.data.msg)
    }
    dialogFormVisible.value = false
 }
  
}
//old

const handleEditShard = (id:number, val:any) => {
  dialogFormVisible.value = true
  formshard.key = val.Key
  formshard.value = val.Value
}

const handleDelete = async (val:any) => {
 if (!val.Key) {
   ElMessage.error("未获取到的删除的内容")
 } else {
  formshard.key = val.Key
   const res = await delCfg(formshard)
   if (res.data.errorCode === 0) {
     ElMessage.success("删除成功")
     await load()
   } else {
     ElMessage.error(res.data.msg)
   }
 }
}
const handleChangeShard = async () => {
 console.log(formshard)
 if (!formshard.key) {
   ElMessage.error("未获取到的变更配置项")
 } else if (!formshard.value) {
   ElMessage.error("未获取到的变更配置值")
 } else {
   const res = await updateCfg(formshard)
   if (res.data.errorCode === 0) {
     ElMessage.success("变更成功")
     formshard.key = ""
     formshard.value = ""
     await load()
   } else {
     ElMessage.error(res.data.msg)
   }
   dialogFormVisible.value = false
 }
}
const handleCancel = () => {
  dialogFormVisible.value = false
  formshard.key = ""
  formshard.value = ""
  dialogFormVisiblePassword.value = false
  formpw.instanceid = ""
  formpw.password = ""
}
// 启动执行
onMounted(async () => {
 await load()
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
 width: 300px;
}
.el-input {
 width: 300px;
}
.dialog-footer button:first-child {
 margin-right: 10px;
}
</style>