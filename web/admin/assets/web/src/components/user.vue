<template>
  <div>
  <el-card style="width: 100%;">
    <span>用户选择</span>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column prop="name" label="名称" width="180" sortable></el-table-column>
      <el-table-column prop="des" label="描述" width="180" sortable></el-table-column>
      <el-table-column fixed="right" label="操作" width="100">
            <template slot-scope="scope">
              <el-button @click="handleUser(scope.row)" type="text" size="small">选择</el-button>
            </template>
      </el-table-column>
    </el-table>
  </el-card>

  </div>
</template>
<script>
  import api from '../api/flow'
  export default {
    data() {
      return {
        tableData:[{"name":"张三","des":""},{"name":"joe","des":"asd"}]
      }
    },
    created() {
      api.getUsers().then(({data}) => {
        console.info(data.items[0].v)
        this.tableData = data.items[0].v
      });
    },
    methods: {
      handleUser(user) {
        this.$emit('handleUser',user)
      }
    }
  }
</script>

<style>
</style>
