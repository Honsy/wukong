<template>
  <div :class="['sub-table', tableData.length <= 0 ? 'empty' : '']">
    <el-table
      v-if="tableData.length > 0"
      :data="tableData"
      style="width: 100%"
      row-class-name="sub-table-row"
    >
      <el-table-column prop="name"></el-table-column>
      <el-table-column prop="address"></el-table-column>
      <el-table-column prop="civilcode"></el-table-column>
      <el-table-column prop="owner"></el-table-column>
      <el-table-column prop="name"></el-table-column>
      <el-table-column prop="modified_on"></el-table-column>
      <el-table-column label="操作" >
        <template slot-scope="scope">
          <el-button type="text" @click="playVideo(scope.row)">启动推流</el-button>
          <el-button type="text">删除</el-button>
          <el-button type="text" @click="stopVideo(scope.row)">停止推流</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { playCameraWithDeviceId } from "@/api";
export default {
  props: {
    tableData: {
      type: Array,
      default() {
        return [];
      },
    },
  },
  methods: {
    playVideo(camera) {
      const params = {
        device_id: camera.device_id
      }
      playCameraWithDeviceId(params).then(res=>{
        console.log(res)
      })
    },
    stopVideo(camera) {
      
    }
  }
}
</script>

<style>

</style>