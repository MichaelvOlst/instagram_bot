
const state = () => ({
  user: null,
  error: null,
})

// getters
const getters = {
  getUser: (state) => {
    return state.user
  },
  loggedIn: (state) => {
    return !!state.user 
  }
}

// actions
const actions = {
  checkLogin ({ state, commit }) {
    fetch("/auth/check", {
      method: 'POST', // *GET, POST, PUT, DELETE, etc.
    })
    .then(response => response.json())
    .then(data => {
      commit('setUser', data.user)
    })
    .catch(error => {
      commit('setError', error)
    });
  }
}

// mutations
const mutations = {

  setUser (state, { user }) {
    state.user = user
  },

  setError (state, { error }) {
    state.error = error
  },

}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}