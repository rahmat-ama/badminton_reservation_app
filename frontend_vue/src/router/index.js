import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/login'
    },
    {
      path: '/login',
      name: 'Login',
      component: () => import('@/views/LoginView.vue'),
      meta: { requiresGuest: true }
    },
    {
      path: '/register',
      name: 'Register',
      component: () => import('@/views/RegisterView.vue'),
      meta: { requiresGuest: true }
    },
    {
      path: '/admin',
      name: 'AdminDashboard',
      component: () => import('@/views/admin/DashboardView.vue'),
      meta: { requiresAuth: true, requiresAdmin: true },
      children: [
        {
          path: 'courts',
          name: 'AdminCourts',
          component: () => import('@/views/admin/CourtsView.vue'),
        },
        {
          path: 'timeslots',
          name: 'AdminTimeslots',
          component: () => import('@/views/admin/TimeslotsView.vue'),
        },
        {
          path: 'bookings',
          name: 'AdminBookings',
          component: () => import('@/views/admin/BookingsView.vue'),
        },
      ]
    },
    {
      path: '/customer',
      name: 'CustomerDashboard',
      component: () => import('@/views/customer/DashboardView.vue'),
      meta: { requiresAuth: true, requiresCustomer: true },
      children: [
        {
          path: 'booking',
          name: 'CustomerBooking',
          component: () => import('@/views/customer/BookingView.vue'),
        },
        {
          path: 'my-bookings',
          name: 'MyBookings',
          component: () => import('@/views/customer/MyBookingsView.vue'),
        },
        {
          path: 'profile',
          name: 'CustomerProfile',
          component: () => import('@/views/customer/ProfileView.vue'),
        },
      ]
    },
  ],
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next('/login')
    return
  }

  if (to.meta.requiresGuest && authStore.isAuthenticated) {
    if (authStore.isAdmin) {
      next('/admin/courts')
      return
    }
    if (authStore.isCustomer) {
      next('/customer/booking')
      return
    }
  }

  if (to.meta.requiresAdmin && !authStore.isAdmin) {
    if (authStore.isCustomer) {
      next('/customer/booking')
    } else {
      next('/login')
    }
    return
  }

  if (to.meta.requiresCustomer && !authStore.isCustomer) {
    if (authStore.isAdmin) {
      next('/admin/courts')
    } else {
      next('/login')
    }
    return
  }

  next()
})

export default router
