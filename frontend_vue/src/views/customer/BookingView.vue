<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <h2 class="text-xl sm:text-2xl font-bold text-emerald-900">Buat Bookingan Baru</h2>
    </div>

    <div class="max-w-5xl mx-auto grid lg:grid-cols-3 gap-6">
      <!-- Form Card -->
      <div class="lg:col-span-2">
        <div class="bg-white rounded-xl border border-emerald-100 shadow-sm p-4 sm:p-6">
          <form @submit.prevent="createBooking" class="space-y-5">
            <div class="grid sm:grid-cols-2 gap-4 items-start">

              <div class="sm:col-span-1">
                <label class="block text-sm font-medium text-gray-700 mb-1">Tanggal Booking</label>
                <input
                  v-model="form.booking_date"
                  @change="handleDateChange"
                  type="date"
                  class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-emerald-500"
                  :min="minDate"
                  required
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">Timeslot</label>
                <select
                  v-model.number="form.timeslot_id"
                  @change="handleTimeslotChange"
                  class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-emerald-500 disabled:bg-gray-50"
                  :disabled="disabledTimeslot"
                  required
                >
                  <option value="">Pilih Timeslot</option>
                  <option
                    v-for="timeslot in displayTimeslots"
                    :key="timeslot.id ?? timeslot.ID"
                    :value="timeslot.id ?? timeslot.ID"
                    :disabled="isTimeslotBooked(timeslot)"
                  >
                    {{ timeslot.display_start }} - {{ timeslot.display_end }}
                    {{ isTimeslotBooked(timeslot) ? ' (Booked)' : '' }}
                  </option>
                </select>
              </div>
              
              <div class="sm:col-span-2">
                <label class="block text-sm font-medium text-gray-700 mb-1">Court</label>
                <select
                  v-model.number="form.court_id"
                  class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-emerald-500 disabled:bg-gray-50"
                  :disabled="disabledCourtSelect"
                  required
                >
                  <option value="">Pilih Court</option>
                  <option
                    v-for="court in courts"
                    :key="court.id ?? court.ID"
                    :value="court.id ?? court.ID"
                    :disabled="isCourtBooked(court)"
                  >
                    {{ court.court_name }} - {{ court.type }} ({{ court.location }}){{ isCourtBooked(court) ? ' - (Booked)' : '' }}
                  </option>
                </select>
              </div>

            </div>

            <div class="flex justify-end gap-2 pt-2">
              <button
                type="button"
                @click="resetForm"
                class="px-4 py-2 border rounded-lg hover:bg-gray-50"
              >
                Reset
              </button>
              <button
                type="submit"
                class="px-4 py-2 bg-emerald-600 text-white rounded-lg hover:bg-emerald-700"
                :disabled="loading"
              >
                {{ loading ? 'Creating...' : 'Create Booking' }}
              </button>
            </div>
          </form>
        </div>
      </div>

      <!-- Summary Card -->
      <div class="lg:col-span-1">
        <div class="bg-gradient-to-b from-emerald-50 to-white rounded-xl border border-emerald-100 shadow-sm p-4 sm:p-6">
          <div class="flex items-center gap-2 mb-4">
            <div class="h-8 w-8 rounded-lg bg-emerald-600 text-white flex items-center justify-center">ℹ️</div>
            <h3 class="font-semibold text-emerald-900">Ringkasan</h3>
          </div>
          <div class="space-y-3 text-sm">
            <div class="flex justify-between"><span class="text-gray-600">Court</span><span class="font-medium">{{ selectedCourtLabel }}</span></div>
            <div class="flex justify-between"><span class="text-gray-600">Timeslot</span><span class="font-medium">{{ selectedTimeslotLabel }}</span></div>
            <div class="flex justify-between"><span class="text-gray-600">Tanggal</span><span class="font-medium">{{ formatDateLocal(form.booking_date) || '-' }}</span></div>
            <div v-if="estimatedPrice > 0" class="pt-3 mt-2 border-t">
              <p class="text-sm text-gray-600 mb-1">Estimasi Harga</p>
              <p class="text-2xl font-bold text-emerald-700">Rp {{ formatPrice(estimatedPrice) }}</p>
              <p class="text-xs text-gray-500 mt-1">Tarif: {{ isWeekend ? 'Weekend' : 'Weekday' }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { courtAPI, timeslotAPI, bookingAPI } from '@/services/api'

const router = useRouter()
const courts = ref([])
const timeslots = ref([])
const bookedCourtIds = ref(new Set())
const bookedTimeslotCourts = ref(new Map())
const loading = ref(false)
const estimatedPrice = ref(0)
const isWeekend = ref(false)

const form = ref({
  court_id: null,
  timeslot_id: null,
  booking_date: ''
})

const disabledTimeslot = computed(() => !form.value.booking_date)

const disabledCourtSelect = computed(() => !form.value.timeslot_id)

const minDate = computed(() => {
  const today = new Date()
  return today.toISOString().split('T')[0]
})

const formatPrice = (price) => {
  return new Intl.NumberFormat('id-ID').format(price)
}

const selectedCourtLabel = computed(() => {
  const c = courts.value.find(c => (c.id ?? c.ID) == form.value.court_id)
  return c ? `${c.court_name}` : '-'
})

const selectedTimeslotLabel = computed(() => {
  const t = displayTimeslots.value.find(t => (t.id ?? t.ID) == form.value.timeslot_id)
  return t ? `${t.display_start} - ${t.display_end}` : '-'
})

const dateOnly = (value) => {
  if (!value) return ''
  return value.includes('T') ? value.split('T')[0] : value.split(' ')[0]
}

const filteredTimeslots = computed(() => {
  if (!form.value.booking_date) return []
  return timeslots.value.filter(ts => {
    const startDate = dateOnly(ts.start_time)
    const endDate = dateOnly(ts.end_time)
    return startDate === form.value.booking_date && endDate === startDate
  })
})

const toHHMM = (value) => {
  if (!value) return ''
  const timePart = value.includes('T')
    ? value.split('T')[1]
    : (value.split(' ')[1] || value)
  const [hh = '', mm = ''] = timePart.split(':')
  return `${hh}:${mm}`
}

const displayTimeslots = computed(() => {
  return filteredTimeslots.value.map(ts => ({
    ...ts,
    display_start: toHHMM(ts.start_time),
    display_end: toHHMM(ts.end_time),
  }))
})

const loadCourts = async () => {
  try {
    const response = await courtAPI.getAll()
    courts.value = response.data.data
  } catch (error) {
    alert('Gagal memuat courts')
  }
}

const loadTimeslots = async () => {
  try {
    const response = await timeslotAPI.getAll()
    timeslots.value = response.data.data
    form.value.timeslot_id = null
    estimatedPrice.value = 0
  } catch (error) {
    alert('Gagal memuat timeslots')
  }
}

const loadBookedCourts = async () => {
  if (!form.value.booking_date || !form.value.timeslot_id) {
    bookedCourtIds.value = new Set()
    return
  }
  try {
    const response = await bookingAPI.getAll()
    const all = response.data?.data || []
    const norm = (s) => String(s ?? '').toLowerCase().trim()
    const isActive = (s) => ['pending', 'confirmed', 'paid'].includes(norm(s))
    const list = all.filter(b => isActive(b.status ?? b.Status))
    const wantedDate = dateOnly(form.value.booking_date)
    const wantedTimeslotId = Number(form.value.timeslot_id)

    const booked = new Set(
      list
        .filter(b => {
          const bDate = dateOnly(b.booking_date)
          const bTimeslotId = Number(b.timeslot_id ?? b.TimeslotID ?? b.timeslot?.id ?? b.timeslot?.ID)
          return bDate === wantedDate && bTimeslotId === wantedTimeslotId
        })
        .map(b => Number(b.court_id ?? b.CourtID ?? b.court?.id ?? b.court?.ID))
    )
    bookedCourtIds.value = booked
  } catch (e) {
    console.warn('Gagal memuat bookings untuk menentukan court yang sudah dibooking')
    bookedCourtIds.value = new Set()
  }
}

const isCourtBooked = (court) => {
  const id = Number(court.id ?? court.ID)
  return bookedCourtIds.value.has(id)
}

const loadBookedTimeslots = async () => {
  if (!form.value.booking_date) {
    bookedTimeslotCourts.value = new Map()
    return
  }
  try {
    const response = await bookingAPI.getAll()
    const all = response.data?.data || []
    const norm = (s) => String(s ?? '').toLowerCase().trim()
    const isActive = (s) => ['pending', 'confirmed', 'paid'].includes(norm(s))
    const list = all.filter(b => isActive(b.status ?? b.Status))
    const targetDate = form.value.booking_date
    const map = new Map()
    list
      .filter(b => dateOnly(b.booking_date) === targetDate)
      .forEach(b => {
        const tId = Number(b.timeslot_id ?? b.TimeslotID ?? b.timeslot?.id ?? b.timeslot?.ID)
        const cId = Number(b.court_id ?? b.CourtID ?? b.court?.id ?? b.court?.ID)
        if (!map.has(tId)) map.set(tId, new Set())
        map.get(tId).add(cId)
      })
    bookedTimeslotCourts.value = map
  } catch (e) {
    console.warn('Gagal memuat bookings untuk menentukan court tiap timeslot')
    bookedTimeslotCourts.value = new Map()
  }
}

const isTimeslotBooked = (timeslot) => {
  const id = Number(timeslot.id ?? timeslot.ID)
  const bookedCourtsSet = bookedTimeslotCourts.value.get(id)
  if (!bookedCourtsSet) return false
  return bookedCourtsSet.size >= courts.value.length && courts.value.length > 0
}

const formatDateLocal = (bookingDate) => {
  if (bookingDate) {
    const [year, month, day] = bookingDate.split('-')
    const joinDate = day+"/"+month+"/"+year
    return joinDate
  }
  return null
}
const handleDateChange = () => {
  form.value.timeslot_id = null
  form.value.court_id = null
  estimatedPrice.value = 0
  isWeekend.value = false
  loadBookedTimeslots()
}

const handleTimeslotChange = () => {
  calculatePrice()
  loadBookedCourts()
}

const calculatePrice = () => {
  if (!form.value.timeslot_id || !form.value.booking_date) {
    estimatedPrice.value = 0
    return
  }

  const selectedTimeslot = filteredTimeslots.value.find(t => (t.id ?? t.ID) == form.value.timeslot_id)
  if (!selectedTimeslot) return

  const date = new Date(form.value.booking_date)
  const dayOfWeek = date.getDay()
  isWeekend.value = dayOfWeek === 0 || dayOfWeek === 6

  estimatedPrice.value = isWeekend.value 
    ? selectedTimeslot.price_weekend 
    : selectedTimeslot.price_weekday

  loadBookedCourts()
}

const createBooking = async () => {
  try {
    loading.value = true
    await bookingAPI.create({
      court_id: parseInt(form.value.court_id),
      timeslot_id: parseInt(form.value.timeslot_id),
      booking_date: form.value.booking_date
    })

    alert('Booking berhasil dibuat!')
    router.push('/customer/my-bookings')
  } catch (error) {
    alert(error.response?.data?.message || 'Gagal membuat booking')
  } finally {
    loading.value = false
  }
}

const resetForm = () => {
  form.value = {
    court_id: null,
    timeslot_id: null,
    booking_date: ''
  }
  estimatedPrice.value = 0
  isWeekend.value = false
  bookedCourtIds.value = new Set()
  bookedTimeslotCourts.value = new Map()
}

onMounted(() => {
  loadCourts()
  loadTimeslots()
})
</script>
