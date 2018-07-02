<template>
  <div class="chat">
    <el-container>
    <el-dialog
      :visible.sync="img_edit"
      width="30%"
      center
      :show-close="false"
      >
      <el-input
        type="textarea"
        :rows="img_list.length"
        placeholder="请输入内容"
        v-model="img_list_string"
        >
      </el-input>
      <div class="img-check">
        <el-button size="mini" type="success" icon="el-icon-check" @click="imgSet()"></el-button>
        <el-button size="mini" type="danger" icon="el-icon-close" @click="img_edit=false"></el-button>
      </div>
    </el-dialog>

    <el-aside width="200px" class="left">
      <div class="i"><img class="im" :src="img"><a class="in">{{name}}</a></div>
      <div @click="checkIt(0)" :class="[to_uid==0?'select':'','home','li']"><i class="ic el-icon-menu"></i><a class="in">群聊</a></div>
      <ul>
        <li @click="checkIt(info[0])" v-for="info in user_list" :class="[to_uid==info[0]?'select':'','li']" ><img class="im" :src="info[2]"><a class="in">{{info[1]}}<el-badge class="msg-bad" :value="info[3]==0?false:info[3]" /></a></li>
      </ul>
    </el-aside>
    <el-main class="right">
      <div class="ch" id="ch">
        <div class="msg" v-for="m in message_list">
          <div :class="[m[0]==uid?'self':'info']"><span class="name">{{m[0]==uid?name:user_map[m[0]][0]}}</span><span class="time">{{m[2]}}</span></div>
          <div class="img_show" v-if="imgParse(m[1])"><img :src="imgParse(m[1])[1]" /></div>
          <div class="text" v-else>{{m[1]}}</div>
        </div>
      </div>
      <div class="chin">
        <el-popover
          placement="top"
          width="265"
          v-model="pop"
          trigger="click">

          <div class="imgs">
            <img @click="imgDump(i)" v-for="i in img_list" :src="i">
          </div>
          <div class="img-edit">
            <el-button @click="imgEdit()" icon="el-icon-edit"></el-button>
          </div>

          <i slot="reference" class="el-icon-picture ic"></i>
        </el-popover>
        <el-input class="inpu" @keyup.enter.native="send" v-model="input" placeholder="请输入内容"></el-input>
        <el-button @click="send" size="small">发送</el-button>
      </div>
    </el-main>
    <i class="el-icon-share share" @click="share"></i>
    </el-container>
  </div>
</template>

