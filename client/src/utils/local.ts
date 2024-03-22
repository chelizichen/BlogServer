export const constants = {
    BLOG_TOKEN:"blog_server_token",
    USER_NAME:"blog_user_name",
    ENCRYPT_PASSWORD:"blog_encrypt_password",
}
export function localGet(key:string){
    return localStorage.getItem(key)
}

export function localSet(key:string,value:string){
    return localStorage.setItem(key,value)
}

export function localDel(key:string){
    return localStorage.removeItem(key)
}
