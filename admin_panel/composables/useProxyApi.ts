import { useRuntimeConfig } from '#imports'
import axios from 'axios'

export function useProxyApi() {
  const config = useRuntimeConfig()
  const client = axios.create({ baseURL: config.public.apiBase })

  const getConnections = async () => {
    const res = await client.get('/api/connections')
    return res.data
  }

  const getAlerts = async () => {
    const res = await client.get('/api/alerts')
    return res.data
  }

  const getStats = async () => {
    const res = await client.get('/api/stats')
    return res.data
  }

  return { getConnections, getAlerts, getStats }
}
