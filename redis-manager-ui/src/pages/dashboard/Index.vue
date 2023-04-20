<template>
  <div class="content">
    <!-- 运营状况 -->
    <div class="main-info">
      <el-card class="info">
        <el-button type="primary" icon="el-icon-user-solid" circle />
        <h2 class="num-info">112</h2>
        <p class="desc">腾讯Redis集群</p>
      </el-card>
      <el-card class="info">
        <el-button type="success" icon="el-icon-s-data" circle />
        <h2 class="num-info">66</h2>
        <p class="desc">阿里Redis集群</p>
      </el-card>
      <el-card class="info">
        <el-button type="danger" icon="el-icon-coin" circle />
        <h2 class="num-info">12</h2>
        <p class="desc">自建Codis集群</p>
      </el-card>
      <el-card class="info">
        <el-button type="warning" icon="el-icon-data-line" circle />
        <h2 class="num-info">99</h2>
        <p class="desc">自建Redis集群</p>
      </el-card>
    </div>
    <!-- 图表 -->
    <div class="chart">
      <!-- <el-card class="e-chart" id="one-chart"></el-card>
      <el-card class="e-chart" id="two-chart"></el-card> -->
    </div>
  </div>
</template>

<script lang="ts" setup>
</script>

<style lange="scss" scoped>
.content {
  margin-top: 20px;
  display: flex;
  flex-direction: column;
  height: 70%;
}
.main-info {
  flex: 1;
  display: flex;
  justify-content: space-around;
  background: rgb(255, 255, 255);
  min-width: 700px;
  border-radius: 8px;
}

.info {
  background: white;
  width: 18%;
  height: 80%;
  align-self: center;
  text-align: center;
  box-sizing: border-box;
  display: flex;
  align-items: center;
  justify-content: center;
}
.info .num-info {
  margin: 10px 0px;
}
.info .desc {
  font-size: 10px;
  color: gray;
}
.chart {
  flex: 1;
  margin-top: 10px;
  display: flex;
  justify-content: space-between;
  border-radius: 8px;
  width: 100%;
  height: 100%;
}

.e-chart {
  padding-top: 10px;
  width: 49%;
  height: auto;
  min-height: 300px;
  min-width: 300px;
}
</style>