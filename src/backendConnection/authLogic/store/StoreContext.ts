import {createContext, useContext} from "react";
import type Store from "./store";

export const StoreContext=createContext<Store | null>(null)
export const useStore=()=>{
    const context=useContext(StoreContext);
    if(!context){
        throw new Error("useStore must be within storeContext.Provider")
    }
    return context
}