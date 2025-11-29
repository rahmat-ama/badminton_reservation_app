// Mengambil nilai dari Environment Variable (yang diset di Dockerfile)
// Jika tidak ada (misal saat dev lokal tanpa docker), baru pakai localhost
export const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8000/api'

export const API_ENDPOINTS = {
  LOGIN: '/auth/login',
  REGISTER: '/auth/register',
  
  USERS: '/user',
  USER_BY_ID: (id) => `/user/${id}`,
  
  COURTS: '/court',
  COURT_BY_ID: (id) => `/court/${id}`,
  
  TIMESLOTS: '/timeslot',
  TIMESLOT_BY_ID: (id) => `/timeslot/${id}`,
  
  BOOKINGS: '/booking',
  BOOKING_BY_ID: (id) => `/booking/${id}`,
}
