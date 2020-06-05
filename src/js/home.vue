<template>
<div>
  <a-row :gutter="24">
  <div v-for="(item, key) in carlogs">
    <a-col :span="8">
    <card v-bind:make="item.make"
          v-bind:model="item.model"
          v-bind:reg="item.reg"
          v-bind:year="item.year"
          v-bind:id="item.id"
    />
  </a-col>
  </div>
</a-row>
</div>
</template>
<script>
import axios from "axios"
import card from './components/card.vue'

export default {
  components: {
    card,
  },
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
