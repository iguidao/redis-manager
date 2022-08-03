import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/vue-admin-template/table/list',
    method: 'get',
    params
  })
}

export function transactionList(query) {
  return request({
    url: '/redis-manager/v1/redis/list',
    method: 'get',
    params: { query }
  })
}