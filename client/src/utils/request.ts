import router from '@/router'
import axios from 'axios'
import { constants, localGet } from './local'

const HttpReq = axios.create({
  // baseURL, // api的base_url
  timeout: 15000, // 请求超时时间,
  method: 'post',
  baseURL: '/blogserver/'
})

HttpReq.interceptors.response.use((resp) => {
  if (resp.data.Code != 0) {
    console.log('resp.data.Message |', resp.data.Message)
    // this.$message.error(resp.data.Message)
  }
  return resp.data
})

HttpReq.interceptors.request.use((config) => {
  const tkn = localGet(constants.BLOG_TOKEN)
  config.headers['Token'] = tkn
  return config
})

export default HttpReq
