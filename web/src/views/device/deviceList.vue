<template>
  <div>
    <el-table :data="deviceList">
      <el-table-column prop="device_id" label="设备ID"> </el-table-column>
      <el-table-column prop="host" label="主机IP"> </el-table-column>
      <el-table-column prop="region" label="设备域"> </el-table-column>
      <el-table-column prop="created_on" label="创建时间"> </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getDeviceList } from "@/api";

export default {
  data() {
    return {
      deviceList: [],
      queryCondition: {
        pageNumber: 1,
        pageSize: 10,
      },
    };
  },
  created() {
    this.getData();
  },
  mounted() {},
  methods: {
    getData() {
      getDeviceList(this.queryCondition).then((res) => {
        this.deviceList = res.data.list;
      });
    },
  },
};
</script>

<style>
</style>