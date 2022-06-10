<template>
  <div>
    <el-table
      :data="deviceList"
      @expand-change="handleExpandChange"
      row-key="device_id"
      :expand-row-keys="expandRows"
    >
      <el-table-column type="expand">
        <template slot-scope="scope">
          <SubDeviceList :tableData="scope.row.subList"></SubDeviceList>
        </template>
      </el-table-column>
      <el-table-column prop="device_id" label="设备ID"> </el-table-column>
      <el-table-column prop="host" label="主机IP"> </el-table-column>
      <el-table-column prop="region" label="设备域"> </el-table-column>
      <el-table-column prop="created_on" label="创建时间"> </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getDeviceList, getSubDeviceList } from "@/api";
import SubDeviceList from "./component/SubDeviceList.vue";

export default {
  components: { SubDeviceList },
  data() {
    return {
      deviceList: [],
      expandRows: [],
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
    handleExpandChange(row, expandRows) {
      this.expandRows = expandRows.map((item) => item.device_id);
      getSubDeviceList({ deviceId: row.device_id }).then((res) => {
        this.deviceList = this.deviceList.map((item) => {
          if (item.device_id === row.device_id) {
            item.subList = res.data.list;
          }
          return item;
        });
      });
    },
  },
};
</script>

<style>
</style>