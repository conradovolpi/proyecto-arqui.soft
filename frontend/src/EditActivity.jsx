import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import api from './api';

export default function EditActivity() {
  const { id } = useParams();
  const [form, setForm] = useState({ titulo: '', horario: '', profesor: '' });

  useEffect(() => {
    api.get(`/activities/${id}`).then(res => setForm(res.data));
  }, [id]);

  const handleChange = e => setForm({ ...form, [e.target.name]: e.target.value });

  const submit = () => {
    api.put(`/activities/${id}`, form)
      .then(() => alert('✅ Editada'))
      .catch(() => alert('❌ Error'));
  };

  const eliminar = () => {
    api.delete(`/activities/${id}`)
      .then(() => alert('🗑️ Eliminada'))
      .catch(() => alert('❌ Error'));
  };

  return (
    <div>
      <h2>Editar actividad</h2>
      <input name="titulo" value={form.titulo} onChange={handleChange} />
      <input name="horario" value={form.horario} onChange={handleChange} />
      <input name="profesor" value={form.profesor} onChange={handleChange} />
      <button onClick={submit}>Actualizar</button>
      <button onClick={eliminar}>Eliminar</button>
    </div>
  );
}
