import axios from '../http';

let getOrderAndIncome = params => {
  return axios.post(`/data/orderandincome`, params)
};
let getUserStatistics = params => {
  return axios.post(`/data/user/statistics`, params)
};
export default {
  getOrderAndIncome,
  getUserStatistics,
}
