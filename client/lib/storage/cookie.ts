import { CookieKey } from "@/constants/common";
import { User } from "@/types/user.type";

export const CookiesStorage = {
  getCookieData(key: string) {
    return localStorage.getItem(key);
  },

  setCookieData(key: string, data: string) {
    localStorage.setItem(key, data);
  },

  clearCookieData(key: string) {
    localStorage.removeItem(key);
  },

  // Access Token
  getAccessToken() {
    return localStorage.getItem(CookieKey.accessToken)
  },

  setAccessToken(accessToken: string) {
    localStorage.setItem(CookieKey.accessToken, accessToken)
  },

  clearAccessToken() {
    localStorage.removeItem(CookieKey.accessToken)
  },

  // Refresh Token
  getRefreshToken() {
    return localStorage.getItem(CookieKey.refreshToken)
  },

  setRefreshToken(refreshToken: string) {
    localStorage.setItem(CookieKey.refreshToken, refreshToken)
  },

  clearRefreshToken() {
    localStorage.removeItem(CookieKey.refreshToken)
  },

  // User
  getUser() {
    if (typeof window !== 'undefined') {
      const user = localStorage.getItem(CookieKey.user);
      return user ? JSON.parse(user) : null;
    }
    return null;
  },

  setUser(data: User) {
    return localStorage.setItem(CookieKey.user, JSON.stringify(data))
  },

  clearUser() {
    localStorage.removeItem(CookieKey.user)
  },
}