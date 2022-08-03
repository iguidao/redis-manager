import request from '@/utils/request'

export function addCluster(data) {
  return request({
    url: '/redis-web-manager/v1/cluster/add',
    method: 'post',
    data
  })
}

export function getList(query) {
  return request({
    url: '/redis-web-manager/v1/cluster/list',
    method: 'get',
    params: { query }
  })
}

// export function logout() {
//   return request({
//     url: '/vue-admin-template/user/logout',
//     method: 'post'
//   })
// }
