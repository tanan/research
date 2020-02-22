<template>
  <v-app>
    <v-app-bar app color="primary" dark>
      <v-toolbar-title>Research</v-toolbar-title>
      <v-spacer></v-spacer>
    </v-app-bar>
    <v-content>
      <NotificationContainer />
      <router-view></router-view>
    </v-content>
    <v-footer
      color="primary lighten-1"
      padless
    >
      <v-row
        justify="center"
        no-gutters
      >
        <!-- <v-btn
          v-for="link in links"
          :key="`${link.label}-header-link`"
          color="white"
          text
          rounded
          class="my-2"
          :to="link.url"
        >
          {{ link.label }}
        </v-btn> -->
        <v-col
          class="primary lighten-2 py-4 text-center white--text"
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
