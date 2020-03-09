<template>
  <div>
    <v-container>
      <v-row>
        <v-col md="8">
          <div class="new-article">
            <h2>新着記事</h2>
          </div>
          <v-row>
            <v-col class="pt-6 pr-8" md="5" v-for="(article, index) in articles" :key="index">
              <v-card tile min-height="340px" :href="article.landingUrl">
                <v-img class="white--text align-end" height="180px" :src="article.thumbnailUrl">
                  <v-card-title>{{ article.title }}</v-card-title>
                </v-img>
                <v-card-text>{{ article.description }}</v-card-text>
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
      articles: this.fetchArticles()
    }
  },
  methods: {
    getMessage () {
      return ""
    },
    async fetchArticles () {
      const res = await ArticleRepository.get()
      if (res.status === 200) {
        console.log(res.data)
        let data = res.data
        let items = data.items
        let asset = data.includes.Asset
        let articles = items.map(l => {
          let r = {}
          r.title = l.fields.title
          if (l.fields.thumbnail) {
            r.thumbnailId = l.fields.thumbnail.sys.id
            r.thumbnailUrl = "https:" + asset.find(e => e.sys.id === r.thumbnailId).fields.file.url
          }
          r.description = l.fields.description
          r.landingUrl = "/article/" + l.sys.id
          return r
        })
        console.log(articles)
        this.articles = articles
      }
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
