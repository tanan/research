<template>
  <v-app>
    <v-app-bar app color="#fff" dense flat>
      <v-toolbar-title>Research</v-toolbar-title>
    </v-app-bar>
    <v-content>
      <NotificationContainer />
      <router-view></router-view>
    </v-content>
    <v-footer
      color="#fff"
      padless
    >
      <v-row
        justify="center"
        no-gutters
      >
        <v-col
          class="grey darken-4 py-4 text-center white--text"
          cols="12"
        >
          {{ new Date().getFullYear() }} — <strong>RESEARCH</strong>
        </v-col>
      </v-row>
    </v-footer>
  </v-app>
</template>

<script>
import NotificationContainer from '@/components/NotificationContainer.vue';
export default {
  name: 'HelloWorld',
  components: {
    NotificationContainer
  },
  data () {
    return {
      links: [
        {
          label: 'Login',
          url: '/login',
        }
      ]
    }
  }
}
</script>

<style lang="scss">
.v-application {
  font-family: "M PLUS Rounded 1c" !important;
}
</style>