import axios from '../http'

let getRechargeList = params => {
  return axios.post(`/data/recharge/list`, params)
};
let getRechargeListByMonth = params => {
  return axios.post(`/data/recharge/month`, params)
};
let getRechargeListByYear = params => {
  return axios.post(`/data/recharge/year`, params)
};
export default {
  getRechargeList,
  getRechargeListByMonth,
  getRechargeListByYear,
}
