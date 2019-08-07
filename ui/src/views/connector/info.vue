<template>
  <div>
    <div>connector info<span>{{ id }}</span></div>
    <mqtt-connector-info />
  </div>
</template>

<script>
import MqttConnectorInfo from './components/MqttConnectorInfo'
import { getConnector } from '@/api/application'
export default {
  name: 'Connector',
  components: { MqttConnectorInfo },
  data: () => {
    return {
      id: '',
      connector: {}
    }
  },
  created() {
    const id = this.$route.params && this.$route.params.id
    this.id = id
    this.setPageTitle(id)
    this.getConnectorInfo(id)
  },
  methods: {
    setPageTitle: function(id) {
      const t = 'connector'
      document.title = `${t} - ${id}`
    },
    getConnectorInfo: function(id) {
        let c = getConnector(id)
        if (c != null) {
            c.props
        }
        this.connector = {

        }
    }
  }
}
</script>

<style scoped>

</style>
