// src/pages/Home.jsx
import { useEffect, useState } from 'react';
import { getActivities } from '../services/mockData';
import ActivityCard from '../components/ActivityCard';

export default function Home() {
  const [activities, setActivities] = useState([]);
  const [filtered, setFiltered] = useState([]);
  const [search, setSearch] = useState('');

  useEffect(() => {
    const data = getActivities();
    setActivities(data);
    setFiltered(data);
  }, []);

  const handleSearch = (e) => {
    const value = e.target.value.toLowerCase();
    setSearch(value);
    const results = activities.filter(
      (act) =>
        act.title.toLowerCase().includes(value) ||
        act.category.toLowerCase().includes(value) ||
        act.schedule.includes(value)
    );
    setFiltered(results);
  };

  return (
    <div>
      <h2>Actividades deportivas</h2>

      <input
        type="text"
        placeholder="Buscar por título, categoría u horario"
        value={search}
        onChange={handleSearch}
      />

      {filtered.length > 0 ? (
        <div>
          {filtered.map((act) => (
            <ActivityCard key={act.id} activity={act} />
          ))}
        </div>
      ) : (
        <p>No hay resultados</p>
      )}
    </div>
  );
}
