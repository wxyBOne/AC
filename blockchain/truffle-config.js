require('dotenv').config();
const HDWalletProvider = require('@truffle/hdwallet-provider');

module.exports = {
  // 合约编译输出目录配置
  contracts_build_directory: "../web/project/src/contracts",

  networks: {
    development: {
      host: "127.0.0.1",
      port: 7545,
      network_id: "*"
    }
  },

  compilers: {
    solc: {
      version: "0.8.18",
      settings: {
        optimizer: {
          enabled: true,
          runs: 200
        }
      }
    }
  },

  mocha: {
    timeout: 100000
  }
};