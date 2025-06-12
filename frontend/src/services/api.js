import { jwtDecode } from 'jwt-decode';

const API_URL = 'http://localhost:8080';

const defaultHeaders = {
  'Content-Type': 'application/json',
  'Accept': 'application/json',
};

function getAuthHeaders() {
  const token = getToken();
  return token ? { ...defaultHeaders, 'Authorization': `Bearer ${token}` } : defaultHeaders;
}

// Función para obtener el token del localStorage
const getToken = () => {
  const token = localStorage.getItem('token');
  console.log('API: Token obtenido:', token ? 'existe' : 'no existe');
  return token;
};

// Función para guardar el token en localStorage
const setToken = (token) => {
  console.log('API: Guardando token');
  localStorage.setItem('token', token);
};

// Función para eliminar el token del localStorage
const removeToken = () => {
  console.log('API: Eliminando token');
  localStorage.removeItem('token');
};

// Función para obtener el usuario actual
export const getCurrentUser = () => {
  const token = getToken();
  if (!token) {
    console.log('API: No hay token, usuario no autenticado');
    return null;
  }

  try {
    const payload = JSON.parse(atob(token.split('.')[1]));
    console.log('API: Usuario decodificado:', payload);
    return payload;
  } catch (error) {
    console.error('API: Error al decodificar token:', error);
    removeToken();
    return null;
  }
};

// Función para hacer login
export const login = async (email, password) => {
  try {
    console.log('API: Iniciando login para:', email);
    const response = await fetch(`${API_URL}/usuarios/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ email, password }),
    });

    const data = await response.json();
    console.log('API: Respuesta de login:', data);

    if (!response.ok) {
      throw new Error(data.message || 'Error en el login');
    }

    if (!data.token) {
      throw new Error('No se recibió el token de autenticación');
    }

    setToken(data.token);
    return data;
  } catch (error) {
    console.error('API: Error en login:', error);
    if (error.name === 'TypeError' && error.message === 'Failed to fetch') {
      throw new Error('No se puede conectar con el servidor. Por favor, verifique que el servidor esté en ejecución en http://localhost:8080');
    }
    throw error;
  }
};

// Función para hacer logout
export const logout = () => {
  console.log('API: Iniciando logout');
  removeToken();
};

// Función para obtener todas las actividades
export const getActivities = async () => {
  try {
    const token = getToken();
    if (!token) {
      throw new Error('No hay token de autenticación');
    }

    const response = await fetch(`${API_URL}/actividades`, {
      headers: {
        'Authorization': `Bearer ${token}`,
      },
      credentials: 'include',
    });

    if (!response.ok) {
      throw new Error('Error al obtener las actividades');
    }

    return await response.json();
  } catch (error) {
    console.error('Error al obtener actividades:', error);
    throw error;
  }
};

// Función para obtener una actividad específica
export const getActivity = async (id) => {
  try {
    const token = getToken();
    if (!token) {
      throw new Error('No hay token de autenticación');
    }

    const response = await fetch(`${API_URL}/actividades/${id}`, {
      headers: {
        'Authorization': `Bearer ${token}`,
      },
      credentials: 'include',
    });

    if (!response.ok) {
      throw new Error('Error al obtener la actividad');
    }

    return await response.json();
  } catch (error) {
    console.error('Error al obtener actividad:', error);
    throw error;
  }
};

// Función para crear una actividad
export const createActivity = async (activity) => {
  try {
    const token = getToken();
    if (!token) {
      throw new Error('No hay token de autenticación');
    }

    const response = await fetch(`${API_URL}/actividades`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(activity),
      credentials: 'include',
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Error al crear la actividad');
    }

    return await response.json();
  } catch (error) {
    console.error('Error al crear actividad:', error);
    throw error;
  }
};

// Función para actualizar una actividad
export const updateActivity = async (id, activity) => {
  try {
    const token = getToken();
    if (!token) {
      throw new Error('No hay token de autenticación');
    }

    const response = await fetch(`${API_URL}/actividades/${id}`, {
      method: 'PUT',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(activity),
      credentials: 'include',
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Error al actualizar la actividad');
    }

    return await response.json();
  } catch (error) {
    console.error('Error al actualizar actividad:', error);
    throw error;
  }
};

// Función para eliminar una actividad
export const deleteActivity = async (id) => {
  try {
    const token = getToken();
    if (!token) {
      throw new Error('No hay token de autenticación');
    }

    const response = await fetch(`${API_URL}/actividades/${id}`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
      credentials: 'include',
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Error al eliminar la actividad');
    }

    return await response.json();
  } catch (error) {
    console.error('Error al eliminar actividad:', error);
    throw error;
  }
};

// Función para inscribirse en una actividad
export const enrollInActivity = async (activityId) => {
  try {
    const token = getToken();
    if (!token) {
      throw new Error('No hay token de autenticación');
    }

    const currentUser = getCurrentUser();
    if (!currentUser || !currentUser.id) {
      throw new Error('No se pudo obtener el ID del usuario actual para la inscripción.');
    }
    const usuarioID = currentUser.id;

    const response = await fetch(`${API_URL}/inscripciones`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ usuario_id: usuarioID, actividad_id: activityId }),
      credentials: 'include',
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Error al inscribirse en la actividad');
    }

    return await response.json();
  } catch (error) {
    console.error('Error al inscribirse en actividad:', error);
    throw error;
  }
};
