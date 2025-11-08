<template>
  <div>
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-3 mb-6">
      <div>
        <h2 class="text-2xl font-bold">Manage Court</h2>
        <p class="text-sm text-gray-500">Tambah, ubah, atau hapus lapangan</p>
      </div>
      <button
        @click="showCreateModal = true"
        class="inline-flex items-center gap-2 bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 shadow-sm transition"
      >
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-5 h-5"><path d="M12 4.5a.75.75 0 01.75.75V11h5.75a.75.75 0 010 1.5H12.75v5.75a.75.75 0 01-1.5 0V12.5H5.5a.75.75 0 010-1.5h5.75V5.25A.75.75 0 0112 4.5z"/></svg>
        Add Court
      </button>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="text-center py-12 text-gray-500">
      <svg class="animate-spin h-6 w-6 mx-auto mb-2 text-gray-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v4a4 4 0 00-4 4H4z"></path></svg>
      Loading courts...
    </div>

    <!-- Courts Table -->
    <div v-else class="bg-white rounded-xl shadow overflow-hidden">
      <table class="min-w-full text-sm">
        <thead class="bg-gray-50 text-gray-600">
          <tr>
            <th class="px-6 py-3 text-left font-semibold">Nama Court</th>
            <th class="px-6 py-3 text-left font-semibold">Tipe</th>
            <th class="px-6 py-3 text-left font-semibold">Lokasi</th>
            <th class="px-6 py-3 text-left font-semibold">Aksi</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-100">
          <tr v-for="court in courts" :key="court.ID" class="hover:bg-gray-50">
            <td class="px-6 py-4">{{ court.court_name }}</td>
            <td class="px-6 py-4">{{ court.type }}</td>
            <td class="px-6 py-4">{{ court.location }}</td>
            <td class="px-6 py-4">
              <div class="flex items-center gap-3">
                <button
                  @click="editCourt(court)"
                  class="inline-flex items-center gap-1 text-yellow-100 hover:text-yellow-600 transition"
                >
                  <span class="bg-yellow-400 px-2 py-1 hover:bg-yellow-200 rounded">Edit</span>
                </button>
                <button
                  @click="deleteCourt(court.ID)"
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
    <div v-if="showCreateModal || editingCourt" class="fixed inset-0 z-50 flex items-center justify-center">
      <div class="absolute inset-0 bg-black/50"></div>
      <div class="relative bg-white p-6 rounded-xl w-full max-w-md shadow-2xl">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-xl font-bold">
            {{ editingCourt ? 'Edit Court' : 'Add New Court' }}
          </h3>
          <button @click="closeModal" class="w-8 h-8 flex items-center justify-center rounded hover:bg-red-400 hover:text-white transition">✕</button>
        </div>
        
        <form @submit.prevent="saveCourt" class="space-y-4">
          <div>
            <label class="block text-gray-700 mb-1">Nama Court</label>
            <input
              v-model="form.court_name"
              type="text"
              class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              required
            />
          </div>

          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div>
              <label class="block text-gray-700 mb-1">Tipe</label>
              <select v-model="form.type" class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" required>
                <option value="Indoor">Indoor</option>
                <option value="Outdoor">Outdoor</option>
              </select>
            </div>

            <div>
              <label class="block text-gray-700 mb-1">Location</label>
              <input
                v-model="form.location"
                type="text"
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
import { courtAPI } from '@/services/api'

const courts = ref([])
const loading = ref(false)
const showCreateModal = ref(false)
const editingCourt = ref(null)
const form = ref({
  court_name: '',
  type: 'Indoor',
  location: ''
})

const loadCourts = async () => {
  try {
    loading.value = true
    const response = await courtAPI.getAll()
    courts.value = response.data.data
  } catch (error) {
    alert('Gagal memuat court')
  } finally {
    loading.value = false
  }
}

const saveCourt = async () => {
  try {
    if (editingCourt.value) {
      await courtAPI.update(editingCourt.value.ID, form.value)
    } else {
      await courtAPI.create(form.value)
    }
    
    await loadCourts()
    closeModal()
    alert('Court berhasil disimpan!')
  } catch (error) {
    alert(error.response?.data?.message || 'Gagal menyimpan court')
  }
}

const editCourt = (court) => {
  editingCourt.value = court
  form.value = {
    court_name: court.court_name,
    type: court.type,
    location: court.location
  }
}

const deleteCourt = async (id) => {
  if (!confirm('Apakah Anda yakin ingin menghapus court ini?')) return
  
  try {
    await courtAPI.delete(id)
    await loadCourts()
    alert('Court berhasil dihapus!')
  } catch (error) {
    alert(error.response?.data?.message || 'Gagal menghapus court')
  }
}

const closeModal = () => {
  showCreateModal.value = false
  editingCourt.value = null
  form.value = {
    court_name: '',
    type: 'Indoor',
    location: ''
  }
}

onMounted(() => {
  loadCourts()
})
</script>
