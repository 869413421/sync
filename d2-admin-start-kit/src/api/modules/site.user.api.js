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
  }
})
