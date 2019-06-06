<template>
  <div style="display:inline-block;">
    <label class="radio-label" style="padding-left:0;">CAS: </label>
    <el-input v-model="name" placeholder="输入要查询的CAS号码" style="width:345px;" prefix-icon="el-icon-document" />
    <el-button type="primary" icon="document" >查询</el-button>
  </div>
</template>

<script>
export default {
  props: {
    value: {
      type: String,
      default: ''
    }
  },
  computed: {
    name: {
      get() {
        return this.value
      },
      set(val) {
        this.$emit('input', val)
      }
    }
  }
}
</script>
