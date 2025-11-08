<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <h2 class="text-xl sm:text-2xl font-bold text-emerald-900">Bookingan Saya</h2>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="grid gap-4">
      <div v-for="i in 3" :key="i" class="bg-white rounded-xl border border-emerald-100 shadow-sm p-6 animate-pulse">
        <div class="h-4 bg-emerald-100 rounded w-1/3 mb-4"></div>
        <div class="grid grid-cols-2 gap-4">
          <div class="h-3 bg-emerald-50 rounded"></div>
          <div class="h-3 bg-emerald-50 rounded"></div>
          <div class="h-3 bg-emerald-50 rounded"></div>
          <div class="h-3 bg-emerald-50 rounded"></div>
        </div>
      </div>
    </div>

    <!-- No Bookings -->
    <div v-else-if="bookings.length === 0" class="text-center py-16 bg-white rounded-xl border border-emerald-100 shadow-sm">
      <div class="h-12 w-12 mx-auto mb-3 rounded-xl bg-emerald-100 text-emerald-700 flex items-center justify-center">📭</div>
      <p class="text-emerald-900 font-semibold">Belum ada booking</p>
      <p class="text-gray-600 text-sm">Ayo buat booking pertama Anda di halaman Create Booking.</p>
    </div>

    <!-- Bookings Grid -->
    <div v-else class="grid md:grid-cols-2 lg:grid-cols-3 gap-4">
      <div
        v-for="booking in bookings"
        :key="booking.id"
        class="bg-white rounded-xl border border-emerald-100 shadow-sm p-6 hover:shadow-md transition-shadow"
      >
        <div class="flex justify-between items-start mb-4">
          <div>
            <h3 class="text-base sm:text-lg font-semibold text-emerald-900">{{ booking.court?.court_name }}</h3>
            <p class="text-gray-600 text-sm">{{ booking.court?.type }} • {{ booking.court?.location }}</p>
          </div>
          <span :class="getStatusClass(booking.status)">
            {{ booking.status }}
          </span>
        </div>

        <div class="grid grid-cols-2 gap-4 text-sm">
          <div>
            <p class="text-gray-600">Tanggal</p>
            <p class="font-medium">{{ formatDate(booking.booking_date) }}</p>
          </div>
          <div>
            <p class="text-gray-600">Waktu</p>
            <p class="font-medium">
              {{ getTimeOnly(booking.timeslot?.start_time) }} - {{ getTimeOnly(booking.timeslot?.end_time) }}
            </p>
          </div>
          <div>
            <p class="text-gray-600">Total Harga</p>
            <p class="font-semibold text-emerald-700 text-lg">
              Rp {{ formatPrice(booking.total_price) }}
            </p>
          </div>
          <div v-if="booking.status === 'pending'" class="flex items-end justify-end">
            <button
              @click="cancelBooking(booking.ID)"
              class="px-3 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600 text-sm"
            >
              Batalkan
            </button>
          </div>
        </div>
      </div>
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
    weekday: 'long',
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

const getTimeOnly = (time) => {
  const timeOnly = new Date(time).toTimeString().split(' ')[0]
  const splitedTime = timeOnly.split(':')
  return splitedTime[0]+":"+splitedTime[1]
}
const getStatusClass = (status) => {
  const baseClass = 'px-3 py-1 rounded-full text-xs font-semibold'
  switch (status) {
    case 'confirmed':
      return `${baseClass} bg-emerald-100 text-emerald-800`
    case 'pending':
      return `${baseClass} bg-amber-100 text-amber-800`
    case 'cancelled':
      return `${baseClass} bg-rose-100 text-rose-800`
    default:
      return `${baseClass} bg-gray-100 text-gray-800`
  }
}

const loadBookings = async () => {
  try {
    loading.value = true
    const response = await bookingAPI.getAll()
    bookings.value = response.data.data
  } catch (error) {
    alert('Failed to load bookings')
  } finally {
    loading.value = false
  }
}

const cancelBooking = async (id) => {
  if (!confirm('Apakah Anda yakin ingin membatalkan booking ini?')) return

  try {
    await bookingAPI.update(id, { status: 'cancelled' })
    await loadBookings()
    alert('Booking berhasil dibatalkan!')
  } catch (error) {
    alert(error.response?.data?.message || 'Gagal membatalkan booking')
  }
}

onMounted(() => {
  loadBookings()
})
</script>
