<template>
  <div>connector info<span>{{ id }}</span></div>
</template>

<script>
export default {
  name: 'Connector',
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
