
const state = () => ({
  user: null,
})

const getters = {
  getUser: (state) => {
    return state.user
  },
  
}

const actions = {
  checkLogin ({ commit }) {
    fetch("/auth/check", {
      method: 'POST',
    })
    .then(response => response.json())
    .then(res => {     
      commit('setUser', res.data)
    })
    .catch(error => {
      commit('setError', error)
    });
  },

  login({commit}, data) {

    fetch("/auth/login", {
      method: 'POST',
      body: JSON.stringify(data),
      headers: {
        'Content-Type': 'application/json'
      },
    })
    .then(response => response.json())
    .then(res => {
      commit('setUser', res.data)
    })
    .catch(error => {
      commit('setError', error)
    });
  },

  logout({commit}) {

    fetch("/auth/logout", {
      method: 'POST',
    })
    .then(response => response.json())
    .then(() => {
      commit('setUser', null)
    })
    .catch(error => {
      commit('setUser', null)
      commit('setError', error)
    });
  }

}

const mutations = {

  setUser (state, payload) {
    state.user = payload
  },

  setError (state, payload) {
    state.error = payload
  },

}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}