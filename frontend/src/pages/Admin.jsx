// src/pages/Admin.jsx
import { useEffect, useState } from 'react';
import {
  getActivities,
  createActivity,
  updateActivity,
  deleteActivity
} from '../services/api';
import { convertScheduleAndDurationToTimeRange } from '../utils/dateUtils';


export default function Admin() {
  const [activities, setActivities] = useState([]);
  const [form, setForm] = useState({
    id: null,
    title: '',
    description: '',
    instructor: '',
    schedule: '',
    duration: '',
    category: '',
    capacity: '',
  });

  const [isEditMode, setIsEditMode] = useState(false);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const fetchActivities = async () => {
    try {
      setLoading(true);
      setError(null);
      const data = await getActivities();
      setActivities(data);
    } catch (err) {
      console.error("Error al obtener actividades:", err);
      setError("Error al cargar las actividades.");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchActivities();
  }, []);

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const resetForm = () => {
    setForm({
      id: null,
      title: '',
      description: '',
      instructor: '',
      schedule: '',
      duration: '',
      category: '',
      capacity: '',
    });
    setIsEditMode(false);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null);

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

    try {
      if (isEditMode) {
        await updateActivity(form.id, dataToSend);
      } else {
        await createActivity(dataToSend);
      }
      await fetchActivities(); // Refrescar la lista de actividades
      resetForm();
    } catch (err) {
      console.error("Error al guardar actividad:", err);
      setError(err.message || "Error al guardar la actividad.");
    }
  };

  const handleEdit = (activity) => {
    // Cuando editas, necesitas convertir las fechas de nuevo a schedule y duration
    // Esto es un placeholder; la implementación real dependerá de cómo se muestren las fechas en la UI para edición
    setForm({
      id: activity.actividad_id,
      title: activity.titulo,
      description: activity.descripcion,
      instructor: activity.instructor,
      schedule: new Date(activity.horario_inicio).toLocaleTimeString('es-ES', { hour: '2-digit', minute: '2-digit' }),
      duration: '', // No tenemos la duración original, se necesitaría un ajuste en el backend o frontend para esto
      category: activity.categoria,
      capacity: activity.cupo,
    });
    setIsEditMode(true);
  };

  const handleDelete = async (id) => {
    if (confirm('¿Estás seguro que quieres eliminar esta actividad?')) {
      setError(null);
      try {
        await deleteActivity(id);
        await fetchActivities(); // Refrescar la lista de actividades
      } catch (err) {
        console.error("Error al eliminar actividad:", err);
        setError(err.message || "Error al eliminar la actividad.");
      }
    }
  };

  if (loading) {
    return <div>Cargando actividades...</div>;
  }

  if (error) {
    return <div className="error-message">{error}</div>;
  }

  return (
    <div>
      <h2>Panel de Administración</h2>

      <form onSubmit={handleSubmit}>
        <input
          type="text"
          name="title"
          placeholder="Título"
          value={form.title}
          onChange={handleChange}
          required
        />
        <input
          type="text"
          name="description"
          placeholder="Descripción"
          value={form.description}
          onChange={handleChange}
          required
        />
        <input
          type="text"
          name="instructor"
          placeholder="Instructor"
          value={form.instructor}
          onChange={handleChange}
          required
        />
        <input
          type="text"
          name="schedule"
          placeholder="Horario (HH:mm)"
          value={form.schedule}
          onChange={handleChange}
          required
        />
        <input
          type="text"
          name="duration"
          placeholder="Duración (ej: 1h, 45m)"
          value={form.duration}
          onChange={handleChange}
          required
        />
        <input
          type="text"
          name="category"
          placeholder="Categoría"
          value={form.category}
          onChange={handleChange}
          required
        />
        <input
          type="number"
          name="capacity"
          placeholder="Cupo"
          value={form.capacity}
          onChange={handleChange}
          required
        />

        <button type="submit">{isEditMode ? 'Guardar Cambios' : 'Crear Actividad'}</button>
        {isEditMode && <button onClick={resetForm}>Cancelar</button>}
      </form>

      <hr />

      <h3>Lista de actividades</h3>
      {activities.map((a) => (
        <div key={a.id} style={{ border: '1px solid gray', padding: '10px', margin: '10px' }}>
          <strong>{a.title}</strong> — {a.schedule} — {a.instructor} <br />
          <small>{a.description}</small><br />
          <button onClick={() => handleEdit(a)}>Editar</button>
          <button onClick={() => handleDelete(a.id)}>Eliminar</button>
        </div>
      ))}
    </div>
  );
}
