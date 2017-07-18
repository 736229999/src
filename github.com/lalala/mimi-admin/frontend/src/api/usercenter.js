import axios from '../http'

let getUserList = params => {
  return axios.post(`/usercenter/list`, params)
};
let getUserDetail = params => {
  return axios.get(`/usercenter/detail/` + params.id)
};
let getFund = params => {
  return axios.get(`/usercenter/fund/` + params.id)
};
let getWithdrawList = params => {
  return axios.get(`/usercenter/withdraw/` + params.id)
};
let getWithdrawApplyList = params => {
  return axios.get('/usercenter/withdrawapply', { params: params })
};
let getWithdrawApplyDetail = params => {
  return axios.get('/usercenter/withdrawapply/detail/' + params.id)
}
let updateWithdrawApplyStatus = params => {
  return axios.put('/usercenter/withdrawapply/detail', params)
}
let claimWithdrawApply = params => {
  return axios.put('/usercenter/withdrawapply/claim', params)
}

let checkWithdrawApply = params => {
  return axios.get('usercenter/withdrawapply/check/' + params.id)
}

let addWithdrawAuditAuth = params => {
  return axios.post('/usercenter/withdraw/auth', params)
}

let getWithdrawAuditAuthList = params => { 
  return axios.get('/usercenter/withdraw/auth', { params: params })
}

let getWithdrawAuditAuthDetail = params => {
  return axios.get('/usercenter/withdraw/auth/detail/' + params.id)
}

let updateWithdrawAuditAuth = params => {
  return axios.put('/usercenter/withdraw/auth/detail', params)
}

let getWithdrawTransferList = params => {
  return axios.get('/usercenter/withdrawtransfer', { params: params })
}

export default {
  getUserList,
  getUserDetail,
  getFund,
  getWithdrawList,
  getWithdrawApplyList,
  getWithdrawApplyDetail,
  updateWithdrawApplyStatus,
  claimWithdrawApply,
  checkWithdrawApply,
  addWithdrawAuditAuth,
  getWithdrawAuditAuthList,
  getWithdrawAuditAuthDetail,
  updateWithdrawAuditAuth,
  getWithdrawTransferList,
}
