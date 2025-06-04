import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import api from './api';

export default function MyActivities() {
  const [activities, setActivities] = useState([]);

  useEffect(() => {
    api.get('/mis-actividades')
      .then(res => setActivities(res.data))
      .catch(err => console.error(err));
  }, []);

  return (
    <div>
      <h2>Mis actividades</h2>
      <ul>
        {activities.map(act => (
          <li key={act.id}>
            <Link to={`/activity/${act.id}`}>{act.titulo}</Link>
          </li>
        ))}
      </ul>
    </div>
  );
}
