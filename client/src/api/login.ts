import type { BasicResp } from '@/dto/dto'
import request from '@/utils/request'


export function LogIn(data:any){
    return request({
        url:'/login',
        method:'post',
        data
    }) as unknown as Promise<BasicResp<any>>
}

export function LoginByCache(data:any){
    return request({
        url:'/loginByCache',
        method:'post',
        data
    }) as unknown as Promise<BasicResp<any>>
}

export function Logout(params:any){
    return request({
        url:'/logout',
        method:'get',
        params
    }) as unknown as Promise<BasicResp<any>>
}