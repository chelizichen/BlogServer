import { useUserStore } from "@/stores/counter";

export const level = {
    name:'level',
    hooks:{
        'beforeMount':function(el,binding,vnode){
            const userStore = useUserStore()
            if(userStore.userInfo.level < binding.value){
                el.style.display = "none"
            }
        }
    }
}