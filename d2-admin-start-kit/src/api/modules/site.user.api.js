export default ({ request }) => ({
  /**
   * @description 获取用户列表
   * @param {BigInteger} page 登录携带的信息
   */
  SYS_USER_LIST(page) {
    // 接口请求
    return request({
      url: '/user?page=' + page,
      method: 'get'
    })
  },
  /**
   * @description 获取用户详情
   * @param {BigInteger} id 用户ID
   */
  SYS_USER_INFO(id) {
    // 接口请求
    return request({
      url: '/user/' + id,
      method: 'get'
    })
  },
  /**
   * @description 更新用户
   * @param {BigInteger} id 用户ID
   * @param {Object} data 用户信息
   */
  SYS_USER_UPDATE(id, data) {
    // 接口请求
    return request({
      url: '/user/' + id,
      method: 'put',
      data: data
    })
  },
  /**
  * @description 新增用户
  * @param {Object} data 用户信息
  */
  SYS_USER_STORE(data) {
    // 接口请求
    return request({
      url: '/user',
      method: 'post',
      data: data
    })
  }
})