<style>
.share{
  height: 10px;
  cursor: pointer
}
.share :hover{
}
.el-dialog__header {
    padding: 0;
}
.el-dialog--center .el-dialog__body {
  padding: 0 0 10px;
}
.el-dialog__headerbtn {
  top: 5px;
  right: 5px;
}
.img-check{
  margin-top: 10px;
  text-align: center
}
.img-edit{
  text-align: center
}
.img-edit button{
  padding: 4px 20px
}
.imgs img{
  padding: 2px;
  width:60px;
  cursor: pointer
}
.img_show{
  padding: 3px 5px
}
.img_show img{
  max-width: 150px
}
.msg{
  font-size: 13px;
  padding-left: 25px;
}
.msg .info{
  color: #409eff
}
.msg .self{
  color: #67c23a
}
.msg .name{
}
.msg .time{
  font-size: 12px;
  margin-left: 6px;
}
.msg .text{
  padding: 4px 6px
}
.el-main{
  padding-left: 40px;
}
.chin {
  padding: 3px 30px
}
.chin .ic{
  padding: 6px;
  cursor: pointer;
}
.chin .ic:hover{
  color: #667;
}
.chin .inpu{
  width: 77%;
  margin: 0 20px;
}
.chin button{
  position: absolute
}
.ch {
  height: 308px;
  position: relative;
  left: 15px;
  top: -10px;
  background-color: #fff;
  overflow: auto;
}
::-webkit-scrollbar { width: 0;height: 0; }
ul{
  padding: 0px;
  margin: 0px;
}
.msg-bad{ float:right }
.chat{
  padding: 24px 20px 12px;
  height: 366px;
}
.left{
  text-overflow:ellipsis;
  overflow: auto;
  height: 350px;
}
.home {
  margin: 10px 0;
}
.home i{
  font-size: 16px;
  position: relative;
  top: 10px;
  left: 4px;
}
.li:hover i { 
  background-color: inherit;
}
.i {
  height: 35px;
  display: block;
  padding-bottom: 10px;
  padding-left: 20px;
  border-bottom: 1px solid #ccc;
  white-space:nowrap;
}
.li {
  height: 35px;
  display: block;
  padding: 4px 10px 3px;
  cursor: pointer;
}
.select{ background-color:#eee }
.select :hover{ background-color:#eee }
.new .in{ color:red!important }
/*?? .li :hover{
  background-color:#eee;
} */
.li .im{
  padding: 2px 5px;
  width: 30px;
  height: 30px;
  border-radius: 50%;
}
.li .ic{
  float: left;
  position: relative;
  top: 12px;
  left: 11px;
}
.li .in{
  color: #555
}
.im {
  width: 35px;
  height: 35px;
  float: left;
  border-radius: 2px
}
.in {
  white-space:nowrap;
  display: block;
  padding: 9px 25px 5px 50px;
  font-size: 15px;
  font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;
}
</style>

<script>
import Config from "../config"
export default {
  data() {
    return {
      uid: 0,
      to_uid: 0,
      name: '',
      img: '',
      input: '',
      ws: null,
      user_map: {
        //id: [name, img]
      },
      user_list: [
        //[id, name, img, news]
      ],
      message_map: {
        // id: [
        //    [id, msg, time]
        // ]
      },
      message_list:[
        // [id, msg, time]
      ],
      img_list: [],
      img_list_string: "",
      img_edit: false,
      pop: false
    }
  },
  //每次页面渲染完之后滚动条在最底部
  updated:function(){
    this.$nextTick(function(){
      var div = document.getElementById('ch')
      div.scrollTop = div.scrollHeight
    })
  },
  methods: {
    share: function(){
      var localhostPaht = window.location.href.substring(0, window.location.href.indexOf(window.location.pathname));
      var url = localhostPaht+'/chat/'+encodeURIComponent(this.name)+'/'+encodeURIComponent(this.img)+'/'+encodeURIComponent(this.img_list.join(','))
      this.$alert('<textarea id="share" class="el-textarea__inner">'+url+'</textarea>', 'link', {
        confirmButtonText: 'Copy',
        dangerouslyUseHTMLString: true
      }).then(() => {
        document.getElementById('share').select()
        document.execCommand("Copy")
        this.$message({
          type: 'success',
          message: '个人URL成功复制到粘贴板'
        })
      }).catch(()=>{})
      
    },
    imgEdit: function(){
      this.img_edit = true
      this.img_list_string = this.img_list_parse
    },
    imgSet: function(){
      this.img_list = this.img_list_string.split("\n")
      this.img_edit = false
    },
    // 图片
    imgDump: function(img_url){
      this.input = '![]('+img_url+')'
      this.pop = false
      // this.send()
    },
    // 图片解析
    imgParse: function(txt){
      return txt.match(/^\!\[\]\((.*?)\)$/)
    },
    // 切换聊天窗口
    checkIt: function(uid) {
      this.to_uid = uid
      this.message_list = this.message_map[uid]
      if(uid!=0){
        this.user_list.forEach((item, index) => {
          if(item[0] == uid){
            this.user_list[index][3] = 0
            return
          }
        })
      }
      //this.$set(this.user_map, uid, [this.user_map[uid][0],this.user_map[uid][1],0])
    },
    // 向某用户发送消息
    send: function() {
      if (this.input == '') return
      //this.message_map[this.to_uid].push([this.uid, this.input])
      this.message_list.push([this.uid, this.input, this.getTime()])
      if(this.to_uid==0)
        this.ws.send('{"c":"broadcast","a":{"u":'+this.to_uid+',"a":["'+this.input+'"]}}')
      else
        this.ws.send('{"c":"send","a":{"u":'+this.to_uid+',"a":["'+this.input+'"]}}')
      this.input = ''
    },
    login: function(data) {
      // 初始化登录
      // {u: uuid, a: [uuid_name_img, uuid_name_img ...]}
      this.uid = data.u
      this.message_map[0] = []
      data.a.forEach((item, index) => {
        var user_info = item.split('_')
        if(this.uid!=user_info[0]){
          // 主动刷新
          // this.$set(this.user_map, user_info[0], [user_info[1], user_info[2]])
          this.user_map[user_info[0]] = [user_info[1], user_info[2]]
          user_info.push(0)
          this.user_list.push(user_info)
          this.message_map[user_info[0]] = []
        }
      })
      this.checkIt(0)
      console.log(data)
    },
    online: function(data){
        // 处理用户上线
        // {u: uuid, a: [name, img]}
        // this.$set(this.user_map, data.u, data.a)
        this.user_map[data.u] = data.a
        this.user_list.push([data.u, data.a[0], data.a[1], 0])
        this.message_map[data.u] = []
        console.log(data)
    },
    offline: function(data){
        // 处理用户下线
        // {u: uuid, a: []}
        if(data.u == this.to_uid){
          this.message_list = []
          this.to_uid = 0
        }
        // this.$delete(this.user_map, data.u)
        this.user_list.forEach((item, index) => {
          if(item[0] == data.u){
            this.user_list.splice(index, 1)
            return
          }
        })
        console.log(data)
    },
    message: function(data){
        // 处理某用户发送消息
        // {u: uuid, a: [massage]}
        if(data.u == this.to_uid){
          this.message_list.push([data.u, data.a[0], this.getTime()])
        } else {
          this.message_map[data.u].push([data.u, data.a[0], this.getTime()])
          this.user_list.forEach((item, index) => {
            if(item[0] == data.u){
              this.user_list.splice(index, 1)
              this.user_list.unshift([item[0], item[1], item[2], item[3]+1])
              return
            }
          })
          //this.$set(this.user_map, data.u, [this.user_map[data.u][0],this.user_map[data.u][1],1])
        }
        console.log(data)
    },
    broadcast: function(data){
        // 处理某用户广播消息
        // {u: uuid, a: [massage]}
        if(0 == this.to_uid){
          this.message_list.push([data.u, data.a[0], this.getTime()])
        } else {
          this.message_map[0].push([data.u, data.a[0], this.getTime()])
        }
        console.log(data)
    },
    getTime: function() {
      // 获得当前时间 h:i
      var t = new Date()
      return t.getHours()+":"+((t.getMinutes()<10)?("0"+t.getMinutes()):t.getMinutes())
    }
  },
  created() {
    if(this.$route.params['f'] == ','){
      this.img_list = Config.img_list
    } else {
      this.img_list = this.$route.params['f'].split(',')
    }
    this.name = this.$route.params['name']
    this.img = decodeURIComponent(this.$route.params['img'])
    // 连接websocket
    this.ws = new WebSocket(Config.ws_url+"/ws?name="+this.name+"&img="+this.$route.params['img'])
    var that = this
    this.ws.onmessage = function(e) {
      console.log(e.data);
      var data = JSON.parse(e.data)
      that[data.e](data.a)
    }
  },
  computed: {
    img_list_parse: function() {
      return this.img_list.join("\n")
    }
  }
}
</script>