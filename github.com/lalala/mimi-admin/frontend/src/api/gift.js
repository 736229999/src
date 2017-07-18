import axios from '../http'

let giftTemplateList = params => {
  return axios.post(`/gift/template/list`, params)
};
let giftTypeList = params => {
  return axios.post(`/gift/type/list`, params)
};
let addGift = params => {
  return axios.post(`/gift/template/add`, params)
};
let getGiftDetail = params => {
  return axios.get(`/gift/template/`+params.id)
};
let delGift = params => {
  return axios.post(`/gift/template/delete`, params)
};
let updateGift = params => {
  return axios.post(`/gift/template/update`, params)
};
let getLotteryTypeList = params => {
  return axios.post(`/lottery/type/list`, params)
};

export default {
  giftTemplateList,
  giftTypeList,
  addGift,
  getGiftDetail,
  delGift,
  updateGift,
  getLotteryTypeList,
}
