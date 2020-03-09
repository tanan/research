<template>
  <div>
    <v-container>
      <v-row justify="center">
        <v-col md="9">
          <v-card outlined tile class="no-border">
            <v-card-title class="justify-center headline font-weight-bold">{{ article.title }}</v-card-title>
            <v-card-text>{{ article.description }}</v-card-text>
            <v-card-text class="body-1 font-weight-bold pl-12" v-html="article.content"></v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import { RepositoryFactory } from '@/repositories/RepositoryFactory'
import { documentToHtmlString } from '@contentful/rich-text-html-renderer';
const ArticleRepository = RepositoryFactory.get('articles')

export default {
  name: 'ArticleView',
  components: {

  },
  data () {
    return {
      article: this.fetchArticle(this.$route.params.id)
    }
  },
  methods: {
    async fetchArticle (id) {
      console.log(id)
      const res = await ArticleRepository.getArticle(id)
      if (res.status === 200) {
        console.log(res.data)
        let data = res.data
        let article = {}
        article.title = data.fields.title
        article.description = data.fields.description
        article.content = documentToHtmlString(data.fields.content)
        this.article = article
      }
    }
  }
}
</script>

<style lang="scss">
.no-border {
  border-color: white !important;
}
</style>