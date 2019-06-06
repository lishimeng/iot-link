<template>
  <div>
    <el-row>
      <app-item :list="sample" />
    </el-row>
  </div>
</template>

<script>
import AppItem from './components/AppItem'
export default {
  name: 'Connectors',
  components: { AppItem },
  data: () => {
    return {
      sample: [
        { AppId: 1, AppDescription: 'Radar', CodecType: 'javascript', Connector: 'CONN13766462136', CreateTime: '14823473274' },
        { AppId: 2, AppDescription: 'Humidity', CodecType: 'Tlv', Connector: 'CONN13764624223', CreateTime: '14823432223' }
      ]
    }
  }
}
</script>

<style scoped>

</style>
