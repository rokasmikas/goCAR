<template>
<div>
  <a-menu mode="horizontal">
    <a-menu-item>
      <router-link to="/">
        <a-icon type="home" />Home
      </router-link>
    </a-menu-item>

    <a-menu-item active key="file-add">
      <router-link to="/addlog">
        <a-icon type="file-add" />Add log
      </router-link>
    </a-menu-item>
  </a-menu>
</div>
</template>
<script>
export default {
  data() {
    return {

    }
  },
}
</script>

<script>
export default {
  data() {
    return {

    }
  }
}
</script>

<style scoped>

</style>
