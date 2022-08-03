<template>
  <el-table :data="list" style="width: 100%;padding-top: 15px;">
    <el-table-column label="Redis Name" min-width="200">
      <template slot-scope="scope">
        {{ scope.row.order_no | orderNoFilter }}
      </template>
    </el-table-column>
    <el-table-column label="Model" width="195" align="center">
      <template slot-scope="scope">
        {{ scope.row.price }}
      </template>
    </el-table-column>
    <el-table-column label="Version" width="120" align="center">
      <template slot-scope="scope">
        {{ scope.row.version }}
      </template>
    </el-table-column>
    <el-table-column label="Nodes" width="120" align="center">
      <template slot-scope="scope">
        {{ scope.row.nodes }}
      </template>
    </el-table-column>
    <el-table-column label="Master" width="120" align="center">
      <template slot-scope="scope">
        {{ scope.row.master }}
      </template>
    </el-table-column>
    <el-table-column label="Slave" width="120" align="center">
      <template slot-scope="scope">
        {{ scope.row.slave }}
      </template>
    </el-table-column>
    <el-table-column label="Total Momery" width="120" align="center">
      <template slot-scope="scope">
        {{ scope.row.totalmomery }}
      </template>
    </el-table-column>
    <el-table-column label="Environment" width="120" align="center">
      <template slot-scope="scope">
        {{ scope.row.environment }}
      </template>
    </el-table-column>
    <el-table-column label="From" width="120" align="center">
      <template slot-scope="scope">
        {{ scope.row.from }}
      </template>
    </el-table-column>
    <el-table-column label="Create Time" width="120" align="center">
      <template slot-scope="scope">
        {{ scope.row.createtime }}
      </template>
    </el-table-column>
    <!-- <el-table-column label="Version" width="195" align="center">
      <template slot-scope="scope">
        ¥{{ scope.row.price | toThousandFilter }}
      </template>
    </el-table-column>
    <el-table-column label="Master" width="195" align="center">
      <template slot-scope="scope">
        ¥{{ scope.row.price | toThousandFilter }}
      </template>
    </el-table-column>
    <el-table-column label="Nodes" width="195" align="center">
      <template slot-scope="scope">
        ¥{{ scope.row.price | toThousandFilter }}
      </template>
    </el-table-column> -->
    <el-table-column label="Status" width="100" align="center">
      <template slot-scope="{row}">
        <el-tag :type="row.status | statusFilter">
          {{ row.status }}
        </el-tag>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
// import { transactionList } from '@/api/remote-search'
import { transactionList } from '@/api/table'
export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        health: 'success',
        fail: 'danger'
      }
      return statusMap[status]
    },
    orderNoFilter(str) {
      return str.substring(0, 30)
    }
  },
  data() {
    return {
      list: null
    }
  },
  created() {
    this.fetchData()
  },
  methods: {    
    // fetchData() {
    //   this.listLoading = true
    //   getList().then(response => {
    //     this.list = response.data.items
    //     this.listLoading = false
    //   })
    // }
    fetchData() {
      transactionList().then(response => {
        this.list = response.data.items.slice(0, 8)
      })
    }
  }
}
</script>
