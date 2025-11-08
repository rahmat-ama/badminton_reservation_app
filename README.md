Badminton Reservation App
=========================

A full-stack reservation system for badminton courts built with Go (Gin) for the backend, Vue 3 + Vite for the frontend, Tailwind CSS for styling, and PostgreSQL for data. Docker Compose is provided to run the entire stack locally.

Features
--------
- Authentication (login/register) with JWT
- Role-based access (Admin, Customer)
- Courts management (Admin)
- Timeslots management with weekday/weekend pricing (Admin)
- Booking flow (Customer)
	- Select date → filter timeslots by that date
	- Show timeslots as disabled when fully booked (all courts taken) for selected date
	- Court selection disabled if already booked for the chosen date+timeslot
	- Pricing automatically adjusts based on weekday/weekend
- Profile management (Customer)

Tech stack
---------
- Backend: Go 1.22, Gin, GORM, JWT
- Frontend: Vue 3, Vite, Pinia, Vue Router, Axios, Tailwind CSS
- Database: PostgreSQL 17
- Containerization: Docker, Docker Compose, Nginx (for static frontend)

Repository layout
-----------------
- `backend_go/` – Go API server (Gin) and services
- `frontend_vue/` – Vue 3 SPA with Tailwind CSS
- `docker-compose.yml` – Orchestrates db, backend, frontend
- `pgdata/` – Local volume for Postgres data (created by Compose)

Backend overview
----------------
- Entrypoint: `backend_go/main.go` (or `cmd/server` if present)
- Key folders:
	- `controllers/` – HTTP handlers (booking, court, timeslot, user)
	- `models/` – GORM models
	- `dto/` – Request/response DTOs
	- `services/` – Business logic (auth, user)
	- `middleware/` – JWT auth middleware
	- `routes/` – Route registration
	- `db/` – Database init
	- `seed/` – Initial data seeding
- Expected env vars (example):
	- `DATABASE_HOST` (default: `db` in Docker)
	- `DATABASE_PORT` (default: `5432`)
	- `DATABASE_USER` (default: `postgres`)
	- `DATABASE_PASSWORD` (default: `postgres`)
	- `DATABASE_NAME` (default: `badminton_reservation`)
	- `JWT_SECRET` (required)
	- `APP_PORT` (default: `8000`)

Frontend overview
-----------------
- Entrypoint: `frontend_vue/index.html`; app bootstrapped in `src/main.js`
- Router guards ensure the login page is accessible without redirect loops
- Axios attaches Bearer token from Pinia store
- Tailwind CSS v4 via `@import 'tailwindcss'` in `src/assets/main.css`
- Config: `src/config/api.js` provides base URL and endpoints
- UI/UX:
	- Customer booking flow on `customer/BookingView.vue`
	- Customer dashboard and profile pages
	- Admin: courts, timeslots, bookings

Business logic: booking rules (frontend)
---------------------------------------
1. User selects a date (YYYY-MM-DD)
2. Timeslots are filtered by that date (derived from each timeslot’s start_time and end_time date)
3. A timeslot is disabled only if all courts are already booked on that date and timeslot
4. After a timeslot is chosen, the court dropdown enables and shows all courts; those already booked are disabled
5. Cancelled bookings do not count towards “booked” capacity and free up the court/timeslot
6. Estimated price auto-updates as weekday or weekend rate

Run locally with Docker
-----------------------
Prereqs: Docker Desktop (or Docker Engine) installed.

1) Build & start all services

```powershell
docker compose build
docker compose up -d
```

2) Verify

- Backend API: http://localhost:8000/api
- Frontend: http://localhost:5173
- Postgres: localhost:5430 (user/pass: postgres/postgres) DB: badminton_reservation_db

3) Stop

```powershell
docker compose down
```

Environment configuration
-------------------------
- Backend reads env variables listed above. In Docker Compose, they’re set for the `backend` service.
- Frontend uses `VITE_API_BASE_URL` at build time (Compose sets it to `http://backend:8000/api`). For local non-Docker dev, adjust `src/config/api.js`.

Local development (without Docker)
----------------------------------
Backend:

```powershell
# From backend_go/
go mod tidy
go run .\main.go
```

Frontend:

```powershell
# From frontend_vue/
npm install
npm run dev
```

Useful scripts
--------------
- Frontend build: `npm run build` (produces `frontend_vue/dist`)
- Docker Compose build & up: see above

