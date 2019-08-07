<template>
  <div class="body">
    <div class="form-row">
      <el-input v-model="broker" placeholder="tcp://localhost:1883">
        <template slot="prepend">broker</template>
      </el-input>
    </div>
    <div class="form-row">
      <el-input v-model="upLink" placeholder="application/+/device/+/uplink">
        <template slot="prepend">UpLink</template>
      </el-input>
    </div>
    <div class="form-row">
      <el-input v-model="downLink" placeholder="application/+/device/+/downlink">
        <template slot="prepend">DownLink</template>
      </el-input>
    </div>
    <div class="form-row">
      <el-input v-model="tpl" placeholder="application/{ApplicationID}/device/{DeviceID}/downlink">
        <template slot="prepend">Template</template>
      </el-input>
    </div>
    <div class="form-row">
      <el-button type="primary">Save</el-button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'MqttConnectorInfo',
  data() {
    return {
      broker: '',
      upLink: '',
      downLink: '',
      tpl: ''
    }
  }
}
</script>

<style scoped>

  .body {
    padding: 32px;
  }
  .form-row{
    margin-top: 15px;
    width: 600px;
  }

</style>
