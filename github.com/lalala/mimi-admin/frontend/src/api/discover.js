import axios from '../http';

let addNews = params => {
     return axios.post('/options/discover/news', params)
}

let getNewsList = params => {
  return axios.get('/options/discover/news', {params: params})
}

let getNewsById = params => {
  return axios.get('/options/discover/news/detail/' + params.id)
}

let updateNews = params => {
    return axios.put('/options/discover/news', params)
}

let queryNewsOfSelect = params => {
  return axios.get('/options/discover/news/select', {params: params})
}

export default {
    addNews,
    getNewsList,
    queryNewsOfSelect,
    getNewsById,
    updateNews,
}