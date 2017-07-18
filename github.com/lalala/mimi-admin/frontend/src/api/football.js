import axios from '../http';
//比赛
let getGamesList = params => {
  return axios.get('/football/games', {params: params})
}

let getGameById = params => {
  return axios.get('/football/games/detail/' + params.id)
}

let addGame = params => {
  return axios.post('/football/games/add', params)
}

let updateGame = params => {
  return axios.put('/football/games/update', params)
}

let queryLeaguesOfSelect = params => {
  return axios.get('/football/league/select', {params: params})
}

let queryTeamsOfSelect = params => {
  return axios.get('/football/team/select', {params: params})
}
//赔率
let getOddbyId = params => {
  return axios.get('/football/odds/detail/' + params.id)
}

let updateOdds = params=>{
  return axios.put('/football/odds/update', params)
}

let addOdds = params => {
  return axios.post('/football/odds/add', params)
}
//开彩
let getOpencaiList = params => {
  return axios.get('/football/opencai', {params: params})
}

let getOpencaiById = params => {
  return axios.get('/football/opencai/detail/' + params.id)
}

let addOpencai= params => {
  return axios.post('/football/opencai/add', params)
}

let updateOpencai = params => {
  return axios.put('/football/opencai/update', params)
}
//球队l
let getTeamList = params => {
  return axios.get('/football/team', {params: params})
}

let addTeam = params => {
  return axios.post('/football/team/add', params)
}

let getTeamById = params => {
  return axios.get('/football/team/detail/' + params.id)
}

let updateTeam = params => {
  return axios.put('/football/team/update', params)
}

let getLeagueList= params => {
  return axios.get('/football/league', {params: params})
}

let addLeague = params => {
  return axios.post('/football/league/add', params)
}

let updateLeague = params => {
  return axios.put('/football/league/update', params)
}

export default {
  getGamesList,
  queryLeaguesOfSelect,
  queryTeamsOfSelect,
  addGame,
  getGameById,
  updateGame,
  getOddbyId,
  updateOdds,
  addOdds,
  getOpencaiList,
  getOpencaiById,
  addOpencai,
  updateOpencai,
  getTeamList,
  addTeam,
  getTeamById,
  updateTeam,
  getLeagueList,
  addLeague,
  updateLeague,
}
