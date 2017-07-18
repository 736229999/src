import axios from "../http";

let addBanner = params => {
  return axios.post('/options/banner', params)
};

let getBannerList = params => {
  return axios.get('/options/banner', { params: params })
}

let getPreBanner = params => {
  return axios.get('/options/banner/pre', { params: params })
}

let getBannerById = params => {
  return axios.get('/options/banner/detail/' + params.id)
}

let updateBanner = params => {
  return axios.put('/options/banner/detail', params)
}

let getFeedbackList = params => {
  return axios.post(`/options/feedback/list`, params)
};
let delFeedbackById = params => {
  return axios.post(`/options/feedback/del`, params)
};
let getFeedbackById = params => {
  return axios.get(`/options/feedback/detail/` + params.id)
};
let updateFeedbackById = params => {
  return axios.post(`/options/feedback/update`, params)
};

let addFaq = params => {
  return axios.post('/options/faq', params)
}

let getFaqList = params => {
  return axios.get('/options/faq', { params: params })
}

let getFaqById = params => {
  return axios.get('/options/faq/detail/' + params.id)
}

let updateFaq = params => {
  return axios.put('/options/faq/detail', params)
}

export default {
  addBanner,
  getFeedbackList,
  delFeedbackById,
  getBannerList,
  getPreBanner,
  getBannerById,
  updateBanner,
  getFeedbackById,
  updateFeedbackById,
  addFaq,
  getFaqList,
  getFaqById,
  updateFaq,
}
