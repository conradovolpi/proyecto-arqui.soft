// src/components/ActivityCard.jsx
import { Link } from 'react-router-dom';

export default function ActivityCard({ activity, onCancel, onEnroll, isEnrolled }) {
  const isFull = activity.currentEnrollments >= activity.capacity;

  return (
    <div className="activity-card">
      <h3>{activity.title}</h3>
      <p><strong>Horario:</strong> {activity.schedule}</p>
      <p><strong>Profesor:</strong> {activity.instructor}</p>
      <p><strong>Categoría:</strong> {activity.category}</p>
      <p><strong>Cupos:</strong> {activity.currentEnrollments}/{activity.capacity}</p>
      {activity.fecha_inscripcion && (
        <p><strong>Fecha de inscripción:</strong> {activity.fecha_inscripcion}</p>
      )}
      <Link to={`/actividad/${activity.id}`}>Ver detalles</Link>
      {onCancel && (
        <button 
          onClick={() => onCancel(activity.id)} 
          className="button-cancel"
        >
          Cancelar Inscripción
        </button>
      )}
      {onEnroll && !isEnrolled && !isFull && (
        <button 
          onClick={() => onEnroll(activity.id)} 
          className="button-enroll"
        >
          Inscribirse
        </button>
      )}
      {onEnroll && isFull && !isEnrolled && (
        <p className="text-red-500">Cupos llenos</p>
      )}
    </div>
  );
}
