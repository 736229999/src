import axios from "../http";

let lotteryList = params => {
  return axios.get(`/lottery/list`, params)
};

let buycaiOptions = params => {
  return axios.post(`/lottery/buycai/options`, params)
};

//获取期号数据.
let buycaioptionsIssue = params => {
  return axios.post(`/lottery/issue`, params)
};

// 根据彩种获取期号.
let buycaiOptionsGetIssueByLottery = params => {
  return axios.post(`/lottery/newissue`, params)
};

let buycaiOptionsAddIssue = params => {
  return axios.post(`/lottery/buycai/options/add`, params)
};

let playTiemSettingGetLottery = params => {
  return axios.post(`/lottery/playtime/list`, params)
};

let addPlayTime = params => {
  return axios.post(`/lottery/playtimesetting/add`, params)
};

let updatePlayTimeSetting = params => {
  return axios.post(`/lottery/playtimesetting/update`, params)
};
let lotteryOptions = params => {
  return axios.post(`/lottery/home/options/add`, params)
};
let lotteryOptionsList = params => {
  return axios.post(`/lottery/home/options/list`, params)
};
let getLotteryOptionsById = params => {
  return axios.get('/lottery/home/options/' + params.id)
};
let uploadLotteryOptions = params => {
  return axios.post(`/lottery/home/options/edit`, params)
};
let getLotteryOptionsNotAddList = params => {
  return axios.post(`/lottery/home/options/notaddlist`, params)
};
let updateContactInfo = params => {
  return axios.post(`/options/update/contact`, params)
};
let getContact = params => {
  return axios.post(`/options/query/contact`, params)
};
let delIssueById = params => {
  return axios.post(`/lottery/buycai/options/delete`, params)
};
let getLotteryBuycaiOptionsById = params => {
  return axios.get(`/lottery/buycai/options/detail/`+params.id+`/lottery/`+params.lottery)
};
let updateBuycaiOptions = params => {
  return axios.post(`/lottery/buycai/options/update`, params)
};
let initBuycaiOptions = params => {
  return axios.post(`/lottery/buycai/options/init`, params)
};

export default {
  lotteryList,
  buycaiOptions,
  buycaioptionsIssue,
  buycaiOptionsGetIssueByLottery,
  buycaiOptionsAddIssue,
  playTiemSettingGetLottery,
  addPlayTime,
  updatePlayTimeSetting,
  lotteryOptions,
  lotteryOptionsList,
  getLotteryOptionsById,
  uploadLotteryOptions,
  getLotteryOptionsNotAddList,
  updateContactInfo,
  getContact,
  delIssueById,
  getLotteryBuycaiOptionsById,
  updateBuycaiOptions,
  initBuycaiOptions,
}
