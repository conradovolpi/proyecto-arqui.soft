import { useParams } from 'react-router-dom';
import { useEffect, useState } from 'react';
import api from './api';

export default function ActivityDetail() {
  const { id } = useParams();
  const [activity, setActivity] = useState(null);
  const [msg, setMsg] = useState('');

  useEffect(() => {
    api.get(`/activities/${id}`)
      .then(res => setActivity(res.data))
      .catch(err => console.error(err));
  }, [id]);

  const inscribirse = () => {
    api.post(`/activities/${id}/inscribir`)
      .then(() => setMsg('✅ Inscripción exitosa'))
      .catch(() => setMsg('❌ Error al inscribirse'));
  };

  if (!activity) return <p>Cargando...</p>;

  return (
    <div>
      <h2>{activity.titulo}</h2>
      <p>Instructor: {activity.profesor}</p>
      <p>Horario: {activity.horario}</p>
      <button onClick={inscribirse}>Inscribirme</button>
      <p>{msg}</p>
    </div>
  );
}
