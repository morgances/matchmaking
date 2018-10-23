import axios from 'axios'

const service = axios.create({
  baseURL: 'https://techcats.s8e.io/api', // api的base_url
  timeout: 15000 // 请求超时时间
})

service.interceptors.request.use(config => {
  console.log(config)
})

service.interceptors.response.use(config => {
  console.log(config)
})

export default service
