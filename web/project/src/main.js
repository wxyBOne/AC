/*
 * @Author: wxyBone 13638456960@163.com
 * @Date: 2024-12-04 11:12:53
 * @LastEditors: wxyBone 13638456960@163.com
 * @LastEditTime: 2024-12-17 11:49:53
 * @FilePath: \project\src\main.js
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
import { createApp } from "vue";
import router from "./router";
import store from "./store";
import Axios from 'axios'
import App from "./App.vue";

// 完整导入 Element Plus
import ElementPlus from "element-plus";
import "../node_modules/element-plus/dist/index.css";
// 黑夜模式
import "../node_modules/element-plus/theme-chalk/dark/css-vars.css";

// 导入 Element Plus 图标库
import * as ElementPlusIconsVue from "@element-plus/icons-vue";

import '@fortawesome/fontawesome-free/css/all.css'

// 创建 Vue 应用实例
const app = createApp(App);

// 注册 Element Plus 图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component);
}

// 检查 Local Storage 中的 token 和用户名
const token = localStorage.getItem('token');
const username = localStorage.getItem('username');
const privateKey = localStorage.getItem('privateKey');
if (token) {
    store.dispatch('setToken', token);
}
if (username) {
    store.dispatch('updateUserInfo', { username, privateKey }); // 假设您没有存储私钥
}

// 检查后端状态
const checkBackendStatus = async () => {
    try {
        const response = await Axios.get('http://localhost:8001/status'); // 检查后端状态
        if (response.data.resetUserInfo) {
            // 如果需要重置用户信息，清空 Local Storage
            localStorage.removeItem('token');
            localStorage.removeItem('username');
            localStorage.removeItem('privateKey');
            // 更新 Vuex 中的用户信息
            store.dispatch('updateUserInfo', {
                username: '',
                privateKey: ''
            });
        }
    } catch (error) {
        console.error('无法检查后端状态', error);
    }
};

// 调用状态检查
checkBackendStatus();

// 使用 Element Plus
app.config.globalProperties.$axios = Axios
app.use(ElementPlus).use(router).use(store).mount("#app");
