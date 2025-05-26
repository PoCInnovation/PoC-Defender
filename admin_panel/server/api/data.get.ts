export default defineEventHandler(() => {
  return {
    message: 'Here is some protected proxy data.',
    connections: 42,
    alerts: 3,
    trafficMB: 120.5
  };
});
