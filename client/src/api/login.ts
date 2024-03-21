import type { BasicResp } from '@/dto/dto'
import request from '@/utils/request'


export function LogIn(data:any){
    return request({
        url:'/login',
        method:'post',
        data
    }) as unknown as Promise<BasicResp<any>>
}