const baseApp = {
  code: 0,
  message: ""
}

const apps = {
  'APP001': {
    id: 'APP001',
    description: 'I am a super Application 001'
  },
  'APP002': {
    id: 'APP002',
    description: 'I am a super Application 002'
  }
}

export default [
  // get app info
  {
    url: `/api/[A-Za-z0-9]/application/[A-Za-z0-9]`,
    type: 'get',
    response: config => {
      baseApp['item'] = apps['APP001']

      return baseApp
    }
  },

  // user logout
  {
    url: '/user/logout',
    type: 'post',
    response: _ => {
      return {
        code: 20000,
        data: 'success'
      }
    }
  }
]
