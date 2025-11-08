<template>
  <div>
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-3 mb-6">
      <div>
        <h2 class="text-2xl font-bold">Manage Timeslot</h2>
        <p class="text-sm text-gray-500">Kelola jam bermain dan harga</p>
      </div>
      <button
        @click="showCreateModal = true"
        class="inline-flex items-center gap-2 bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 shadow-sm"
      >
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-5 h-5"><path d="M12 4.5a.75.75 0 01.75.75V11h5.75a.75.75 0 010 1.5H12.75v5.75a.75.75 0 01-1.5 0V12.5H5.5a.75.75 0 010-1.5h5.75V5.25A.75.75 0 0112 4.5z"/></svg>
        Add Timeslot
      </button>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="text-center py-12 text-gray-500">
      <svg class="animate-spin h-6 w-6 mx-auto mb-2 text-gray-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v4a4 4 0 00-4 4H4z"></path></svg>
      Loading timeslots...
    </div>

    <!-- Timeslots Table -->
    <div v-else class="bg-white rounded-xl shadow overflow-hidden">
      <table class="min-w-full text-sm">
        <thead class="bg-gray-50 text-gray-600">
          <tr>
            <th class="px-6 py-3 text-left font-semibold">Timeslot</th>
            <th class="px-6 py-3 text-left font-semibold">Harga Weekday</th>
            <th class="px-6 py-3 text-left font-semibold">Harga Weekend</th>
            <th class="px-6 py-3 text-left font-semibold">Aksi</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-100">
          <tr v-for="timeslot in timeslots" :key="timeslot.ID" class="hover:bg-gray-50">
            <td class="px-6 py-4">Tanggal : {{ timeslot.base_date }} 
              <br>Waktu : {{ timeslot.start_time }} - {{ timeslot.end_time }}</td>
            <td class="px-6 py-4">Rp {{ formatPrice(timeslot.price_weekday) }}</td>
            <td class="px-6 py-4">Rp {{ formatPrice(timeslot.price_weekend) }}</td>
            <td class="px-6 py-4">
              <div class="flex items-center gap-3">
                <button
                  @click="editTimeslot(timeslot)"
                  class="inline-flex items-center gap-1 text-yellow-100 hover:text-yellow-600 transition"
                >
                  <span class="bg-yellow-400 px-2 py-1 hover:bg-yellow-200 rounded">Edit</span>
                </button>
                <button
                  @click="deleteTimeslot(timeslot.ID)"
                  class="inline-flex items-center gap-1 text-red-100 hover:text-red-600 transition"
                >
                  <span class="bg-red-400 px-2 py-1 hover:bg-red-200 rounded">Delete</span>
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Create/Edit Modal -->
    <div v-if="showCreateModal || editingTimeslot" class="fixed inset-0 z-50 flex items-center justify-center">
      <div class="absolute inset-0 bg-black/50"></div>
      <div class="relative bg-white p-6 rounded-xl w-full max-w-md shadow-2xl">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-xl font-bold">
            {{ editingTimeslot ? 'Edit Timeslot' : 'Add New Timeslot' }}
          </h3>
          <button @click="closeModal" class="w-8 h-8 flex items-center justify-center rounded hover:bg-red-400 hover:text-white transition">✕</button>
        </div>
        
        <form @submit.prevent="saveTimeslot" class="space-y-4">
          <div>
            <label class="block text-gray-700 mb-1">Tanggal (DD-MM-YYYY)</label>
            <input
              v-model="form.base_date"
              type="date"
              class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              required
            />
          </div>

          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div>
              <label class="block text-gray-700 mb-1">Start Time</label>
              <input
                v-model="form.start_time"
                type="time"
                step="1"
                class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                required
              />
            </div>
            <div>
              <label class="block text-gray-700 mb-1">End Time</label>
              <input
                v-model="form.end_time"
                type="time"
                step="1"
                class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                required
              />
            </div>
          </div>

          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div>
              <label class="block text-gray-700 mb-1">Harga Weekday</label>
              <input
                v-model.number="form.price_weekday"
                type="number"
                class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                required
              />
            </div>
            <div>
              <label class="block text-gray-700 mb-1">Harga Weekend</label>
              <input
                v-model.number="form.price_weekend"
                type="number"
                class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                required
              />
            </div>
          </div>

          <div class="flex justify-end gap-2 pt-2">
            <button
              type="button"
              @click="closeModal"
              class="px-4 py-2 border rounded-lg hover:bg-gray-50"
            >
              Cancel
            </button>
            <button
              type="submit"
              class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
            >
              Save
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { timeslotAPI } from '@/services/api'

