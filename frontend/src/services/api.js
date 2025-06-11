import jwtDecode from 'jwt-decode';

const API_URL = 'http://localhost:8080';

export async function login(email, password) {
  const res = await fetch(`${API_URL}/usuarios/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password }),
  });

  if (!res.ok) throw new Error('Credenciales incorrectas');

  const { token } = await res.json();
  localStorage.setItem('token', token);
  return token;
}

export function getToken() {
  return localStorage.getItem('token');
}

export function logout() {
  localStorage.removeItem('token');
}

export async function getActividades() {
  const res = await fetch(`${API_URL}/actividades`);
  return await res.json();
}

export async function getActividadById(id) {
  const res = await fetch(`${API_URL}/actividades/${id}`);
  return await res.json();
}

export async function inscribirUsuario(usuarioID, actividadID) {
  const token = getToken();
  const res = await fetch(`${API_URL}/inscripciones`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify({ usuario_id: usuarioID, actividad_id: actividadID }),
  });
  return await res.json();
}



export function getCurrentUser() {
  const token = getToken();
  if (!token) return null;
  try {
    return jwtDecode(token);
  } catch {
    return null;
  }
}
