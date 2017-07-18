 import axios from '../http';

// export const userLogin = params => {
//   return axios.post( `/login`, params )
// };
export const userLogout = params => {
  return axios.get( `/logout`, { params: params });
};
export const userChangepwd = params => {
  return axios.post( `/u/user/changepwd`, params )
};
export const userUpdate = params => {
  return axios.post( `/user/update`, params )
};

// 角色列表，添加角色，删除角色，编辑橘色，角色权限
export const roleList = params => {
  return axios.get( `/user/roles/list`, { params: params } )
};
export const roleListSimple = params => {
  return axios.get( `/u/roles/simple`, { params: params } )
};
export const addRole = params => {
  return axios.post( `/user/roles/add`, params )
};
export const deleteRole = params => {
  return axios.delete( `/u/delete/role`, { params: params } )
};
export const editRole = params => {
  return axios.put( `/u/update/role`, params )
};
export const detailRole = params => {
  return axios.get( `/u/detail/role`, { params: params } )
};
// 用户列表，添加用户，删除用户，编辑用户
export const userList = params => {
  return axios.get( `/user/list`, { params: params } )
};
export const addUser = params => {
  return axios.post( `/u/add/user`, params )
};
export const deleteUser = params => {
  return axios.delete( `/u/delete/user`, { params: params } )
};
export const editUser = params => {
  return axios.put( `/u/update/user`, params )
};

// 获取所有权限，简要权限
export const privilegeList =  params => {
  return axios.get( `/user/privileges/list`, { params: params } )
};
export const addPrivilege =  params => {
    return axios.post( `/user/privileges/add`, params)
};
export const privilegeListSimple =  params => {
  return axios.get( `/u/privileges/simple`, { params: params } )
};
export const editPrivileges =  params => {
  return axios.post( `/user/privileges/edit`, params )
};

export const deletePrivilege =  params => {
  return axios.post( `/user/privileges/delete`, params )
};


//以下为关于彩票开奖的模块
//彩票开奖
export const getLotteryListInfo = params => {
  return axios.post(`/u/lottery`, params)
};
//获取地区json.
//控制台地图使用.
export const getAddressJson = params => {
  return axios.get("http://echarts.baidu.com/asset/map/json/china.json")
};


//彩票.
export const GetLotteryList = params => {
    return axios.post(`/lottery/list`)
};
export const BuycaiOptions = params => {
    return axios.post(`/lottery/buycai/options`, params)
};
export const GetLotteryIssue = params => {
    return axios.post(`/lottery/issue`, params)
};


export const AddCdKey = params => {
    return axios.post(`/activity/cdkey/add`, params)
};

//新闻相关
export const addNews = params => {
  // console.log("params:", params)
  return axios.post('/news', params)
};
export const NewsList = params => {
  return axios.get('/news', {params: params})
}
export const AddInvite = params => {
    return axios.post(`/activity/invite/add`, params)
};
export const CdKeyList = params => {
    return axios.get(`/activity/cdkey/list`, {params: params})
};
export const deleteCdkey = params => {
    return axios.post(`/activity/cdkey/del`, params)
};

//日志.
export const LogList = params => {
  return axios.get(`/log/list`, {params: params})
};
