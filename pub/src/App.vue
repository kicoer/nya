<template>
  <div id="app">
    <el-row :gutter="20">
      <el-col :span="12" :offset="6"><div class="nya-content">
        <transition name="fade"><router-view></router-view></transition>  
      </div></el-col>
    </el-row>   
  </div>
</template>

<script>
  export default {
    
  }
</script>

<style lang="scss">
	@import './style/common';
  .nya-content{
    margin-top: 100px;
    box-shadow: rgba(0,0,0,0.1) 0 1px 2px;
  }
</style>