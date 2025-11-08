<template>
  <div class="min-h-screen bg-gray-100">
    <!-- Navbar -->
    <nav class="bg-blue-600 text-white p-4">
      <div class="container mx-auto flex justify-between items-center">
        <h1 class="text-xl font-bold">🏸 Admin Dashboard</h1>
        <div class="flex items-center gap-4">
          <span>{{ authStore.user?.username }}</span>
          <button
            @click="handleLogout"
            class="bg-red-500 px-4 py-2 rounded hover:bg-red-600"
          >
            Logout
          </button>
        </div>
      </div>
    </nav>

    <div class="container mx-auto flex">
      <!-- Sidebar -->
      <aside class="w-64 bg-white h-screen p-4">
        <nav class="space-y-2">
          <router-link
            to="/admin/courts"
            class="block px-4 py-2 rounded hover:bg-blue-100"
            active-class="bg-blue-500 text-white"
          >
            🏸 Manage Court
          </router-link>
          <router-link
            to="/admin/timeslots"
            class="block px-4 py-2 rounded hover:bg-blue-100"
            active-class="bg-blue-500 text-white"
          >
            ⏰ Manage Timeslot
          </router-link>
          <router-link
            to="/admin/bookings"
            class="block px-4 py-2 rounded hover:bg-blue-100"
            active-class="bg-blue-500 text-white"
          >
            📋 All Booking
          </router-link>
        </nav>
      </aside>

      <!-- Main Content -->
      <main class="flex-1 p-6">
        <router-view />
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
