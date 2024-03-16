import type { BasicResp } from '@/dto/dto'
import request from '@/utils/node_req'

export function getEvents(params: any) {
  return request({
    method: 'get',
    url: 'getEvents',
    params
  }) as unknown as Promise<BasicResp<any>>
}

export function saveEvent(data: any) {
  return request({
    method: 'post',
    url: 'saveEvent',
    data
  }) as unknown as Promise<BasicResp<any>>
}


export function changeStatus(data: any) {
  return request({
    method: 'post',
    url: 'changeStatus',
    data
  }) as unknown as Promise<BasicResp<any>>
}

export function getCommentsByEventId(params: any) {
  return request({
    method: 'get',
    url: 'getCommentsByEventId',
    params
  }) as unknown as Promise<BasicResp<any>>
}
