<template>
  <div>
    <div>MyApp</div>
    <button v-on:click="showMessage">This is test</button>
    <router-link to="/addlog">Add log</router-link>
    <div>{{message}}</div>
  </div>
</template>
<script>
import axios from "axios";
const appData = {
  message: ""
};
export default {
  data() {
    return appData;
  },
  methods: {
    showMessage: showMessage
  }
};
function showMessage() {
  axios.get("/api").then(res => {
    console.log(res);
    appData.message = res.data.message;
  });
}
</script>
<style>
</style>
