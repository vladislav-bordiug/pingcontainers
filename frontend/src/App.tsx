import React, { useEffect, useState } from 'react';
import { PingStatus } from './models/PingStatus';
import { fetchPingStatuses } from './services/api';
import { intervaltime } from './utils/parseInterval'
import { Navbar, Table } from 'react-bootstrap';
import "bootswatch/dist/cosmo/bootstrap.min.css";

const App: React.FC = () => {
  const [statuses, setStatuses] = useState<PingStatus[]>([]);
  
  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await fetchPingStatuses();
        setStatuses(data);
      } catch (error) {
        console.error('Ошибка получения данных:', error);
      }
    };

    fetchData();
    const interval = setInterval(fetchData, intervaltime);
    return () => clearInterval(interval);
  }, []);

  return (
    <>
      <Navbar className="d-flex align-items-center justify-content-center">
        <Navbar.Brand className="d-flex align-items-center justify-content-center">
          Статус пинга контейнеров
        </Navbar.Brand>
      </Navbar>
      <Table striped bordered hover style={{ tableLayout: 'fixed', width: '100%' }}>
        <thead>
          <tr className="text-center">
            <th style={{ width: '30%' }}>IP-адрес</th>
            <th style={{ width: '30%' }}>Время пинга (ms)</th>
            <th style={{ width: '30%' }}>Дата последней успешной попытки</th>
          </tr>
        </thead>
        <tbody>
          {statuses.map((status) => (
            <tr key={status.ip}>
              <td>{status.ip}</td>
              <td>{status.ping_time}</td>
              <td>{new Date(status.last_success).toLocaleString()}</td>
            </tr>
          ))}
        </tbody>
      </Table>
    </>
  );
};

export default App;