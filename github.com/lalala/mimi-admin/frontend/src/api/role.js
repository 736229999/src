import axios from '../http'

let getRoleList = params => {
  return axios.post(`/user/role/list`, params)
};

export default {
  getRoleList,
}
