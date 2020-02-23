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
                <v-img class="white--text align-end" height="180px" :src="article.imgurl">
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
import { RepositoryFactory } from '@/repositories/RepositoryFactory'
const ArticleRepository = RepositoryFactory.get('articles')

export default {
  name: 'ResearchTop',
  components: {
    
  },
  computed: {
  },
  data () {
    return {
      // articles: [
      //   {
      //     title: "タイトル1",
      //     imgurl: "https://cdn.vuetifyjs.com/images/cards/docks.jpg",
      //     content: "コンテンツが流れています"
      //   },
      //   {
      //     title: "タイトル2",
      //     imgurl: "https://cdn.vuetifyjs.com/images/cards/docks.jpg",
      //     content: "コンテンツが流れています"
      //   }
      // ]
      articles: this.fetchArticles()
    }
  },
  methods: {
    getMessage () {
      return ""
    },
    async fetchArticles () {
      const { data } = await ArticleRepository.get()
      return data.articles
    }
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
