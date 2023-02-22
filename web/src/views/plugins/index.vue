<template>
  <div class="wk-plugin">
    <div class="wk-plugin-header"></div>
    <div class="wk-plugin-container">
      <el-table :data="channels"> 
        
      </el-table>
      <el-pagination>

      </el-pagination>
    </div>
  </div>
</template>
<script>
import AddPluginDialog from "./components/AddPluginDialog.vue";

export default {
  components: { AddPluginDialog },
  data() {
    return {
      channels: [],
    };
  },
};
</script>