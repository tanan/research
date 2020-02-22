<template>
  <div>
    <v-container>
      <v-row>
        <v-col md="8">
          <div class="new-article">
            <h2>新着記事</h2>
          </div>
          <v-row>
            <v-col class="pt-6 pr-8" md="6" v-for="(article, index) in articles" :key="index">
              <v-card min-height="340px">
                <v-img class="white--text align-end" height="180px" src="https://cdn.vuetifyjs.com/images/cards/docks.jpg">
                  <v-card-title>{{ article.title }}</v-card-title>
                </v-img>
                <v-card-text>{{ article.content }}</v-card-text>
              </v-card>
            </v-col>
          </v-row>
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
      articles: [
        {
          title: "タイトル1",
          content: "コンテンツが流れています"
        },
        {
          title: "タイトル2",
          content: "コンテンツが流れています"
        },
        {
          title: "タイトル1",
          content: "コンテンツが流れています"
        },
        {
          title: "タイトル2",
          content: "コンテンツが流れています"
        }
      ]
    }
  },
  methods: {
    getMessage () {
      return ""
    },
  }
}
</script>

<style lang="scss">
.new-article {
  h2 {
    font-size: 18px;
  }
}
</style>
