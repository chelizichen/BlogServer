import type { BasicResp } from '@/dto/dto'
import request from '@/utils/request'

export function getArticleList(params: any) {
  return request({
    method: 'get',
    url: 'getArticleList',
    params
  }) as unknown as Promise<BasicResp<any>>
}

export function saveArticle(data: any) {
  return request({
    method: 'post',
    url: 'saveArticle',
    data
  }) as unknown as Promise<BasicResp<any>>
}

export function getArticle(params: any) {
  return request({
    method: 'get',
    url: 'getArticleById',
    params
  }) as unknown as Promise<BasicResp<any>>
}

export function delArticle(params: any) {
  return request({
    method: 'get',
    url: 'delArticleById',
    params
  }) as unknown as Promise<BasicResp<any>>
}

export function getPics(params: any) {
  return request({
    method: 'get',
    url: 'getPics',
    params
  }) as unknown as Promise<BasicResp<any>>
}
