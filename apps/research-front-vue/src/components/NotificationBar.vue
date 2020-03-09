<template>
  <v-alert
    :value="true"
    type="error"
    class="notification-bar"
    :class="notificationTypeClass"
    outlined
  >
    <p>{{ notification.message }}</p>
  </v-alert>
</template>

<script>
import { mapActions } from 'vuex'
export default {
  props: {
    notification: {
      type: Object,
      required: true
    }
  },
  mounted() {
    this.timeout = setTimeout(() => this.remove(this.notification), 5000)
  },
  computed: {
    notificationTypeClass() {
      return `-type-${this.notification.type}`
    }
  },
  methods: mapActions('notification', ['remove'])
}
</script>

<style lang="scss" scoped>
.notification-bar {
  margin: 1em 0 1em;
  font-size: 14px;
  p {
    margin: auto;
  }
}
</style>