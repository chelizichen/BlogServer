import type { BasicResp } from '@/dto/dto'
import request from '@/utils/request'


export function getUserList(params:any){
    return request({
        url:'/getUserList',
        method:'get',
        params
    }) as unknown as Promise<BasicResp<any>>
}

export function changeUserLevel(data:any){
    return request({
        url:'/changeUserLevel',
        method:'post',
        data
    }) as unknown as Promise<BasicResp<any>>
}


