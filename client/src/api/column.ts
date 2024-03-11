import type { BasicResp } from '@/dto/dto'
import request from '@/utils/request'

export function getColumns(params: any) {
  return request({
    method: 'get',
    url: 'getColumns',
    params
  }) as unknown as Promise<BasicResp<any>>
}

export function getColumnDetail(params: { id: string }) {
  return request({
    method: 'get',
    url: 'getColumns',
    params
  }) as unknown as Promise<BasicResp<any>>
}

export function saveColumn(data: any) {
  return request({
    method: 'post',
    url: 'saveColumn',
    data
  }) as unknown as Promise<BasicResp<any>>
}


export function saveArticleInColumn(params: any) {
  return request({
    method: 'get',
    url: 'saveArticleInColumn',
    params
  }) as unknown as Promise<BasicResp<any>>
}
