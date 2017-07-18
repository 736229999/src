import moment from "moment";

export default {
  timeStampFormat: function (value) {
    return moment.unix(value).format('YYYY-MM-DD HH:mm:ss')
  },
  switchStatusFormat: function (value) {
    if (value) {
      return '开'
    } else {
      return '关'
    }
  },
  switchPlayTypeFormat: function (value) {
    if (value) {
      return '可以单关投注'
    } else {
      return '不可以单关投注'
    }
  }
}
