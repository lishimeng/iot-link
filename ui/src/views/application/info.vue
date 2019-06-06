<template>
  <div>
    <application-description :id="id" />
  </div>
</template>

<script>
import ApplicationDescription from './components/ApplicationDescription'
export default {
  name: 'ApplicationInfo',
  components: { ApplicationDescription },
  data: () => {
    return {
      id: ''
    }
  },
  created() {
    const id = this.$route.params && this.$route.params.id
    this.id = id
    this.setPageTitle(id)
  },
  methods: {
    setPageTitle(id) {
      const t = 'Application'
      document.title = `${t} - ${id}`
    }
  }
}
</script>

<style scoped>

</style>
