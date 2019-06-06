import request from '@/utils/request'

export function getCodecJs(appId) {
  return request({
    url: '/api/1/application/' + appId + '/codec/js',
    method: 'get'
  })
}
