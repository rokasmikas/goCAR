<template>
<div>
  <form @submit="postMessage">

    <!-- TODO: add bootstrap styles -->
    <label>Name</label>
    <input type="text" v-model="carLogData.name" lazy>
    </br>
    <label>Make</label>
    <input type="text" v-model="carLogData.make" lazy>
    </br>
    <label>Model</label>
    <input type="text" v-model="carLogData.model" lazy>
    </br>
    <label>Reg</label>
    <input type="text" v-model="carLogData.reg" lazy>
    </br>
    <label>Year</label>
    <input type="text" v-model="carLogData.year" lazy>
    </br>

    <input type="submit" value="Submit">
  </form>
  {{errors}}
</div>
</template>
<script>
import axios from "axios"

export default {
  data() {
    return {
      carLogData: {
        name: null,
        make: null,
        model: null,
        reg: null,
        year: null,
      },
      errors: [],
    }
  },
  methods: {
    checkForm: function() {
      if (this.carLogData.name && this.carLogData.make && this.carLogData.model && this.carLogData.reg && this.carLogData.year) {
        return true
      }

      this.errors = []

      for (let x in this.carLogData) {
        if (!this.carLogData[x])
          this.errors.push(x.charAt(0).toUpperCase() + x.slice(1) + ' is missing'); // TODO: Convert "x to upercase" function to filter
      }
      return false
    },
    postMessage: function() {
      if (this.checkForm()) {
        let o = this.carLogData
        let that = this

        axios.post('/api/carlog/create', {
            name: o.name,
            make: o.make,
            model: o.model,
            reg: o.reg,
            year: o.year,
          })
          .then(function(res) {
            if(res.status == 201) {
              that.$router.push({ name: 'showlog', params: {id: res.data.data} })
            }
          })
          .catch(function(err) {
            console.log(err)
          })
      }
    }
  }
}
</script>

<style>

</style>
