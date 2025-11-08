<template>
  <div class="min-h-screen bg-gradient-to-b from-emerald-50 via-white to-white">
    <!-- Topbar -->
    <header class="sticky top-0 z-40 backdrop-blur supports-[backdrop-filter]:bg-white/70 bg-white/60 border-b border-emerald-100">
      <div class="mx-auto px-4 sm:px-6 lg:px-8 h-16 flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="inline-flex h-9 w-9 items-center justify-center rounded-lg bg-emerald-600 text-white shadow-sm">🏸</div>
          <h1 class="text-base sm:text-lg font-semibold text-emerald-900">Customer Dashboard</h1>
        </div>
        <div class="flex items-center gap-3">
          <div class="hidden sm:flex items-center gap-2 rounded-full bg-emerald-50 text-emerald-800 px-3 py-1 text-sm">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-4 h-4"><path d="M15.75 6a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0z"/><path fill-rule="evenodd" d="M2.25 20.1a8.25 8.25 0 0119.5 0 .75.75 0 01-.75.9H3a.75.75 0 01-.75-.9z" clip-rule="evenodd"/></svg>
            <span class="font-medium">{{ authStore.user?.username }}</span>
          </div>
          <button @click="handleLogout" class="inline-flex items-center gap-2 bg-red-500 text-white px-3 py-2 rounded-lg hover:bg-red-600 shadow-sm">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-4 h-4"><path d="M15.75 9V5.25A2.25 2.25 0 0013.5 3h-6A2.25 2.25 0 005.25 5.25v13.5A2.25 2.25 0 007.5 21h6a2.25 2.25 0 002.25-2.25V15"/><path d="M12 9l3-3m0 0l3 3m-3-3v12"/></svg>
            <span class="hidden sm:inline">Logout</span>
          </button>
        </div>
      </div>
    </header>

    <div class=" mx-auto px-4 sm:px-6 lg:px-8 py-6 grid grid-cols-1 md:grid-cols-12 gap-6">
      <!-- Sidebar -->
      <aside class="md:col-span-3 lg:col-span-3">
        <div class="rounded-xl bg-white border border-emerald-100 shadow-sm p-3">
          <nav class="space-y-1">
            <router-link
              to="/customer/booking"
              class="flex items-center gap-3 px-3 py-2 rounded-lg text-emerald-800 hover:bg-emerald-50 hover:text-emerald-900"
              active-class="bg-emerald-600 text-white hover:bg-emerald-600"
            >
              <span class="text-lg">➕</span>
              <span class="font-medium">Create Booking</span>
            </router-link>
            <router-link
              to="/customer/my-bookings"
              class="flex items-center gap-3 px-3 py-2 rounded-lg text-emerald-800 hover:bg-emerald-50 hover:text-emerald-900"
              active-class="bg-emerald-600 text-white hover:bg-emerald-600"
            >
              <span class="text-lg">📋</span>
              <span class="font-medium">My Booking</span>
            </router-link>
            <router-link
              to="/customer/profile"
              class="flex items-center gap-3 px-3 py-2 rounded-lg text-emerald-800 hover:bg-emerald-50 hover:text-emerald-900"
              active-class="bg-emerald-600 text-white hover:bg-emerald-600"
            >
              <span class="text-lg">👤</span>
              <span class="font-medium">Profile</span>
            </router-link>
          </nav>
        </div>
      </aside>

      <!-- Main Content -->
      <main class="md:col-span-9 lg:col-span-9">
        <div class="rounded-xl bg-white border border-emerald-100 shadow-sm p-4 sm:p-6">
          <router-view />
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const handleLogout = () => {
  authStore.logout()
  router.push('/login')
}
</script>
