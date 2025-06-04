import { useState } from 'react';
import api from './api';

export default function CreateActivity() {
  const [form, setForm] = useState({ titulo: '', horario: '', profesor: '' });

  const handleChange = e => setForm({ ...form, [e.target.name]: e.target.value });

  const submit = () => {
    api.post('/activities', form)
      .then(() => alert('✅ Creada'))
      .catch(() => alert('❌ Error'));
  };

  return (
    <div>
      <h2>Crear actividad</h2>
      <input name="titulo" onChange={handleChange} placeholder="Título" />
      <input name="horario" onChange={handleChange} placeholder="Horario" />
      <input name="profesor" onChange={handleChange} placeholder="Profesor" />
      <button onClick={submit}>Crear</button>
    </div>
  );
}
