import axios from "../http";

let getBuycaiByDay = params => {
  return axios.post(`/data/buycai/day`, params)
};
export default {
  getBuycaiByDay,
}
