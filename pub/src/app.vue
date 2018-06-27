<template>
  <div id="app">
    <header>
      <!-- <router-link to="/login">login</router-link>
      <router-link to="/chat">chat</router-link> -->
    </header>
    <el-row :gutter="20">
      <el-col :span="12" :offset="6"><div class="nya-content">
        <transition name="fade"><router-view></router-view></transition>  
      </div></el-col>
    </el-row>   
  </div>
</template>

<style>
  .nya-content{
    margin-top: 100px;
    box-shadow: rgba(0,0,0,0.1) 0 1px 2px; 
  }
</style>


<script>
export default {
  
}
</script>