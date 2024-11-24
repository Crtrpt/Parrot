import {GET, Resp} from "./request"

export const GetMenuList = async ():Promise<Resp> => {
    return GET("/menu/list", { name: "top" });
};