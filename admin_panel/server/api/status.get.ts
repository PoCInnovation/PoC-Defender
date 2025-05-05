export default defineEventHandler(() => ({
  status: 'OK',
  uptime: process.uptime(),
  timestamp: new Date().toISOString()
}))
