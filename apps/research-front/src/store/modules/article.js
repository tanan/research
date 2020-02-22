import ArticleService from '@/services/ArticleService.js'
export const namespaced = true

export const state = {
  deleteGlobalIds: []
}

export const mutations = {
  SET_DELETE_GLOBALIDS(state, globalIds) {
    state.deleteGlobalIds = globalIds
  }
}

export const actions = {
  addArticle({ commit }, article) {
    commit('ADD_ARTICLE', article)
  },
  async post({ dispatch }, request) {
    try {
      let r = await ArticleService.postRequest(request)
      console.log(r.data)
    } catch (error) {
      let m = error.message
        if (error.response) {
          m = error.response.data.message
        }
        const notification = {
          type: 'error',
          message: m
        }
        dispatch('notification/add', notification, { root: true})
    }
  }
}
