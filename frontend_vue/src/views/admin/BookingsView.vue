<template>
  <div>
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-3 mb-6">
      <div>
        <h2 class="text-2xl font-bold">Data Booking</h2>
        <p class="text-sm text-gray-500">Daftar semua booking pelanggan</p>
      </div>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="text-center py-12 text-gray-500">
      <svg class="animate-spin h-6 w-6 mx-auto mb-2 text-gray-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v4a4 4 0 00-4 4H4z"></path></svg>
      Loading booking...
    </div>

    <!-- Bookings Table -->
    <div v-else class="bg-white rounded-xl shadow overflow-hidden">
      <table class="min-w-full text-sm">
        <thead class="bg-gray-50 text-gray-600">
          <tr>
            <th class="px-6 py-3 text-left font-semibold">Customer</th>
            <th class="px-6 py-3 text-left font-semibold">Court</th>
            <th class="px-6 py-3 text-left font-semibold">Timeslot</th>
            <th class="px-6 py-3 text-left font-semibold">Date</th>
            <th class="px-6 py-3 text-left font-semibold">Total Harga</th>
            <th class="px-6 py-3 text-left font-semibold">Status</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-100">
          <tr v-for="booking in bookings" :key="booking.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">{{ booking.user?.username || 'N/A' }}</td>
            <td class="px-6 py-4">{{ booking.court?.court_name || 'N/A' }}</td>
            <td class="px-6 py-4">
              {{ booking.start_time }} - {{ booking.end_time }}
            </td>
            <td class="px-6 py-4">{{ booking.base_date }}</td>
            <td class="px-6 py-4">Rp {{ formatPrice(booking.total_price) }}</td>
            <td class="px-6 py-4">
              <span :class="getStatusClass(booking.status)">
                {{ booking.status }}
              </span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { bookingAPI } from '@/services/api'

const bookings = ref([])
const loading = ref(false)

const formatPrice = (price) => {
  return new Intl.NumberFormat('id-ID').format(price)
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('id-ID', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

const getStatusClass = (status) => {
  const baseClass = 'px-2 py-1 rounded text-sm'
  switch (status) {
    case 'confirmed':
      return `${baseClass} bg-green-100 text-green-800`
    case 'pending':
      return `${baseClass} bg-yellow-100 text-yellow-800`
    case 'cancelled':
      return `${baseClass} bg-red-100 text-red-800`
    default:
      return `${baseClass} bg-gray-100 text-gray-800`
  }
}

const loadBookings = async () => {
  try {
    loading.value = true
    const response = await bookingAPI.getAll()
    bookings.value = response.data.data
    for (const bookng in bookings.value) {
      const element = bookings.value[bookng];
      element.base_date = new Date(element.timeslot.start_time).toLocaleDateString('id-ID')
      element.start_time = new Date(element.timeslot.start_time).toTimeString().split(' ')[0]
      element.end_time = new Date(element.timeslot.end_time).toTimeString().split(' ')[0]
      const start_split = element.start_time.split(':')
      const start_join = start_split[0]+":"+start_split[1]
      element.start_time = start_join
      const end_split = element.end_time.split(':')
      const end_join = end_split[0]+":"+end_split[1]
      element.end_time = end_join
    }
  } catch (error) {
    alert('Gagal memuat bookings')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadBookings()
})
</script>
