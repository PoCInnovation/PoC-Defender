export default defineEventHandler(() => {
  return {
    timestamps: ['10:00', '10:01', '10:02', '10:03'],
    requests: [30, 42, 51, 37]
  };
});
