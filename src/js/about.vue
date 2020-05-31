<template>
<div>
  <div>MyApp</div>
  <button v-on:click="showMessage">This is test</button>
  <div>Test</div>
</div>
</template>
<script>
import axios from "axios"

export default {
  data() {
    return {
      message: ""
    }
  },
  methods: {
    showMessage() {
      axios.get("/api/v1/hello").then(res => {
        console.log(res)
        appData.message = res.data.message
      });
    }
  }
};
</script>
<style>
</style>
