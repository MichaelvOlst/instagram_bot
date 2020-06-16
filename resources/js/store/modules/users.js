
const state = () => ({
    users: [],
    error: null
  })
  
  const getters = {
    getUsers: (state) => {
      return state.users
    },
    
  }
  
  const actions = {
    getUsers ({ commit }) {
      fetch("/api/users")
      .then(response => response.json())
      .then(res => {     
        commit('setUsers', res.data)
      })
      .catch(error => {
        commit('setError', error)
      });
    },

    addUser ({ commit }, data) {
      fetch("/api/users", {
        method: 'POST',
        body: JSON.stringify(data),
        headers: {
          'Content-Type': 'application/json'
        },
      })
      .then(response => response.json())
      .then(res => {
        commit('setUsers', res.data)
      })
      .catch(error => {
        commit('setError', error)
      });
    },

  }
  
  const mutations = {
  
    setUsers (state, payload) {
      state.users = payload
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