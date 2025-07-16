/*
 * @Author: wxyBone 13638456960@163.com
 * @Date: 2024-12-04 11:10:42
 * @LastEditors: wxyBone 13638456960@163.com
 * @LastEditTime: 2024-12-19 17:20:11
 * @FilePath: \project\truffle-config.js
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
module.exports = {
  networks: {
    development: {
      host: "127.0.0.1",
      port: 7545,  // 改为 Ganache 默认端口
      network_id: "*"
    }
  },
  compilers: {
    solc: {
      version: "0.8.21",  // 使用当前版本
      settings: {
        optimizer: {
          enabled: true,
          runs: 200
        }
      }
    }
  }
};