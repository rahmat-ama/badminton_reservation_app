import axios from './axios'
import { API_ENDPOINTS } from '@/config/api'

export const authAPI = {
  login: (credentials) => axios.post(API_ENDPOINTS.LOGIN, credentials),
  register: (userData) => axios.post(API_ENDPOINTS.REGISTER, userData),
}

export const userAPI = {
  getAll: () => axios.get(API_ENDPOINTS.USERS),
  getById: (id) => axios.get(API_ENDPOINTS.USER_BY_ID(id)),
  update: (id, data) => axios.put(API_ENDPOINTS.USER_BY_ID(id), data),
  delete: (id) => axios.delete(API_ENDPOINTS.USER_BY_ID(id)),
}

export const courtAPI = {
  getAll: () => axios.get(API_ENDPOINTS.COURTS),
  getById: (id) => axios.get(API_ENDPOINTS.COURT_BY_ID(id)),
  create: (data) => axios.post(API_ENDPOINTS.COURTS, data),
  update: (id, data) => axios.put(API_ENDPOINTS.COURT_BY_ID(id), data),
  delete: (id) => axios.delete(API_ENDPOINTS.COURT_BY_ID(id)),
}

export const timeslotAPI = {
  getAll: () => axios.get(API_ENDPOINTS.TIMESLOTS),
  getById: (id) => axios.get(API_ENDPOINTS.TIMESLOT_BY_ID(id)),
  create: (data) => axios.post(API_ENDPOINTS.TIMESLOTS, data),
  update: (id, data) => axios.put(API_ENDPOINTS.TIMESLOT_BY_ID(id), data),
  delete: (id) => axios.delete(API_ENDPOINTS.TIMESLOT_BY_ID(id)),
}

export const bookingAPI = {
  getAll: () => axios.get(API_ENDPOINTS.BOOKINGS),
  getById: (id) => axios.get(API_ENDPOINTS.BOOKING_BY_ID(id)),
  create: (data) => axios.post(API_ENDPOINTS.BOOKINGS, data),
  update: (id, data) => axios.put(API_ENDPOINTS.BOOKING_BY_ID(id), data),
  delete: (id) => axios.delete(API_ENDPOINTS.BOOKING_BY_ID(id)),
}