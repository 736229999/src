import axios from "../http";

let getWinningList = params => {
  return axios.post(`/lottery/home/winning/list`, params)
};
let winningList = params => {
  return axios.post(`/lottery/open/list`, params)
};
let getWinningByIssue = params => {
  return axios.post(`/lottery/open/search`, params)
};

let winningHistoryList = params => {
  return axios.post(`/lottery/open/history`, params)
}

export default {
  getWinningList,
  winningList,
  getWinningByIssue,
  winningHistoryList,
}
