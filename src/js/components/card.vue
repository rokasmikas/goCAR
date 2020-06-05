<template>
<div style="marginBottom: 20px">
  <a-card :loading="loading" v-bind:title="getCarName">
    {{reg}}
    {{year}}
    <router-link :to="{ name: 'showlog', params: {id: id} }">Show more</router-link>
  </a-card>
</div>
</template>
<script>
export default {
  props: {
    make: String,
    model: String,
    reg: String,
    year: String,
    id: String,
  },
  data() {
    return {
      loading: true,
    };
  },
  computed: {
    getCarName: function() {
      return this.make + " " + this.model
    }
  },
  mounted() {
    // fake some delay for better UX
    let delay = 500
    setTimeout(() => {
      this.loading = false
    }, delay)
  },
};
</script>
