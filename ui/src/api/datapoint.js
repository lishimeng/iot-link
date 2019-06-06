import request from '@/utils/request'

export function getDataPoint(appId) {
  return request({
    url: '/api/1/application/' + appId + '/dp',
    method: 'get'
  })
}
