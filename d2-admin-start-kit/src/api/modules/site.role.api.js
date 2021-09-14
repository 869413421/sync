export default ({ request }) => ({
  /**
   * @description 获取列表
   * @param {BigInteger} page 登录携带的信息
   */
  SYS_ROLE_LIST(page) {
    // 接口请求
    return request({
      url: '/role?page=' + page,
      method: 'get'
    })
  },
  /**
   * @description 获取详情
   * @param {BigInteger} id ID
   */
  SYS_ROLE_INFO(id) {
    // 接口请求
    return request({
      url: '/role/' + id,
      method: 'get'
    })
  },
  /**
   * @description 更新
   * @param {BigInteger} id ID
   * @param {Object} data 数据
   */
  SYS_ROLE_UPDATE(id, data) {
    // 接口请求
    return request({
      url: '/role/' + id,
      method: 'put',
      data: data
    })
  },
  /**
  * @description 新增
  * @param {Object} data 数据
  */
  SYS_ROLE_STORE(data) {
    // 接口请求
    return request({
      url: '/role',
      method: 'post',
      data: data
    })
  },
  /**
   * @description 删除
   * @param {BigInteger} id ID
   */
  SYS_ROLE_DELETE(id) {
    // 接口请求
    return request({
      url: '/role/' + id,
      method: 'delete'
    })
  }
})
