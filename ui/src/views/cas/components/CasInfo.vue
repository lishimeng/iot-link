<template>
  <div class="shape">
    <div slot="header">
      <h1>{{ info.label }}</h1><sub>{{ info.subLabel }}</sub>
    </div>
    <div v-if="info.pic" class="box-center">
      <pan-thumb :image="info.pic" :height="'100px'" :width="'100px'" :hoverable="false" />
    </div>
  </div>
</template>

<script>
import PanThumb from '@/components/PanThumb'

export default {
  name: 'CasInfo',
  components: { PanThumb },
  props: {
    info: {
      type: Object,
      default: () => {
        return {
          label: '',
          subLabel: '',
          pic: ''
        }
      }
    }
  },
  data: () => {
    return {
    }
  }
}
</script>

<style scoped>

  .shape{
    margin-top: 15px;
    margin-bottom: 15px;
  }
  .box-center {
    margin: 0 auto;
    display: table;
  }
</style>
