<template>
  <div class="index">
    <div class="typo-PingFang tit"><i class="el-icon-picture-outline"></i>获取头像</div>
    <div class="icon">
      <img src="/static/qq.jpg" @click="qq" v-bind:style="sty.qaq">
      <img  src="/static/gv.jpg" @click="ava" v-bind:style="sty.g">
    </div>
    <div class="head-url"><el-input :placeholder="pla" v-model="url" @keyup.enter.native="next_get" ></el-input></div>
    <div class="next"><el-button size="small" @click="next_get">下一步</el-button></div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      url: '',
      pla: '请选择头像格式啦！',
      sty: {
        qaq: {},
        g: {}
      }
    }
  },
  methods: {
    qq() {
      this.pla = 'QQ'
      this.sty.g = {}
      this.sty.qaq = {opacity: 1}
    },
    ava() {
      this.pla = 'Email'
      this.sty.qaq = {}
      this.sty.g = {opacity: 1}
    },
    // 根据URL获取用户头像
    next_get() {
      if(this.url == ''){
        this.$message.error('请填写头像相关信息哦')
        return
      }
      var email_reg = /^([a-zA-Z0-9_\.\-])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})+$/
      var qq_reg = /^[1-9][0-9]{4,14}$/
      if(this.pla == 'QQ') {
        if(qq_reg.test(this.url)) {
          this.$router.push('/login/'+'m=q&u='+this.url)
        } else {
          this.$message.error('请填写正确的qq')
        }
        return
      }
      if(this.pla == 'Email') {
        if(email_reg.test(this.url)) {
          this.$router.push('/login/'+'m=g&u='+this.url)
        } else {
          this.$message.error('请填写正确的邮箱')
        }
        return
      }
      this.$message.error('请点击选择头像格式呢')
    },
    //async 
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
.icon, .head-url, .next {
  text-align: center
}
.icon img{
  width: 30px;
  height: 30px;
  border-radius: 50%;
  padding: 10px 20px;
  cursor: pointer;
}
.icon img{
  opacity: 0.5;
}
.icon img:hover{
  opacity: 0.6;
}
.head-url {
  margin-bottom: 30px;
}
.head-url input{
  width: 200px;
}
</style>