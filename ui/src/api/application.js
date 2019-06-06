import request from '@/utils/request'

export function getApplication(id) {
  return request({
    url: '/api/1/application/' + id,
    method: 'get'
  })
}
