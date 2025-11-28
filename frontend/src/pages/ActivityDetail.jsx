// src/pages/ActivityDetail.jsx
import { useParams, useNavigate } from 'react-router-dom';
import { useState, useEffect } from 'react';
import {
  getActivity,
  enrollInActivity,
  getCurrentUser,
  updateActivity,
  deleteActivity
} from '../services/api';
import { convertScheduleAndDurationToTimeRange } from '../utils/dateUtils';
import CommentForm from '../components/CommentForm';

export default function ActivityDetail() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [activity, setActivity] = useState(null);
  const [message, setMessage] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(true);
  const [isEditMode, setIsEditMode] = useState(false);
  const [successMessage, setSuccessMessage] = useState('');
  const [form, setForm] = useState({
    title: '',
    description: '',
    instructor: '',
    schedule: '',
    duration: '',
    category: '',
    capacity: '',
  });
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

  const handleFormChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleEdit = () => {
    if (!activity) return;
    
    // Calcular duración desde horario_inicio y horario_fin
    const startDate = new Date(activity.horario_inicio);
    const endDate = new Date(activity.horario_fin);
    const durationMs = endDate - startDate;
    const durationHours = Math.floor(durationMs / (1000 * 60 * 60));
    const durationMinutes = Math.floor((durationMs % (1000 * 60 * 60)) / (1000 * 60));
    
    let durationStr = '';
    if (durationHours > 0) {
      durationStr += `${durationHours}h`;
    }
    if (durationMinutes > 0) {
      durationStr += ` ${durationMinutes}m`;
    }
    if (!durationStr) {
      durationStr = '1h'; // Default
    }

    setForm({
      title: activity.titulo,
      description: activity.descripcion,
      instructor: activity.instructor,
      schedule: new Date(activity.horario_inicio).toLocaleTimeString('es-ES', { hour: '2-digit', minute: '2-digit' }),
      duration: durationStr.trim(),
      category: activity.categoria,
      capacity: activity.cupo.toString(),
    });
    setIsEditMode(true);
    setError('');
    setMessage('');
    setSuccessMessage('');
  };

  const handleCancelEdit = () => {
    setIsEditMode(false);
    setForm({
      title: '',
      description: '',
      instructor: '',
      schedule: '',
      duration: '',
      category: '',
      capacity: '',
    });
    setError('');
    setSuccessMessage('');
  };

  const handleUpdate = async (e) => {
    e.preventDefault();
    setError('');
    setSuccessMessage('');

    try {
      const { horario_inicio, horario_fin } = convertScheduleAndDurationToTimeRange(form.schedule, form.duration);

      const dataToSend = {
        titulo: form.title,
        descripcion: form.description,
        instructor: form.instructor,
        horario_inicio: horario_inicio,
        horario_fin: horario_fin,
        cupo: Number(form.capacity),
        categoria: form.category,
      };

      await updateActivity(id, dataToSend);
      setSuccessMessage('Congrats');
      setIsEditMode(false);
      // Recargar la actividad actualizada
      await fetchActivity();
    } catch (err) {
      console.error("Error al actualizar actividad:", err);
      setError(err.message || 'Error al actualizar la actividad.');
    }
  };

  const handleDelete = async () => {
    if (!window.confirm('¿Estás seguro que quieres eliminar esta actividad?')) {
      return;
    }

    setError('');
    setSuccessMessage('');

    try {
      await deleteActivity(id);
      setSuccessMessage('Congrats');
      // Redirigir a Home después de un breve delay para mostrar el mensaje
      setTimeout(() => {
        navigate('/');
      }, 1000);
    } catch (err) {
      console.error("Error al eliminar actividad:", err);
      setError(err.message || 'Error al eliminar la actividad.');
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
      {!isEditMode ? (
        <>
          <h2>{activity.titulo}</h2>
          <p><strong>Descripción:</strong> {activity.descripcion}</p>
          <p><strong>Profesor:</strong> {activity.instructor}</p>
          <p><strong>Horario:</strong> {new Date(activity.horario_inicio).toLocaleTimeString('es-ES', { hour: '2-digit', minute: '2-digit' })} - {new Date(activity.horario_fin).toLocaleTimeString('es-ES', { hour: '2-digit', minute: '2-digit' })}</p>
          <p><strong>Categoría:</strong> {activity.categoria}</p>
          <p><strong>Cupo:</strong> {activity.cupo}</p>

          <div style={{ marginTop: '20px', marginBottom: '20px' }}>
            <button onClick={handleInscription}>Inscribirme</button>
            
            {user && user.rol === 'admin' && (
              <div style={{ marginTop: '10px' }}>
                <button onClick={handleEdit} style={{ marginRight: '10px', backgroundColor: '#4CAF50', color: 'white' }}>
                  Editar
                </button>
                <button onClick={handleDelete} style={{ backgroundColor: '#f44336', color: 'white' }}>
                  Eliminar
                </button>
              </div>
            )}
          </div>

          {successMessage && <p style={{ color: 'green', fontWeight: 'bold' }}>{successMessage}</p>}
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
        </>
      ) : (
        <div>
          <h2>Editar Actividad</h2>
          <form onSubmit={handleUpdate} style={{ maxWidth: '500px' }}>
            <div style={{ marginBottom: '15px' }}>
              <label>Título:</label>
              <input
                type="text"
                name="title"
                value={form.title}
                onChange={handleFormChange}
                required
                style={{ width: '100%', padding: '8px' }}
              />
            </div>
            <div style={{ marginBottom: '15px' }}>
              <label>Descripción:</label>
              <textarea
                name="description"
                value={form.description}
                onChange={handleFormChange}
                required
                style={{ width: '100%', padding: '8px', minHeight: '80px' }}
              />
            </div>
            <div style={{ marginBottom: '15px' }}>
              <label>Instructor:</label>
              <input
                type="text"
                name="instructor"
                value={form.instructor}
                onChange={handleFormChange}
                required
                style={{ width: '100%', padding: '8px' }}
              />
            </div>
            <div style={{ marginBottom: '15px' }}>
              <label>Horario (HH:mm):</label>
              <input
                type="text"
                name="schedule"
                value={form.schedule}
                onChange={handleFormChange}
                placeholder="10:00"
                required
                style={{ width: '100%', padding: '8px' }}
              />
            </div>
            <div style={{ marginBottom: '15px' }}>
              <label>Duración (ej: 1h 30m):</label>
              <input
                type="text"
                name="duration"
                value={form.duration}
                onChange={handleFormChange}
                placeholder="1h 30m"
                required
                style={{ width: '100%', padding: '8px' }}
              />
            </div>
            <div style={{ marginBottom: '15px' }}>
              <label>Categoría:</label>
              <input
                type="text"
                name="category"
                value={form.category}
                onChange={handleFormChange}
                required
                style={{ width: '100%', padding: '8px' }}
              />
            </div>
            <div style={{ marginBottom: '15px' }}>
              <label>Cupo:</label>
              <input
                type="number"
                name="capacity"
                value={form.capacity}
                onChange={handleFormChange}
                required
                min="1"
                style={{ width: '100%', padding: '8px' }}
              />
            </div>
            <div style={{ marginTop: '20px' }}>
              <button type="submit" style={{ marginRight: '10px', backgroundColor: '#4CAF50', color: 'white', padding: '10px 20px' }}>
                Guardar Cambios
              </button>
              <button type="button" onClick={handleCancelEdit} style={{ padding: '10px 20px' }}>
                Cancelar
              </button>
            </div>
          </form>
          {successMessage && <p style={{ color: 'green', fontWeight: 'bold', marginTop: '10px' }}>{successMessage}</p>}
          {error && <p style={{ color: 'red', marginTop: '10px' }}>{error}</p>}
        </div>
      )}
    </div>
  );
}
