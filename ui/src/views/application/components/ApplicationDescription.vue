<template>
  <div>
    <span>TODO show application description-{{ id }}</span>
  </div>
</template>

<script>
export default {
  name: 'ApplicationDescription',
  props: {
    id: {
      type: String,
      default: () => {
        return ''
      }
    }
  }
}
</script>

<style scoped>

</style>
