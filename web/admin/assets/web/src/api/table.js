import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/api/kv/flow/flow',
    method: 'get',
    params
  })
}
