import { addCluster, getList } from '@/api/cluster'

const state = () => ({
    clusterlist: [],
  })

const mutations = {
  SET_Cluster_List: (state, data) => {
    state.clusterlist = data
  }
}

const actions = {
  // add cluster
  addCluster({ commit }, userInfo) {
    const { username, password } = userInfo
    return new Promise((resolve, reject) => {
      addCluster({ username: username.trim(), password: password }).then(response => {
        const { result } = response
        commit('SET_Cluster_List', result.data)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  // get cluster list
  getList({ commit }) {
    return new Promise((resolve, reject) => {
      getList().then(response => {
        const { result } = response
        commit('SET_Cluster_List', result.data)
        resolve(result)
      }).catch(error => {
        reject(error)
      })
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

