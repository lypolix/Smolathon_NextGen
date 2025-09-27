import { makeAutoObservable } from "mobx"
import AuthService from "../services/AuthService"
import axios from "axios"
import type { AuthResponse } from "../model/AuthResponse"
import { API_URL } from "../http"
import type { Admin1 } from "../../../types"

class Store {
  user = {} as Admin1
  isAuth = false

  constructor() {
    makeAutoObservable(this)
  }

  setAuth(bool: boolean) {
    this.isAuth = bool
  }

  setUser(user: Admin1) {
    this.user = user
  }

  async login(email: string, password: string) {
    try {
      const response = await AuthService.loginAdmin(email, password)
      console.log(response)
      localStorage.setItem("token", response.data.token)
      this.setAuth(true)
    } catch (error: any) {
      console.log(error.response?.data?.message)
    }
  }

  async logout() {
    try {
      await AuthService.logout()
      localStorage.removeItem("token")
      this.setAuth(false)
      this.setUser({} as Admin1)
    } catch (error: any) {
      console.log(error.response?.data?.message)
    }
  }

  async checkAuth() {
    console.log("начинаем проверять авторизацию")
    try {
      await axios.get<AuthResponse>(`${API_URL}/refresh`, {
        withCredentials: true,
      })
    } catch (error: any) {
      console.log(error.response?.data?.message, "ошибка в store")
    }
  }
}

export const rootStore = new Store()   
export default Store