const timeslots = ref([])
const loading = ref(false)
const showCreateModal = ref(false)
const editingTimeslot = ref(null)
const todayISO = () => new Date().toISOString().split('T')[0]
const form = ref({
  base_date: todayISO(),
  start_time: '',
  end_time: '',
  price_weekday: 0,
  price_weekend: 0
})

const formatPrice = (price) => {
  return new Intl.NumberFormat('id-ID').format(price)
}

const toSeconds = (t) => {
  if (!t) return null
  const parts = t.split(':').map((n) => parseInt(n, 10))
  if (parts.length < 2 || parts.length > 3 || parts.some(Number.isNaN)) return null
  const [h, m, s = 0] = parts
  if (h < 0 || h > 23 || m < 0 || m > 59 || s < 0 || s > 59) return null
  return h * 3600 + m * 60 + s
}

const loadTimeslots = async () => {
  try {
    loading.value = true
    const response = await timeslotAPI.getAll()
    timeslots.value = response.data.data
    for (const timeslt in timeslots.value) {
      const element = timeslots.value[timeslt];
      element.base_date = new Date(element.start_time).toLocaleDateString('id-ID')
      element.start_time = new Date(element.start_time).toTimeString().split(' ')[0]
      element.end_time = new Date(element.end_time).toTimeString().split(' ')[0]
      const start_split = element.start_time.split(':')
      const start_join = start_split[0]+":"+start_split[1]
      element.start_time = start_join
      const end_split = element.end_time.split(':')
      const end_join = end_split[0]+":"+end_split[1]
      element.end_time = end_join
    }
  } catch (error) {
    alert('Gagal memuat timeslots')
  } finally {
    loading.value = false
  }
}

const saveTimeslot = async () => {
  try {
    const startSec = toSeconds(form.value.start_time)
    const endSec = toSeconds(form.value.end_time)
    if (startSec == null || endSec == null) {
      alert('Waktu tidak valid. Gunakan format HH:MM atau HH:MM:SS')
      return
    }
    if (endSec <= startSec) {
      alert('End Time harus lebih besar dari Start Time')
      return
    }
    if (form.value.price_weekday <= 0 || form.value.price_weekend <= 0) {
      alert('Harga Weekday/Weekend harus lebih dari 0')
      return
    }

    const normalizeTime = (t) => {
      if (!t) return ''
      return /^\d{2}:\d{2}:\d{2}$/.test(t) ? t : `${t}:00`
    }
    const normalizeDate = (d) => {
      return d
    }

    const payload = {
      start_time: `${normalizeDate(form.value.base_date)} ${normalizeTime(form.value.start_time)}`,
      end_time: `${normalizeDate(form.value.base_date)} ${normalizeTime(form.value.end_time)}`,
      price_weekday: form.value.price_weekday,
      price_weekend: form.value.price_weekend,
    }
  
    if (editingTimeslot.value) {
      await timeslotAPI.update(editingTimeslot.value.ID, payload)
    } else {
      await timeslotAPI.create(payload)
    }
    
    await loadTimeslots()
    closeModal()
    alert('Timeslot berhasil disimpan!')
  } catch (error) {
    alert(error.response?.data?.message || 'Gagal menyimpan timeslot')
  }
}

const editTimeslot = (timeslot) => {
  editingTimeslot.value = timeslot
  var [day, month, year] = timeslot.base_date.split('/')
  day = String(day).padStart(2, '0');
  month = String(month).padStart(2, '0');
  const dateObj = year+"-"+month+"-"+day

  form.value = {
    base_date: dateObj,
    start_time: timeslot.start_time,
    end_time: timeslot.end_time,
    price_weekday: timeslot.price_weekday,
    price_weekend: timeslot.price_weekend
  }
}

const deleteTimeslot = async (id) => {
  if (!confirm('Apakah Anda yakin ingin menghapus timeslot ini?')) return
  
  try {
    await timeslotAPI.delete(id)
    await loadTimeslots()
    alert('Timeslot berhasil dihapus!')
  } catch (error) {
    alert(error.response?.data?.message || 'Gagal menghapus timeslot')
  }
}

const closeModal = () => {
  showCreateModal.value = false
  editingTimeslot.value = null
  form.value = {
    base_date: todayISO(),
    start_time: '',
    end_time: '',
    price_weekday: 0,
    price_weekend: 0
  }
}

onMounted(() => {
  loadTimeslots()
})
</script>
