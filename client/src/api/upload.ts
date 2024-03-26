import axios from "axios";

const HttpReq = axios.create({
    // baseURL, // api的base_url
    timeout: 3 * 1000 * 60, // 请求超时时间,
    method: "post",
    baseURL: '/blogserver'
});
HttpReq.interceptors.response.use((resp) => {
return resp.data
})
export default {
  GetFiles: function () {
    return HttpReq({
      url:"/getFiles",
      method: 'get'
    })
  },
  GetPacakge: function (url:string) {
    return HttpReq({
      url,
      method: 'get'
    })
  },
  DeletePackage: function (params) {
    return HttpReq({
      url: '/deletePackage',
      method: 'get',
      params
    })
  },
  uploadFileChunk(hash, file) {
    const formData = new FormData()
    formData.append('file', file)
    formData.append('hash', hash)
    return HttpReq('/uploadChunk', {
      method: 'POST',
      data: formData
    })
  },

  checkFileChunkState(hash) {
    return HttpReq({
      url: '/checkChunk?hash=' + hash,
      method: 'get'
    })
  },

  megerChunkFile(hash, fileName, len) {
    return HttpReq({
      url: `/megerChunk?hash=${hash}&fileName=${fileName}&size=${len}`,
      method: 'get'
    })
  },
  saveUploadInfo(data){
    return HttpReq({
      url: `/saveUploadInfo`,
      method: 'post',
      data
    })
  },
  getUploadInfoList(keyword:string){
    return HttpReq({
      url: `/getUploadInfoList`,
      method: 'get',
      params:{
        offset:0,
        size:9999,
        keyword
      }
    })
  }
}
