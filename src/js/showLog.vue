<template>
<div>

  <!-- TODO: add bootstrap styles -->
  <a-col :span="6">
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


    <button @click="editLog">Edit</button> </br></br>
    <button @click="showOrderAddition = !showOrderAddition">Add Order</button> </br>
  </a-col>


<a-col :span="6">
  <div v-if="showOrderAddition">
    <label>Name</label>
    <input type="text" v-model="order.name" lazy>
    <button @click="addOrder()">Add the order</button>
  </div>
</a-col>
  <a-col :span="12">
  <li v-for="(item, key) in carLogData.orders">
    {{item}}
  </li>
  </a-col>
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
        orders: [],
      },
      order: {
        name: null,
      },
      noneditable: true,
      showOrderAddition: true,
    }
  },
  methods: {
    setData(data) {
      this.carLogData = data.data
      this.carLogData.orders = data.associated
    },
    editLog() {
      this.noneditable = !this.noneditable
    },
    addOrder() {
      let o = this.order
      axios.post('/api/orders/' + this.$route.params.id + '/create', {
          name: o.name,
        })
        .then(function(res) {
          console.log(res)
        })
        .catch(function(err) {
          console.log(err)
        })
    },
    update() {
      let o = this.carLogData
      let that = this
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
        vm.setData(res.data)
      })
    })
  }

}
</script>

<style>

</style>
