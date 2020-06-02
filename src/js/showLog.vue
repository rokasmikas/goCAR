<template>
<div>

  <!-- TODO: add bootstrap styles -->
  <label>Name</label>
  <input type="text" v-model="carLogData.name" :disabled="noneditable" lazy>
  </br>
  <label>Make</label>
  <input type="text" v-model="carLogData.make" :disabled="noneditable" lazy>
  </br>
  <label>Model</label>
  <input type="text" v-model="carLogData.model" :disabled="noneditable" lazy>
  </br>
  <label>Reg</label>
  <input type="text" v-model="carLogData.reg" :disabled="noneditable" lazy>
  </br>
  <label>Year</label>
  <input type="text" v-model="carLogData.year" :disabled="noneditable" lazy>
  </br>
  <button @click="update" :disabled="noneditable">Update</button>

  <button @click="editLog">Edit</button>
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
      noneditable: true,
    }
  },
  methods: {
    setData(data) {
      this.carLogData = data
    },
    editLog() {
      this.noneditable = !this.noneditable
    },
    update() {
      let o = this.carLogData
      let that = this
      console.log(this.$route.params.id)
      axios.put('/api/carlog/update/' + this.$route.params.id, {
          name: o.name,
          make: o.make,
          model: o.model,
          reg: o.reg,
          year: o.year,
        })
        .then(function(res) {
          console.log(res)
          that.edit = false
        })
        .catch(function(err) {
          console.log(err)
        })
    }
  },
  beforeRouteEnter(to, from, next) {
    axios.get("/api/carlog/get/" + to.params.id).then(res => {
      next(vm => {
        vm.setData(res.data.data)
      })
    })
  }

}
</script>

<style>

</style>
