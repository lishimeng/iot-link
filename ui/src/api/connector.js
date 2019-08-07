import request from '@/utils/request'

export function getConnector(id) {
  let c = request({
    url: '/api/1/connector/' + id,
    method: 'get'
  })
  if (c != null && c.code === 0) {
    return c.item
  } else {
    return null
  }
  return request({
    url: '/api/1/connector/' + id,
    method: 'get'
  })
}
