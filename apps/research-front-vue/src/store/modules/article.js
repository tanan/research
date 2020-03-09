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
  }
}
