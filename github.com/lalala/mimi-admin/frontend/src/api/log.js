import axios from '../http'

let logList = params => {
    return axios.post(`/log/list`, params)
};

export default {
    logList
}
