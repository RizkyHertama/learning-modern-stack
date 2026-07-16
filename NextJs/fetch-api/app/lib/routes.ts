// lib/routes.ts

export const ROUTES = {
  HOME: '/',
  // Route Statis
  EMPLOYEE: '/employee',
  ATTENDANCE: '/attendance',
  SETTINGS: '/settings',

  // Route Dinamis (menggunakan function)
  EMPLOYEE_DETAIL: (id: string | number) => `/employee/${id}`,
  EMPLOYEE_EDIT: (id: string | number) => `/employee/${id}/edit`,
};