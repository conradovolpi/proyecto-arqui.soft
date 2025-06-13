// src/pages/MyActivities.jsx
import { useEffect, useState } from 'react';
import { getCurrentUser, getUserInscriptions, cancelInscription } from '../services/api';
import ActivityCard from '../components/ActivityCard';
import { useNavigate } from 'react-router-dom';

export default function MyActivities() {
  const [myActivities, setMyActivities] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  const fetchActivities = async () => {
    const user = getCurrentUser();
    if (!user) {
      setError("No hay usuario autenticado");
      setLoading(false);
      navigate('/login');
      return;
    }

    try {
      setLoading(true);
      setError(null);
      console.log('Obteniendo actividades para usuario:', user.id);
      const activities = await getUserInscriptions(user.id);
      setMyActivities(activities);
    } catch (err) {
      console.error("Error al obtener actividades:", err);
      setError("Error al cargar tus actividades inscritas.");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchActivities();
  }, [navigate]);

  const handleCancelInscription = async (activityId) => {
    const user = getCurrentUser();
    if (!user || !user.id) {
      setError("No hay usuario autenticado para cancelar la inscripción.");
      return;
    }
    
    if (window.confirm('¿Estás seguro de que quieres cancelar la inscripción a esta actividad?')) {
      try {
        setLoading(true);
        setError(null);
        console.log('Intentando cancelar inscripción:', { userId: user.id, activityId });
        await cancelInscription(user.id, activityId);
        alert('Inscripción cancelada con éxito.');
        await fetchActivities(); // Refrescar la lista de actividades
      } catch (err) {
        console.error("Error al cancelar inscripción:", err);
        setError(err.message || "Error al cancelar la inscripción.");
      } finally {
        setLoading(false);
      }
    }
  };

  if (loading) {
    return <div>Cargando tus actividades...</div>;
  }

  if (error) {
    return <div className="error-message">{error}</div>;
  }

  return (
    <div>
      <h2>Mis Actividades</h2>
      {myActivities.length > 0 ? (
        <div className="activities-grid">
          {myActivities.map((activity) => {
            console.log('Actividad en MyActivities:', activity);
            return (
              <ActivityCard 
                key={activity.actividad_id} 
                activity={{
                  id: activity.actividad_id,
                  title: activity.titulo,
                  schedule: `${new Date(activity.horario_inicio).toLocaleTimeString('es-ES', { hour: '2-digit', minute: '2-digit' })} - ${new Date(activity.horario_fin).toLocaleTimeString('es-ES', { hour: '2-digit', minute: '2-digit' })}`,
                  instructor: activity.instructor,
                  category: activity.categoria,
                  description: activity.descripcion,
                  fecha_inscripcion: new Date(activity.fecha_inscripcion).toLocaleDateString('es-ES')
                }} 
                onCancel={handleCancelInscription}
              />
            );
          })}
        </div>
      ) : (
        <p>No estás inscrito en ninguna actividad aún.</p>
      )}
    </div>
  );
}
