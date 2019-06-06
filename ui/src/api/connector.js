import request from '@/utils/request'

export function getConnector(id) {
  return request({
    url: '/api/1/connector/' + id,
    method: 'get'
  })
}
