<template>
  <div class="content">
    <div>
      <el-card shadow="never">
        <el-row>
          <el-col :span="3">
            <span>自建Cluster集群</span>
          </el-col>
          <el-col :offset="18"  :span="3" style="min-width: 120px">
            <el-button size="small" type="primary" @click="dialogFormVisible = true">添加集群</el-button>
          </el-col>
        </el-row>
      <el-divider></el-divider>
     <el-table :data="clustertable" stripe style="width: 100%" v-loading="loading" lazy @expand-change="GetNodes">
        <el-table-column  label="查看" type="expand" width="60">
          <template #default="scope">
            <el-table :data="nodetable" 
            :border=true
            style="width: 100%"
            row-key="id"
            lazy
            :load="load"
            :tree-props="{ children: 'Children', hasChildren: 'hasChildren' }"
            >
              <el-table-column prop="Flags" label="身份" width="160" />
              <el-table-column prop="Address" label="地址" width="160" />
              <el-table-column prop="LinkState" label="LinkState" width="160" />
              <el-table-column prop="RunStatus" label="RunStatus" width="160" />
              <el-table-column prop="SlotRange" label="SlotRange" width="160" />
              <el-table-column prop="SlotNumber" label="SlotNumber" width="160" />
              <el-table-column prop="CreateTime" :formatter="dateFormat"  label="创建时间" width="160" />
            </el-table>
          </template>
        </el-table-column>
        <el-table-column prop="ID" label="ID" width="60" />
        <el-table-column prop="Name" label="名称" width="100" />
        <el-table-column prop="Nodes" label="内网IP" />
        <el-table-column prop="Password" label="密码" width="120"  >
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
        <el-table-column prop="CreatedAt" :formatter="dateFormat" label="创建时间" width="150" />
       <!-- <el-table-column label="操作" width="300">
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
       </el-table-column> -->
     </el-table>
    </el-card>
    </div>

    <div>
      <el-dialog v-model="dialogFormVisible" title="添加集群" width="30%" align-center>
          <el-form :model="formcluster">
            <el-form-item label="集群名称"  :label-width="formLabelWidth">
                <el-input v-model="formcluster.name" placeholder="Cluster Name" autocomplete="off" />
            </el-form-item>      
            <el-form-item label="集群地址" :label-width="formLabelWidth">
                <el-input v-model="formcluster.nodes"  placeholder="127.0.0.1:6379,127.0.0.1:6380"  autocomplete="off" />
            </el-form-item>
            <el-form-item label="集群密码"  :label-width="formLabelWidth">
                <el-input v-model="formcluster.password" type="password" placeholder="Cluster Password" autocomplete="off" />
            </el-form-item>
          </el-form>
          <template #footer>
          <span class="dialog-footer">
              <el-button @click="dialogFormVisible = false">取消</el-button>
              <el-button type="primary" @click="handleChange()">Confirm</el-button>
          </span>
          </template>
      </el-dialog>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref, reactive} from 'vue';
import { addClusterCfg, listClusterNodes, listCluster } from '../../../api/cluster';
import { ElMessage, RowProps } from 'element-plus';
import moment from 'moment';

const loading = ref(false)
const formLabelWidth = '100px'
const clustertable = ref<any[]>([])
const nodetable = ref<any[]>([])
const dialogFormVisible = ref(false)
const formcluster = reactive({
  name: '',
  nodes: '',
  password: '',
})
const formnodes = reactive({
  cluster_id: '',
})
const dateFormat = (row:any, column:any) => {
 const date = row[column.property]
 if (date===undefined) {
   return ''
 }
 return moment(date).utcOffset(8).format('YYYY-MM-DD HH:mm')
}
const load = async () => {
  let clusterlist = (await listCluster()).data
  // let codislist = (await listCodis()).data
  if (clusterlist.errorCode === 0 ) {
    loading.value = false
    clustertable.value = clusterlist.data
  } else {
    loading.value = false
    ElMessage.error(clusterlist.msg)
  }
}
const expandChange = async () => {

}
const GetNodes = async (row: any) => {
  formnodes.cluster_id = row.ID
  if (!formnodes.cluster_id) {
   ElMessage.error("未获取到的添加cluster_id")
  } else {
   const res = (await listClusterNodes(formnodes)).data
   if (res.errorCode === 0) {
      nodetable.value = res.data
   } else {
     ElMessage.error(res.msg)
   }
   dialogFormVisible.value = false
 }
}
const handleChange = async () => {
 if (!formcluster.name) {
   ElMessage.error("未获取到的添加cluster集群的名称")
 } else if (!formcluster.nodes) {
   ElMessage.error("未获取到的添加cluster集群的地址")
 } else {
   const res = await addClusterCfg(formcluster)
   if (res.data.errorCode === 0) {
     ElMessage.success("添加成功")
     await load()
   } else {
     ElMessage.error(res.data.msg)
   }
   dialogFormVisible.value = false
 }
}
// 启动执行
onMounted(async () => {
  loading.value = true
  await load()
})
</script>
<style lange="scss" scoped>
.content {
  margin: 20px 8px;
}
</style>