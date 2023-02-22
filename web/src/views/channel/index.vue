<template>
  <div class="wk-body wk-channel">
    <AddChannelDialog :dialogVisible="addChannelVisible" :dataSource="currentRow" :handleClose="handleAddChannelClose"></AddChannelDialog>
    <div class="wk-channel-header">
      <el-button size="small" @click="currentRow=null;addChannelVisible=true">新增通道</el-button>
    </div>
    <div class="wk-channel-container">
      <el-table :data="channels"> 
        
      </el-table>
      <el-pagination>

      </el-pagination>
    </div>
  </div>
</template>
<script>
import AddChannelDialog from "./components/AddChannelDialog.vue";

export default {
  components: { AddChannelDialog },
  data() {
    return {
      addChannelVisible: false,
      currentRow: null,
      channels: [],
    };
  },
  methods: {
    handleAddChannelClose() {

    }
  }
};
</script>