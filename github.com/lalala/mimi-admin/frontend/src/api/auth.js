import axios from '../http'

let addAuth = params => {
  return axios.post("/user/privilege/add", params)
};
let getAuthList = params => {
  return axios.post("/user/privilege/list", params)
};

export default {
  addAuth,
  getAuthList,
}
