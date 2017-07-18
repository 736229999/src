import axios from '../http'

let userLogin = params => {
    return axios.post( `/login`, params )
};
let userLogout = params => {
    return axios.get( `/logout`, { params: params });
};

export default {
    userLogin,
    userLogout,
}
