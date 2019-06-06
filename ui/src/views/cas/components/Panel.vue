<template>
  <el-card shadow="always" class="border">
    <div slot="header" class="clearfix">
      <span>{{ title }}</span>
    </div>
    <div>
      <div v-for="(item, index) in items" :key="index">
        <el-row :gutter="10" class="row">
          <el-col :xs="24" :span="10">
            {{ item.label }}:
          </el-col>
          <el-col :xs="24" :span="12">
            {{ item.value }}
          </el-col>
        </el-row>
      </div>
    </div>
  </el-card>
</template>

<script>
export default {
  name: 'Panel',
  props: {
    title: {
      type: String,
      default: () => {
        return ''
      }
    },
    items: {
      type: Array,
      default: () => {
        return []
      }
    },
    showShadow: {
      type: String,
      default: () => {
        return 'always'
      }
    }
  }
}
</script>

<style scoped>
  .border {
    border: 0;
  }
  .row {
    margin-top: 15px;
  }
</style>
