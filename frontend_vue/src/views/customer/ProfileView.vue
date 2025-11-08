<template>
  <div class="space-y-6">
    <h2 class="text-xl sm:text-2xl font-bold text-emerald-900">Profile Saya</h2>

    <!-- Loading Skeleton -->
    <div v-if="profileLoading" class="bg-white rounded-xl border border-emerald-100 shadow-sm p-6 max-w-2xl animate-pulse">
      <div class="flex items-center gap-4 mb-6">
        <div class="h-14 w-14 rounded-full bg-emerald-100"></div>
        <div class="space-y-2 flex-1">
          <div class="h-4 bg-emerald-100 rounded w-1/2"></div>
          <div class="h-3 bg-emerald-50 rounded w-1/3"></div>
        </div>
      </div>
      <div class="grid gap-4">
        <div class="h-10 bg-emerald-50 rounded"></div>
        <div class="h-10 bg-emerald-50 rounded"></div>
        <div class="h-10 bg-emerald-50 rounded"></div>
      </div>
    </div>

    <div v-else class="bg-white rounded-xl border border-emerald-100 shadow-sm p-6 max-w-2xl">
      <div class="flex items-center gap-4 mb-6">
        <div class="h-14 w-14 rounded-full bg-emerald-100 text-emerald-800 flex items-center justify-center font-bold">
          {{ usernameInitials }}
        </div>
        <div>
          <p class="text-emerald-900 font-semibold leading-tight">{{ form.username || '—' }}</p>
          <p class="text-sm text-gray-600">Kelola informasi akun Anda</p>
        </div>
      </div>

      <form @submit.prevent="updateProfile" class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-emerald-900 mb-1">Username</label>
          <input
            v-model="form.username"
            type="text"
            class="block w-full rounded-lg border border-emerald-200 px-3 py-2 text-emerald-900 placeholder:text-emerald-400 focus:outline-none focus:ring-2 focus:ring-emerald-500/50 focus:border-emerald-500"
            required
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-emerald-900 mb-1">Email</label>
          <input
            v-model="form.email"
            type="email"
            class="block w-full rounded-lg border border-emerald-200 px-3 py-2 text-emerald-900 placeholder:text-emerald-400 focus:outline-none focus:ring-2 focus:ring-emerald-500/50 focus:border-emerald-500"
            required
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-emerald-900 mb-1">Kontak</label>
          <input
            v-model="form.kontak"
            type="text"
            class="block w-full rounded-lg border border-emerald-200 px-3 py-2 text-emerald-900 placeholder:text-emerald-400 focus:outline-none focus:ring-2 focus:ring-emerald-500/50 focus:border-emerald-500"
            required
          />
        </div>

        <div class="pt-4 mt-6">
          <h3 class="text-base font-semibold text-emerald-900 mb-3">Ubah Password</h3>

          <label class="block text-sm font-medium text-emerald-900 mb-1">Password Baru</label>
          <div class="relative group">
            <input
              v-model="form.password"
              :type="showPassword ? 'text' : 'password'"
              class="block w-full rounded-lg border border-emerald-200 px-3 py-2 pr-24 text-emerald-900 placeholder:text-gray-400 focus:outline-none focus:ring-2 focus:ring-emerald-500/50 focus:border-emerald-500"
              placeholder="Kosongkan untuk mempertahankan password saat ini"
            />
            <button
              type="button"
              @click="showPassword = !showPassword"
              class="absolute inset-y-1 right-2 my-1 mx-1 h-8 w-8 grid place-items-center rounded-md text-emerald-700 hover:bg-emerald-50 hover:text-emerald-900"
              :aria-label="showPassword ? 'Sembunyikan password' : 'Tampilkan password'"
              >
              <span class="sr-only">Toggle password</span>
              <svg v-if="showPassword" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" class="w-5 h-5">
                <path d="M1 12s4-7 11-7 11 7 11 7-4 7-11 7-11-7-11-7Z"/>
                <circle cx="12" cy="12" r="3"/>
              </svg>
              <svg v-else xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" class="w-5 h-5">
                <path d="M17.94 17.94A10.94 10.94 0 0 1 12 19c-7 0-11-7-11-7a20.29 20.29 0 0 1 5.06-5.94M9.9 4.24A10.94 10.94 0 0 1 12 4c7 0 11 7 11 7a20.3 20.3 0 0 1-3.23 4.62"/>
                <path d="M1 1l22 22"/>
              </svg>
              <span class="pointer-events-none absolute -top-9 right-0 whitespace-nowrap rounded-md bg-emerald-500/90 px-2 py-1 text-xs font-medium text-white opacity-0 transition group-hover:opacity-100">
                {{ showPassword ? 'Sembunyikan' : 'Tampilkan' }}
              </span>
            </button>
          </div>
          <p class="mt-1 text-xs text-gray-600">Biarkan kosong jika tidak ingin mengubah password.</p>
        </div>

        <div class="flex justify-end gap-2 pt-2">
          <button
            type="button"
            @click="loadProfile"
            class="px-4 py-2 rounded-lg border border-emerald-200 text-emerald-700 hover:bg-emerald-50"
          >
            Reset
          </button>
          <button
            type="submit"
            class="px-4 py-2 rounded-lg bg-emerald-600 text-white hover:bg-emerald-700 disabled:opacity-60"
            :disabled="loading"
          >
            {{ loading ? 'Menyimpan…' : 'Simpan Perubahan' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { userAPI } from '@/services/api'

const authStore = useAuthStore()
const loading = ref(false)
const profileLoading = ref(true)
const form = ref({
  username: '',
  email: '',
  kontak: '',
  password: ''
})
const showPassword = ref(false)

const usernameInitials = computed(() => {
  const name = form.value.username?.trim()
  if (!name) return '👤'
  return name.slice(0, 2).toUpperCase()
})

const loadProfile = async () => {
  try {
    const response = await userAPI.getById(authStore.user.id)
    const user = response.data.data

    form.value = {
      username: user.username,
      email: user.email,
      kontak: user.kontak,
      password: ''
    }
    profileLoading.value = false
  } catch (error) {
    alert('Gagal memuat profil')
    profileLoading.value = false
  }
}

const updateProfile = async () => {
  try {
    loading.value = true
    
    const updateData = {
      username: form.value.username,
      email: form.value.email,
      kontak: form.value.kontak
    }
    
    if (form.value.password) {
      updateData.password = form.value.password
    }
    
    const response = await userAPI.update(authStore.user.id, updateData)
    authStore.setUser(response.data.data)

    alert('Profile berhasil diperbarui!')
    form.value.password = '' // Clear password field
  } catch (error) {
    alert(error.response?.data?.message || 'Gagal memperbarui profil')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadProfile()
})
</script>
