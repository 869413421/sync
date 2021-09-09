export default ({ request }) => ({
  /**
   * @description 获取规则列表
   * @param {BigInteger} page 登录携带的信息
   */
  SYS_CASBIN_LIST(page, type) {
    // 接口请求
    return request({
      url: '/casbin?page=' + page + "&ptype=" + type,
      method: 'get'
    })
  },
  /**
   * @description 获取规则详情
   * @param {BigInteger} id 规则ID
   */
  SYS_CASBIN_INFO(id) {
    // 接口请求
    return request({
      url: '/casbin/' + id,
      method: 'get'
    })
  },
  /**
   * @description 更新规则
   * @param {BigInteger} id 规则ID
   * @param {Object} data 规则信息
   */
  SYS_CASBIN_UPDATE(id, data) {
    // 接口请求
    return request({
      url: '/casbin/' + id,
      method: 'put',
      data: data
    })
  },
  /**
  * @description 新增规则
  * @param {Object} data 用户规则
  */
  SYS_CASBIN_STORE(data) {
    // 接口请求
    return request({
      url: '/casbin',
      method: 'post',
      data: data
    })
  },
  /**
  * @description 获取权限树
  */
  SYS_CASBIN_TREE() {
    // 接口请求
    return request({
      url: '/casbin/tree',
      method: 'get',
    })
  },
  /**
   * @description 删除规则
   * @param {BigInteger} id 规则ID
   */
  SYS_CASBIN_DELETE(id) {
    // 接口请求
    return request({
      url: '/casbin/' + id,
      method: 'delete'
    })
  }
})
