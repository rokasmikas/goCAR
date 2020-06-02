<template>
<div>

  <!-- TODO: add bootstrap styles -->
  <div v-show="showInput">
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

    <button @click="update">Update</button>
  </div>

  <h4>{{carLogData.name}}</h4>
  <h4>{{carLogData.make}}</h4>
  <h4>{{carLogData.model}}</h4>
  <h4>{{carLogData.reg}}</h4>
  <h4>{{carLogData.year}}</h4>

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
      showInput: false,
    }
  },
  methods: {
    setData(data) {
      this.carLogData = data
    },
    editLog() {
      this.showInput = !this.showInput
    },
    update() {
      let o = this.carLogData
      let that = this
      console.log(this.$route.params.id)
      axios.put('/api/carlog/update/'+this.$route.params.id, {
          name: o.name,
          make: o.make,
          model: o.model,
          reg: o.reg,
          year: o.year,
        })
        .then(function(res) {
            console.log(res)
            that.showInput = false
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
