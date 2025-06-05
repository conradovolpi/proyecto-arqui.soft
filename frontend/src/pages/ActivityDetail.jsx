// src/pages/ActivityDetail.jsx
import { useParams } from 'react-router-dom';
import { useState, useEffect } from 'react';
import {
  getActivityById,
  enrollUser,
  getCurrentUser
} from '../services/mockData';
import CommentForm from '../components/CommentForm';

export default function ActivityDetail() {
  const { id } = useParams();
  const [activity, setActivity] = useState(null);
  const [message, setMessage] = useState('');
  const [error, setError] = useState('');
  const user = getCurrentUser();

  useEffect(() => {
    const found = getActivityById(id);
    if (found) {
      setActivity({ ...found });
    }
  }, [id]);

  const handleInscription = () => {
    const result = enrollUser(activity.id, user.id);
    if (result.success) {
      setMessage(result.msg);
      setError('');
    } else {
      setError(result.msg);
      setMessage('');
    }
  };

  const handleRefresh = () => {
    const updated = getActivityById(id);
    setActivity({ ...updated });
  };

  if (!activity) return <p>Cargando...</p>;

  return (
    <div>
      <h2>{activity.title}</h2>
      <p><strong>Descripción:</strong> {activity.description}</p>
      <p><strong>Profesor:</strong> {activity.instructor}</p>
      <p><strong>Horario:</strong> {activity.schedule}</p>
      <p><strong>Duración:</strong> {activity.duration}</p>
      <p><strong>Categoría:</strong> {activity.category}</p>
      <p><strong>Cupo:</strong> {activity.enrolledUsers.length}/{activity.capacity}</p>

      <button onClick={handleInscription}>Inscribirme</button>
      {message && <p style={{ color: 'green' }}>{message}</p>}
      {error && <p style={{ color: 'red' }}>{error}</p>}

      <hr />
      <h3>Comentarios</h3>
      {activity.comments.length > 0 ? (
        activity.comments.map((c, i) => (
          <div key={i} style={{ borderTop: '1px solid #ccc', margin: '5px 0' }}>
            <p><strong>⭐ {c.rating}</strong> — {c.text}</p>
          </div>
        ))
      ) : (
        <p>No hay comentarios aún.</p>
      )}

      <CommentForm activityId={activity.id} onCommentAdded={handleRefresh} />
    </div>
  );
}
