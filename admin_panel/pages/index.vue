<template>
  <div class="min-h-screen bg-blue-50 p-6">
    <div class="max-w-7xl mx-auto space-y-8">
      <header class="bg-gradient-to-r from-blue-600 to-indigo-700 rounded-xl shadow-lg p-6 flex flex-col md:flex-row items-center justify-between gap-4">
        <div class="flex items-center gap-3">
          <div class="bg-white p-3 rounded-full shadow">
            <span class="text-blue-600 text-2xl">🛡️</span>
          </div>
          <h1 class="text-3xl font-bold text-white">Proxy Admin Dashboard</h1>
        </div>
        <div class="flex gap-4">
          <button class="px-4 py-2 bg-white text-blue-600 font-medium rounded-lg hover:bg-blue-50 transition flex items-center gap-2 shadow">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
            Refresh Data
          </button>
          <button class="px-4 py-2 bg-blue-800 text-white font-medium rounded-lg hover:bg-blue-900 transition flex items-center gap-2 shadow">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
            Settings
          </button>
        </div>
      </header>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="bg-white p-6 rounded-xl shadow hover:shadow-md transition">
          <div class="flex justify-between items-start">
            <div>
              <p class="text-gray-500 text-sm font-medium">Total Traffic</p>
              <h2 class="text-3xl font-bold text-gray-800 mt-1">1.4 TB</h2>
            </div>
            <div class="bg-blue-100 p-3 rounded-lg">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
              </svg>
            </div>
          </div>
          <p class="text-green-600 text-sm font-medium mt-4 flex items-center">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 10l7-7m0 0l7 7m-7-7v18" />
            </svg>
            8.2% increase
          </p>
        </div>

        <div class="bg-white p-6 rounded-xl shadow hover:shadow-md transition">
          <div class="flex justify-between items-start">
            <div>
              <p class="text-gray-500 text-sm font-medium">Active Connections</p>
              <h2 class="text-3xl font-bold text-gray-800 mt-1">248</h2>
            </div>
            <div class="bg-blue-100 p-3 rounded-lg">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24" stroke="currentColor" class="text-blue-600">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
              </svg>
            </div>
          </div>
          <p class="text-red-600 text-sm font-medium mt-4 flex items-center">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="none" viewBox="0 0 24 24" stroke="currentColor" class="mr-1">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3" />
            </svg>
            3.1% decrease
          </p>
        </div>

        <div class="bg-white p-6 rounded-xl shadow hover:shadow-md transition">
          <div class="flex justify-between items-start">
            <div>
              <p class="text-gray-500 text-sm font-medium">Detected Threats</p>
              <h2 class="text-3xl font-bold text-gray-800 mt-1">7</h2>
            </div>
            <div class="bg-blue-100 p-3 rounded-lg">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24" stroke="currentColor" class="text-blue-600">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
              </svg>
            </div>
          </div>
          <p class="text-yellow-600 text-sm font-medium mt-4 flex items-center">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="none" viewBox="0 0 24 24" stroke="currentColor" class="mr-1">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            2 high severity
          </p>
        </div>
      </div>

      <section class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="bg-white p-6 rounded-xl shadow">
          <div class="flex justify-between items-center mb-6">
            <h2 class="text-xl font-bold text-gray-800">Traffic Analysis</h2>
            <div class="flex gap-2">
              <button class="px-3 py-1 bg-blue-100 text-blue-600 text-sm rounded-lg hover:bg-blue-200 transition">Day</button>
              <button class="px-3 py-1 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 transition">Week</button>
              <button class="px-3 py-1 bg-blue-100 text-blue-600 text-sm rounded-lg hover:bg-blue-200 transition">Month</button>
            </div>
          </div>
          <StatsChart :data="statsData" />
        </div>
        
        <div class="bg-white p-6 rounded-xl shadow">
          <div class="flex justify-between items-center mb-6">
            <h2 class="text-xl font-bold text-gray-800">Security Alerts</h2>
            <span class="px-3 py-1 bg-red-100 text-red-600 text-sm font-medium rounded-full">
              {{ alerts.length }} active
            </span>
          </div>
          <AttackAlerts :alerts="alerts" />
        </div>
      </section>
      
      <section class="bg-white p-6 rounded-xl shadow">
        <div class="flex justify-between items-center mb-6">
          <h2 class="text-xl font-bold text-gray-800">Connection Overview</h2>
          <div class="flex gap-2">
            <div class="relative">
              <input type="text" placeholder="Search connections..." class="pl-9 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 outline-none">
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="none" viewBox="0 0 24 24" stroke="currentColor" class="text-gray-400 absolute left-3 top-2.5">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </div>
            <button class="px-4 py-2 bg-blue-100 text-blue-600 font-medium rounded-lg hover:bg-blue-200 transition flex items-center gap-2">
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.293A1 1 0 013 6.586V4z" />
              </svg>
              Filter
            </button>
          </div>
        </div>
        <ConnectionList :connections="connections" />
      </section>
    </div>

    <footer class="mt-12 text-center text-gray-500 text-sm">
      <p>Proxy Admin Dashboard © {{ new Date().getFullYear() }} | Server Uptime: 99.98%</p>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useProxyApi } from '~/composables/useProxyApi'
import StatsChart from '~/components/StatsChart.vue'
import AttackAlerts from '~/components/AttackAlerts.vue'
import ConnectionList from '~/components/ConnectionList.vue'

const { getConnections, getAlerts, getStats } = useProxyApi()
const connections = ref([])
const alerts = ref([])
const statsData = ref({})

onMounted(async () => {
  try {
    connections.value = await getConnections()
    alerts.value = await getAlerts()
    statsData.value = await getStats()
  } catch (error) {
    console.error('Failed to fetch dashboard data:', error)
    // Could add toast notification here
  }
})
</script>

<style scoped>
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

::-webkit-scrollbar {
  width: 8px;
}

::-webkit-scrollbar-track {
  background: #f1f5f9;
  border-radius: 8px;
}

::-webkit-scrollbar-thumb {
  background: #94a3b8;
  border-radius: 8px;
}

::-webkit-scrollbar-thumb:hover {
  background: #64748b;
}
</style>
