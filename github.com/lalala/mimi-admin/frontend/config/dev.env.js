var merge = require('webpack-merge')
var prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  API_ROOT: 'http://127.0.0.1:11111',
  ASSETS_API_ROOT: 'http://cptest.kxkr.com:8088',
  // ASSETS_API_ROOT: 'http://127.0.0.1:7009',
});
