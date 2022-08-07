<template>
  <div>
    <el-row>
      <el-col :span="12">
        <el-input placeholder="请输入内容" v-model="name">
            <template slot="prepend">流程名称</template>
        </el-input>
      </el-col>
      <el-col :span="12">
         <el-button @click="saveFlow">保存</el-button>
      </el-col>
    </el-row>
    <workflow-chart
      :style="size"
      :transitions="transitions"
      :states="states"
      :stateSemantics="stateSemantics"
      :orientation="'horizontal'"
      @state-click="onLabelClicked('state',$event)"
      @transition-click="onLabelClicked('transition', $event)"
      @size-change="sizeChanged" />
    <el-dialog :visible.sync="dialogVisible" close-on-click-modal width="600px">
      <el-card>
        <span>{{state.label}}</span>
        <div v-if="state.id === 'start'">
          <el-button type="primary" size="mini" @click="addState('child')">添加下级</el-button>
        </div>
        <div v-else-if="state.id === 'end'">
          <el-button type="primary" size="mini" @click="addState('parent')">添加上级</el-button>
        </div>
        <div v-else>
          <el-button type="primary" size="mini" @click="addState('parent')">添加上级</el-button>
          <el-button type="primary" size="mini" @click="addState('same')">添加同级</el-button>
          <el-button type="primary" size="mini" @click="addState('child')">添加下级</el-button>
          <el-button type="danger" size="mini" @click="deleteState">删除</el-button>
          <el-divider></el-divider>
          <user @handleUser="setUser"></user>
        </div>
      </el-card>
    </el-dialog>
  </div>
</template>

<script>
import WorkflowChart from 'vue-workflow-chart';
import api from '../../api/flow'
import user from "../../components/user.vue"

export default {
  name: "App",
  components: {
    WorkflowChart,
    user,
  },
  data: () => ({
    dialogVisible: false,
    state:{},
    stateIndex:0,
    id:'',
    name:'',

    states: [],
    transitions: [],
    stateSemantics: [{
      "classname": "approve",
      "id":"start",
      },
      {
        "classname": "approve",
        "id":"end",
      }
    ],
    size: { width: '0px', height: '0px' },
  }),
  created() {
    this.id = this.$route.query.id
    if (this.id) {
      api.getFlow(this.id).then(({data}) => {
        this.states = data.items.v.states;
        this.transitions = data.items.v.transitions
        this.name = data.items.v.name
      });
    }else{
      this.states=[{id: "start", label: "开始"}, {id: "end", label: "结束"}]
      this.transitions = [{"id": "Dz2un1r","source": "start","target": "end"}]
    }
  },
  methods: {
    saveFlow(){
      var  flow = {"name":this.name,"states":this.states,"transitions":this.transitions,"time":this.getTime()}
      if (this.id){
        api.updateFlow(this.id,flow)
      }else{
        api.saveFlow(flow)
      }
    },
    getId(){
      return Math.random().toString(32).substr(2)
    },
    getTime(){
      var date = new Date();
      return date.getFullYear() + '-' + (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '-'
      + (date.getDate()<10?'0'+date.getDate():date.getDate()) + ' ' + date.getHours() + ':' + date.getMinutes() + ':' + date.getSeconds()
    },
    onLabelClicked(type, id) {
      this.dialogVisible = true
      this.state=this.getValue(this.states,id)
      this.stateIndex = this.states.indexOf(this.state)
    },
    setUser(user){
      this.state.label = user.name
      this.updateState()
    },
    updateState(){
      console.info(this.state)
      this.states[this.stateIndex]=this.state
      this.dialogVisible = false
    },
    addState(type){
      var st={"id":this.getId(),"label":"待设置"}
      this.states.push(st)
      var addTrans=[]

      if(type=="parent"){
          for (var i =0;i< this.transitions.length;i++){
            if(this.transitions[i].target==this.state.id){
               this.transitions[i].target=st.id
               var stt = {"id":this.getId(),"label":"","source":st.id,"target":this.state.id}
               addTrans.push(stt)
            }
          }
      }else if (type=="child"){
          for (var i =0;i< this.transitions.length;i++){
            if(this.transitions[i].source==this.state.id){
               this.transitions[i].source=st.id
               var stt = {"id":this.getId(),"label":"","source":this.state.id,"target":st.id}
               addTrans.push(stt)
            }
          }
      }else{
        for (var i =0;i< this.transitions.length;i++){
          if(this.transitions[i].source==this.state.id){
             var stt = {"id":this.getId(),"label":"","source":st.id,"target":this.transitions[i].target}
             addTrans.push(stt)
          }
          if(this.transitions[i].target==this.state.id){
             var stt = {"id":this.getId(),"label":"","source":this.transitions[i].source,"target":st.id}
             addTrans.push(stt)
          }
        }
      }

    this.transitions=this.transitions.concat(addTrans)
    this.dialogVisible = false
    },
    deleteState(){
      //delete from transitions
      var stateTrans=[]
      var tmpTarget=""
      var tmpSource=[]
      for (var i =0;i< this.transitions.length;i++){
        //将source与target连接
        if(this.transitions[i].source==this.state.id){
            tmpTarget = this.transitions[i].target
        }else if (this.transitions[i].target==this.state.id){
            tmpSource.push(this.transitions[i].source)
        }else{
          stateTrans.push(this.transitions[i])
        }
      }

      if(tmpTarget!=""&&tmpSource.length>0){
        for ( var s in tmpSource) {
          stateTrans.push({"id":this.getId(),"label":"","source":tmpSource[s],"target":tmpTarget})
        }
      }


      this.transitions = stateTrans

      //delete from states
      const index = this.states.indexOf(this.state);
      if (index > -1) { // only splice array when item is found
        this.states.splice(index, 1); // 2nd parameter means remove one item only
      }

      this.dialogVisible = false
    },
    getValue(list,id){
       for (var i = 0; i < list.length; i++) {
         if(list[i].id==id){
           console.log(list[i].label)
           return list[i];
         }
      }
    },
    sizeChanged(size) {
      this.size = {
        width: `${size.width}px`,
        height: `${size.height}px`,
      };
    },
  },
};
</script>

<style lang="scss">
@import '~vue-workflow-chart/dist/vue-workflow-chart.css';
$approve-color: #1eb2a4;
$delete-color: #d64b61;
$approve-color: #1eb2a4;
$delete-color: #d64b61;
.vue-workflow-chart-state {
&-approve {
   color: white;
   background: $approve-color;
 }
&-delete {
   color: white;
   background: $delete-color;
 }
}
.vue-workflow-chart-transition-arrow {
&-approve {
   fill: $approve-color;
 }
&-delete {
   fill: $delete-color;
 }
}
.vue-workflow-chart-transition-path {
&-approve {
   stroke: $approve-color;
 }
&-delete {
   stroke: $delete-color;
 }
}
</style>
<style lang="scss" scoped>
#app {
  display: flex;
  justify-content: center;
  padding-top: 100px;
}
</style>
