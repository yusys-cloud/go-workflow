import request from '@/utils/request'
export default {
  getFlow(id) {
    return request({
      url: '/api/kv/flow/flow/'+id,
      method: 'get'
    })
  },
  saveFlow(data){
    return request({
      url: '/api/kv/flow/flow',
      method: 'post',
      data
    })
  },
  updateFlow(id,data){
    return request({
      url: '/api/kv/flow/flow/'+id,
      method: 'put',
      data
    })
  },
  
  getUsers() {
    return request({
      url: '/api/kv/user/user',
      method: 'get'
    })
  }
}
