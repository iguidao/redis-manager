<template>
  <div class="content">
    <div>
        <el-card shadow="never">
        <el-row>
          <el-col :span="3">
              <span>自建codis集群</span>
          </el-col>
          <el-col :span="2">
            <el-select v-model="codisurl" placeholder="请选择平台地址">
              <el-option 
                v-for="item in codismanager" 
                :key="item.Cname" 
                :label="item.Cname" 
                :value="item.Curl" />
            </el-select>
          </el-col>
          <el-col :offset="1" :span="10">
            <el-link :href="codisurl" target="_blank"  type="success">{{ codisurl }} </el-link>
          </el-col>
          <el-col :offset="5"  :span="3" style="min-width: 120px">
            <el-button size="small" type="primary" @click="dialogFormVisible = true">添加codis平台</el-button>
          </el-col>  
        </el-row>
      </el-card>
    </div>
    <div>
        <el-card shadow="never" >
          <iframe :src="codisurl" frameborder="0" width="100%" height="100%" class="codis_dashboard"></iframe>
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
                <el-button type="primary" @click="handleChange()">Confirm</el-button>
            </span>
            </template>
        </el-dialog>
    </div>
  </div>

</template>

<script lang="ts" setup>
import { onMounted, ref, reactive} from 'vue';
import { addCodisCfg, listCodis, listCodisCluster } from '../../../api/codis'
// import moment from 'moment';
import { ElMessage } from 'element-plus';

// 配置列表
const codismanager = ref<any[]>([])
const codisurl = ref("")
// 弹窗
const dialogFormVisible = ref(false)
const formLabelWidth = '100px'
const formcodis = reactive({
  curl: '',
  cname: '',
})
const fromcluster = reactive({
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
