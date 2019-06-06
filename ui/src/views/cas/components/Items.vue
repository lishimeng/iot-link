<template>
    <div>
      <el-table ref="dragTable" v-loading="listLoading" :data="list" row-key="id" border fit highlight-current-row style="width: 100%">
        <el-table-column align="center" label="ID" prop="id"></el-table-column>

        <el-table-column align="center" label="中文名" prop="cn"></el-table-column>

        <el-table-column align="center" label="英文名" prop="en"></el-table-column>

        <el-table-column label="CAS号" prop="cas"></el-table-column>

        <el-table-column align="center" label="分子式" prop="formula"></el-table-column>
      </el-table>
    </div>
</template>

<script>
export default {
  name: "Items",
  props: {
    list: {
      type: Array,
      default: () => {
        return null
      }
    }
  }
}
</script>

<style scoped>

</style>
