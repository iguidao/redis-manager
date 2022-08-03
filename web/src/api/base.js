import request from '@/utils/request'

export function getOk(params) {
  return request({
    url: '/base/v1/health',
    method: 'get',
    params
  })
}
