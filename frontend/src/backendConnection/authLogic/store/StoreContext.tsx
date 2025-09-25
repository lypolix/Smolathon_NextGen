import React, { createContext, useContext, type ReactNode } from "react"
import { rootStore } from "./store"

type RootStoreType = typeof rootStore

const StoreContext = createContext<RootStoreType>(rootStore)

export const useStore = (): RootStoreType => useContext(StoreContext)

interface StoreProviderProps {
  children: ReactNode
}

export const StoreProvider: React.FC<StoreProviderProps> = ({ children }) => {
  return (
    <StoreContext.Provider value={rootStore}>
      {children}
    </StoreContext.Provider>
  )
}
