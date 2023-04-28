<template>
  <div class="content">
    <div>
        <el-card shadow="never">
        <el-row>
          <el-col :span="3">
              <span>自建codis集群</span>
          </el-col>
          <el-col :span="2">
            <el-select v-model="formcodis.curl" placeholder="请选择平台地址" @change="GetCluster()">
              <el-option 
                v-for="item in codismanager" 
                :key="item.Cname" 
                :label="item.Cname" 
                :value="item.Curl" />
            </el-select>
          </el-col>
          <el-col :offset="1" :span="10">
            <el-link :href="formcodis.curl" target="_blank"  type="success">{{ formcodis.curl }} </el-link>
          </el-col>
          <el-col :offset="2"  :span="2" style="min-width: 120px">
            <el-button size="small"  @click="dialogFormVisibleopnode = true">扩缩容</el-button>
          </el-col>  
          <el-col  :span="3" style="min-width: 120px">
            <el-button size="small" type="primary" @click="dialogFormVisible = true">添加codis平台</el-button>
          </el-col>  
        </el-row>
      </el-card>
    </div>
    <div>
        <el-card shadow="never" >
          <iframe :src="formcodis.curl" frameborder="0" width="100%" height="100%" class="codis_dashboard"></iframe>
        </el-card>
    </div>
    <div>
        <el-dialog v-model="dialogFormVisible" title="新增codis平台地址" width="30%" align-center>
            <el-form :model="formcodis">
              <el-form-item label="平台名称" :label-width="formLabelWidth">
                  <el-input v-model="formcodis.cname" autocomplete="off" />
              </el-form-item>      
              <el-form-item label="平台地址" :label-width="formLabelWidth">
                  <el-input v-model="formcodis.curl" autocomplete="off" />
              </el-form-item>
            </el-form>
            <template #footer>
            <span class="dialog-footer">
                <el-button @click="dialogFormVisible = false">取消</el-button>
                <el-button type="primary" @click="handleChange()">确定</el-button>
            </span>
            </template>
        </el-dialog>
        <el-dialog v-model="dialogFormVisibleopnode" title="扩缩容codis集群" width="30%" align-center>
            <el-form :model="formcodis">
              <el-form-item label="操作类型" :label-width="formLabelWidth">
                <el-select v-model="formopnode.op_type" class="m-2" placeholder="操作" size="large">
                  <el-option
                    v-for="item in options"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                  />
                </el-select>
              </el-form-item> 
              <el-form-item label="选择集群" :label-width="formLabelWidth">
                <el-select v-model="formopnode.cluster_name" class="m-2" placeholder="集群" size="large">
                  <el-option
                    v-for="item in clusterlist"
                    :key="item"
                    :label="item"
                    :value="item"
                  />
                </el-select>
              </el-form-item>   
              <div v-if="formopnode.op_type === 'dilatation'">
                <el-form-item label="proxy列表" :label-width="formLabelWidth">
                  <el-input v-model="formopnode.add_proxy" placeholder="127.0.0.1:11081,127.0.0.1:11082" autocomplete="off" />
                </el-form-item>      
                <el-form-item label="Redis列表" :label-width="formLabelWidth">
                    <el-input v-model="formopnode.add_server" placeholder="127.0.0.1:6379,127.0.0.1:6380" autocomplete="off" />
                </el-form-item>
              </div>
              <div v-if="formopnode.op_type === 'shrinkage'">
                <el-form-item label="proxy数量" :label-width="formLabelWidth">
                  <el-input v-model.number="formopnode.del_proxy" placeholder="1" autocomplete="off" />
                </el-form-item>      
                <el-form-item label="group数量" :label-width="formLabelWidth">
                    <el-input v-model.number="formopnode.del_group" placeholder="1" autocomplete="off" />
                </el-form-item>
              </div>
            </el-form>
            <template #footer>
            <span class="dialog-footer">
                <el-button @click="dialogFormVisibleopnode = false">取消</el-button>
                <el-button type="primary" @click="handleChangeOpnode()">确定</el-button>
            </span>
            </template>
        </el-dialog>
    </div>
  </div>

</template>

<script lang="ts" setup>
import { onMounted, ref, reactive} from 'vue';
import { addCodisCfg, listCodis, listCodisCluster, opCodisNode } from '../../../api/codis'
// import moment from 'moment';
import { ElMessage } from 'element-plus';

// 配置列表
const codismanager = ref<any[]>([])
const clusterlist = ref<any[]>([])
// 弹窗
const dialogFormVisible = ref(false)
const dialogFormVisibleopnode = ref(false)
const formLabelWidth = '100px'
const options = [
  {
    value: 'dilatation',
    label: '扩容',
  },
  {
    value: 'shrinkage',
    label: '缩容',
  }
]
const formopnode = reactive({
  curl: '',
  cluster_name: '',
  add_proxy: '',
  add_server: '',
  del_proxy: 0,
  del_group: 0,
  op_type: ''
})
const formcodis = reactive({
  curl: '',
  cname: '',
})

// 数据请求
const load = async () => {
  let codislist = (await listCodis()).data
  if (codislist.errorCode === 0 ) {
    codismanager.value = codislist.data.lists
  } else {
    ElMessage.error(codislist.msg)
  }
}
const GetCluster = async () => {
  const res = await listCodisCluster(formcodis)
  if (res.data.errorCode === 0) {
      clusterlist.value = res.data.data
     await load()
   } else {
     ElMessage.error(res.data.msg)
   }
}
const handleChange = async () => {
 console.log(formcodis)
 if (!formcodis.curl) {
   ElMessage.error("未获取到的新增codis平台地址")
 } else if (!formcodis.cname) {
   ElMessage.error("未获取到的新增codis平台名称")
 } else {
   const res = await addCodisCfg(formcodis)
   if (res.data.errorCode === 0) {
     ElMessage.success("添加成功")
     await load()
   } else {
     ElMessage.error(res.data.msg)
   }
   dialogFormVisible.value = false
 }
}
const handleChangeOpnode = async () => {
  formopnode.curl = formcodis.curl
  if (!formopnode.cluster_name || !formopnode.curl){
    ElMessage.error("未获取到集群名和集群地址")
  } else {
    console.log(formopnode)
    const res = await opCodisNode(formopnode)
    if (res.data.errorCode === 0) {
      ElMessage.success(res.data.data)
      await load()
    } else {
      ElMessage.error(res.data.msg)
      ElMessage.error(res.data.data)
    }
    dialogFormVisibleopnode.value = false
  }
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
 width: 130px;
}
.el-input {
 width: 300px;
}
.dialog-footer button:first-child {
 margin-right: 10px;
}


.el-link {
  margin-right: 8px;
}
.el-link .el-icon--right.el-icon {
  vertical-align: text-bottom;
}
.codis_dashboard {
  width: 100%;
  height: 70vh;
  /* height: auto; */
}
</style>
