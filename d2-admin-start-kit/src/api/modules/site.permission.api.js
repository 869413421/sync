export default ({ request }) => ({
  /**
   * @description 获取列表
   * @param {BigInteger} page 登录携带的信息
   */
  SYS_PERMISSION_LIST(page) {
    // 接口请求
    return request({
      url: '/permission?page=' + page,
      method: 'get'
    })
  },
  /**
   * @description 获取详情
   * @param {BigInteger} id ID
   */
  SYS_PERMISSION_INFO(id) {
    // 接口请求
    return request({
      url: '/permission/' + id,
      method: 'get'
    })
  },
  /**
   * @description 更新
   * @param {BigInteger} id ID
   * @param {Object} data 数据
   */
  SYS_PERMISSION_UPDATE(id, data) {
    // 接口请求
    return request({
      url: '/permission/' + id,
      method: 'put',
      data: data
    })
  },
  /**
  * @description 新增
  * @param {Object} data 数据
  */
  SYS_PERMISSION_STORE(data) {
    // 接口请求
    return request({
      url: '/permission',
      method: 'post',
      data: data
    })
  },
  /**
  * @description 获取权限树
  */
  SYS_PERMISSION_TREE() {
    // 接口请求
    return request({
      url: '/permission/tree',
      method: 'get',
    })
  },
  /**
   * @description 删除
   * @param {BigInteger} id ID
   */
  SYS_PERMISSION_DELETE(id) {
    // 接口请求
    return request({
      url: '/permission/' + id,
      method: 'delete'
    })
  }
})
