import Vue from 'vue'
import Vuex from 'vuex'
import * as article from '@/store/modules/article.js';
import * as notification from '@/store/modules/notification.js';

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    article,
    notification
  },
  state: {
  }
})