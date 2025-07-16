/*
 * @Author: wxyBone 13638456960@163.com
 * @Date: 2024-12-04 11:12:53
 * @LastEditors: wxyBone 13638456960@163.com
 * @LastEditTime: 2024-12-09 14:17:52
 * @FilePath: \project\src\store\index.js
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
import { createStore } from 'vuex';

export default createStore({
  // 存储应用状态
  state: {
    userInfo: {
      username: '',
      privateKey: ''
    },
    token: ''
  },
  // 修改状态
  mutations: {
    setUserInfo(state, userInfo) {
      state.userInfo = userInfo;
    },
    setToken(state, token) {
      state.token = token;
    }
  },
  // 处理异步等复杂逻辑
  actions: {
    updateUserInfo({ commit }, userInfo) {
      commit('setUserInfo', userInfo);
    },
    setToken({ commit }, token) {
      commit('setToken', token);
    }
  },
  // 获取状态数据
  getters: {
    getUserInfo: (state) => state.userInfo,
    getToken: (state) => state.token
  }
});