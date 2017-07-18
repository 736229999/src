import axios from '../http'


let getUserOrderList = params => {
  return axios.post(`/order/user`, params)
};
let searchUserOrder = params => {
  return axios.post(`/order/select`, params)
};

export default {
  getUserOrderList,
  searchUserOrder,
}
