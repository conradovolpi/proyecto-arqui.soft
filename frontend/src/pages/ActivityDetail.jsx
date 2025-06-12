// src/pages/ActivityDetail.jsx
import { useParams } from 'react-router-dom';
import { useState, useEffect } from 'react';
import {
  getActivity,
  enrollInActivity,
  getCurrentUser
} from '../services/api';
import CommentForm from '../components/CommentForm';

export default function ActivityDetail() {
  const { id } = useParams();
  const [activity, setActivity] = useState(null);
  const [message, setMessage] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(true);
  const user = getCurrentUser();

  const fetchActivity = async () => {
    try {
      setLoading(true);
      setError('');
      const foundActivity = await getActivity(id);
      setActivity(foundActivity);
    } catch (err) {
      console.error("Error al obtener la actividad:", err);
      setError(err.message || 'Error al cargar la actividad.');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchActivity();
  }, [id]);

  const handleInscription = async () => {
    if (!user) {
      setError('Debes iniciar sesión para inscribirte.');
      return;
    }
    try {
      setMessage('');
      setError('');
      await enrollInActivity(id);
      setMessage('¡Inscripción exitosa!');
      // Opcional: Actualizar la actividad para reflejar el nuevo cupo
      fetchActivity(); 
    } catch (err) {
      console.error("Error al inscribirse:", err);
      setError(err.message || 'Error al inscribirse en la actividad.');
    }
  };

  if (loading) {
    return <p>Cargando actividad...</p>;
  }

  if (error) {
    return <p style={{ color: 'red' }}>Error: {error}</p>;
  }

  if (!activity) return <p>Actividad no encontrada.</p>;

  return (
    <div>
      <h2>{activity.titulo}</h2>
      <p><strong>Descripción:</strong> {activity.descripcion}</p>
      <p><strong>Profesor:</strong> {activity.instructor}</p>
      <p><strong>Horario:</strong> {new Date(activity.horario_inicio).toLocaleTimeString('es-ES', { hour: '2-digit', minute: '2-digit' })} - {new Date(activity.horario_fin).toLocaleTimeString('es-ES', { hour: '2-digit', minute: '2-digit' })}</p>
      <p><strong>Categoría:</strong> {activity.categoria}</p>
      <p><strong>Cupo:</strong> {activity.cupo}</p>

      <button onClick={handleInscription}>Inscribirme</button>
      {message && <p style={{ color: 'green' }}>{message}</p>}
      {error && <p style={{ color: 'red' }}>{error}</p>}

      <hr />
      <h3>Comentarios</h3>
      {activity.comments && activity.comments.length > 0 ? (
        activity.comments.map((c, i) => (
          <div key={i} style={{ borderTop: '1px solid #ccc', margin: '5px 0' }}>
            <p><strong>⭐ {c.rating}</strong> — {c.text}</p>
          </div>
        ))
      ) : (
        <p>No hay comentarios aún.</p>
      )}

      <CommentForm activityId={activity.actividad_id} onCommentAdded={fetchActivity} />
    </div>
  );
}
