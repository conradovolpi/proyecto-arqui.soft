// src/components/ActivityCard.jsx
import { Link } from 'react-router-dom';

export default function ActivityCard({ activity }) {
  return (
    <div className="activity-card">
      <h3>{activity.title}</h3>
      <p><strong>Horario:</strong> {activity.schedule}</p>
      <p><strong>Profesor:</strong> {activity.instructor}</p>
      <p><strong>Categoría:</strong> {activity.category}</p>
      <Link to={`/actividad/${activity.id}`}>Ver detalles</Link>
    </div>
  );
}
