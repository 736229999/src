import axios from '../http'

let cdkeyList = params => {
  return axios.post(`/activity/cdkey/list`, params)
};
let giftList = params => {
  return axios.post(`/activity/cdkey/allGiftTemplate`, params)
};
let addCdkeyBatch = params => {
  return axios.post(`/activity/cdkey/add`, params)
};
let getCdkeyDetail = params => {
  return axios.get(`/activity/cdkey/detail/`+params.id)
};
let editCdkeyBatch = params => {
  return axios.post('/activity/cdkey/update', params)
};
let delCdkeyById = params => {
  return axios.post( `/activity/cdkey/delete`, params)
};
let exportCsv =  params => {
  return axios.post(`/activity/cdkey/export`, params)
};
let addTask =  params => {
  return axios.post(`/activity/task/addTask`, params)
};
let taskList =  params => {
  return axios.post(`/activity/task/taskList`, params)
};
let allTask =  params => {
  return axios.post(`/activity/task/allTask`, params)
};
let deleteTask = params => {
  return axios.post( `/activity/task/delete`, params)
};
let queryTaskById = params => {
  return axios.post( `/activity/task/detail`, params)
};
let updateTask = params => {
  return axios.post( `/activity/task/update`, params)
};
let taskTypes = params => {
  return axios.post( `/activity/task/taskTypes`, params)
};
let addActivity =  params => {
  return axios.post(`/activity/activity/addActivity`, params)
};
let activityList =  params => {
  return axios.post(`/activity/activity/activityList`, params)
};
let editActivity = params => {
  return axios.post('/activity/activity/detail', params)
}
let deleteActivity = params => {
  return axios.post('/activity/activity/delete', params)
}
let updateActivity = params => {
  return axios.post('/activity/activity/update', params)
}
let getGiftTemplateList = params => {
  return axios.post(`/activity/gift/template/list`, params)
};

export default {
  cdkeyList,
  giftList,
  addCdkeyBatch,
  getCdkeyDetail,
  editCdkeyBatch,
  delCdkeyById,
  exportCsv,
  addTask,
  taskList,
  deleteTask,
  updateTask,
  taskTypes,
  queryTaskById,
  allTask,
  addActivity,
  activityList,
  editActivity,
  deleteActivity,
  updateActivity,
  getGiftTemplateList,
}
