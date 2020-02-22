<template>
  <div>
    <v-container>
      <v-row>
        <v-col v-for="i of articleCount" :key="i">
          <div :id="`form-${i}`">
            <p>コンテンツを書く</p>
          </div>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import { mapState } from 'vuex';

export default {
  name: 'ResearchTop',
  components: {
    
  },
  computed: {
    ...mapState(['article']),
  },
  data () {
    return {
      articleCount: 2
    }
  },
  methods: {
    getMessage () {
      return ""
    },
  }
}
</script>

<style scoped>
</style>
