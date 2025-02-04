import { PingStatus } from '../models/PingStatus';

const backendUrl = import.meta.env.VITE_BACKEND_URL || 'http://localhost:8080';

export const fetchPingStatuses = async (): Promise<PingStatus[]> => {
  const res = await fetch(`${backendUrl}/api/status`);
  if (!res.ok) {
    throw new Error('Ошибка при получении данных с сервера');
  }
  return res.json();
};