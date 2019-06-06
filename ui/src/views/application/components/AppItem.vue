<template>
  <div>
    <el-table ref="dragTable" :data="list" row-key="id" border fit highlight-current-row style="width: 100%">
      <el-table-column align="center" label="ID" prop="AppId" />
      <el-table-column align="center" label="Description">
        <template slot-scope="{row}">
          <router-link :to="'/application/'+row.AppId" class="link-type">
            <span>{{ row.AppDescription }}</span>
          </router-link>
        </template>
      </el-table-column>
      <el-table-column align="center" label="Codec" prop="CodecType" />
      <el-table-column align="center" label="Connector" prop="Connector" />
      <el-table-column align="center" label="CreateTime" prop="CreateTime" />
    </el-table>
  </div>
</template>

<script>
export default {
  name: 'AppItem',
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
