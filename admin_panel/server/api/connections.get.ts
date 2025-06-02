export default defineEventHandler(() => {
  return [
    { id: 1, ip: '192.168.0.10', endpoint: '/login', duration: 12 },
    { id: 2, ip: '10.0.0.5', endpoint: '/api/data', duration: 37 }
  ];
});
