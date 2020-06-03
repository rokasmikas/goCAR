<template>
<div>
  <table>
    <tr>
      <th>Name</th>
      <th>Brand (Make + model)</th>
      <th>Reg</th>
      <th>Year</th>
      <th>Delete</th>
    </tr>
    <tr v-for="(item, key) in carlogs">
      <th>
        <router-link :to="{ name: 'showlog', params: {id: item.id} }">{{item.name}}</router-link>
      </th>
      <th> {{item.make}} {{item.model}} </th>
      <th> {{item.reg}} </th>
      <th> {{item.model}} </th>
      <th> <button @click="deleteCarlog(item.id, key)">X</button></th>
    </tr>
  </table>
</div>
</template>
<script>
import axios from "axios"

export default {
  data() {
    return {
      carlogs: null,
    }
  },
  methods: {
    setData(data) {
      this.carlogs = data
    },
    deleteCarlog(id, key) {
      axios.delete('/api/carlog/delete/' + id)
        .then(res => {
          this.$delete(this.carlogs, key)
        })
        .catch(err => {
          console.log(err)
        })
    }
  },
  beforeRouteEnter(to, from, next) {
    axios.get("/api/carlog/getAll").then(res => {
      next(vm => vm.setData(res.data.data))
    })
  }
}
</script>
