import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  const user = ref(null)
  const token = ref(localStorage.getItem('token') || null)
  const refreshToken = ref(localStorage.getItem('refresh_token') || null)

  const isAuthenticated = computed(() => !!token.value && !!user.value)
  const userRole = computed(() => user.value?.role_name || null)
  const isAdmin = computed(() => userRole.value === 'Admin')
  const isCustomer = computed(() => userRole.value === 'Customer')

  function setAuth(authData) {
    user.value = authData.user || null
    token.value = authData.access_token || null
    refreshToken.value = authData.refresh_token || null

    if (token.value) localStorage.setItem('token', token.value)
    else localStorage.removeItem('token')

    if (refreshToken.value) localStorage.setItem('refresh_token', refreshToken.value)
    else localStorage.removeItem('refresh_token')

    if (user.value) localStorage.setItem('user', JSON.stringify(user.value))
    else localStorage.removeItem('user')
  }

  function setUser(userData) {
    user.value = userData
    if (user.value) localStorage.setItem('user', JSON.stringify(user.value))
    else localStorage.removeItem('user')
  }

  function logout() {
    user.value = null
    token.value = null
    refreshToken.value = null

    localStorage.removeItem('token')
    localStorage.removeItem('refresh_token')
    localStorage.removeItem('user')
  }

  function loadUserFromStorage() {
    const storedUser = localStorage.getItem('user')
    if (storedUser) {
      try {
        user.value = JSON.parse(storedUser)
      } catch (e) {
        console.error('Gagal memuat pengguna dari localStorage', e)
        logout()
      }
    }
  }

  loadUserFromStorage()

  return {
    user,
    token,
    refreshToken,
    isAuthenticated,
    isAdmin,
    isCustomer,
    userRole,
    setAuth,
    setUser,
    logout,
  }
})
