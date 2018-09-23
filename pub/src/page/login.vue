<template>
  <div class="index"
  v-loading="loading"
  element-loading-text="拼命加载中" >
    <div class="typo-PingFang tit"><i class="el-icon-edit"></i>名称</div>
    <div class="head">
      <img :src="url">
    </div>
    <div class="inn"><el-input placeholder="your name." v-model="name" @keyup.enter.native="next_get" ></el-input></div>
    <div class="next"><el-button size="small" @click="next_get">进入</el-button></div>
  </div>
</template>

<script>
import API from "axios"
import {web_url} from "@/config/web"
export default {
  data() {
    return {
      loading: true,
      url: '',
      name: ''
    }
  },
  methods: {
    // 根据URL获取用户头像
    next_get() {
      if(this.name != "") {
        // login
        var name_reg = /[\/\?\:\@\&\=\+\$\#]/
        if(name_reg.test(this.name)){
          this.$message.error('名字中不能包含特殊字符呢')
          return
        }
        this.$router.push('/chat/'+this.name+'/'+encodeURIComponent(this.url)+'/,')
        return
      }
      this.$message.error('请点填写用户名')
    },
    loadI(isrc) {
      var img = new Image()
      img.src = isrc
      img.onerror = () => {
        img.src = 'http://kicoe.com/favicon.ico'
        this.loading = false
      }
      img.onload = () => {
        this.url = img.src
        this.loading = false
      }
    }
  },
  created: function() {
    var isrc = ''
    if(this.$route.params['url'].substring(0,2) != "m="){
      isrc = this.$route.params['url']
      this.loadI(isrc)
    } else {
      API.get(web_url+'/nya?'+this.$route.params['url']).then(response => {
        isrc = response.data
        this.loadI(isrc)
      })
    }
  }
}
</script>

<style>
.index{
  padding: 30px 70px 50px 70px;
}
.tit {
  padding-bottom: 20px;
  color: #303133;
  font-size: 15px
}
.tit i {
  padding-right: 5px
}
.head, .inn, .next {
  text-align: center
}
.head img{
  width: 30px;
  height: 30px;
  border-radius: 50%;
  padding: 10px 20px;
  cursor: pointer;
}
.inn {
  margin-bottom: 30px;
}
.inn input{
  width: 200px;
}
</style>