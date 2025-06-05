// src/pages/Admin.jsx
import { useEffect, useState } from 'react';
import {
  getActivities,
  addActivity,
  editActivity,
  deleteActivity
} from '../services/mockData';


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

  useEffect(() => {
    setActivities(getActivities());
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

  const handleSubmit = (e) => {
    e.preventDefault();

    const data = {
      ...form,
      capacity: Number(form.capacity),
    };

    if (isEditMode) {
      editActivity(form.id, data);
    } else {
      addActivity(data);
    }

    setActivities(getActivities());
    resetForm();
  };

  const handleEdit = (activity) => {
    setForm(activity);
    setIsEditMode(true);
  };

  const handleDelete = (id) => {
    if (confirm('¿Estás seguro que quieres eliminar esta actividad?')) {
      deleteActivity(id);
      setActivities(getActivities());
    }
  };

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
