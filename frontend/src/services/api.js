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
  const userStr = localStorage.getItem('user');
  if (!userStr) return null;
  
  try {
    const user = JSON.parse(userStr);
    // Asegurarnos de que el usuario tenga un ID
    if (!user || !user.id) {
      console.error('Usuario no tiene ID:', user);
      return null;
    }
    return user;
  } catch (error) {
    console.error('Error al parsear usuario:', error);
    return null;
  }
};

// Función para iniciar sesión
export const login = async (email, password) => {
  try {
    const response = await fetch(`${API_URL}/usuarios/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ email, password }),
      credentials: 'include',
    });

    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.message || 'Error al iniciar sesión');
    }

    if (!data.token) {
      throw new Error('No se recibió el token de autenticación');
    }

    // Guardar el token
    localStorage.setItem('token', data.token);
    
    // Crear el objeto usuario con los datos que vienen del login
    const user = {
      id: data.usuario.usuario_id,
      email: data.usuario.email,
      nombre: data.usuario.nombre,
      rol: data.usuario.rol
    };

    // Verificar que tenemos todos los datos necesarios
    if (!user.id) {
      throw new Error('No se recibió el ID del usuario');
    }
    
    localStorage.setItem('user', JSON.stringify(user));
    return user;
  } catch (error) {
    console.error('Error al iniciar sesión:', error);
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
    
    const headers = {
      'Content-Type': 'application/json',
      'Accept': 'application/json'
    };
    
    if (token) {
      headers['Authorization'] = `Bearer ${token}`;
    }

    console.log('Intentando obtener actividades...');
    const response = await fetch(`${API_URL}/actividades/`, {
      method: 'GET',
      headers: headers,
      credentials: 'include',
    });

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}));
      console.error('Error en la respuesta:', response.status, errorData);
      
      if (response.status === 401) {
        localStorage.removeItem('token');
        localStorage.removeItem('user');
        window.location.href = '/login';
        throw new Error('Sesión expirada');
      }
      
      throw new Error(errorData.message || `Error al obtener las actividades (${response.status})`);
    }

    const data = await response.json();
    console.log('Actividades obtenidas:', data);
    return data;
  } catch (error) {
    console.error('Error al obtener actividades:', error);
    if (error.message === 'Failed to fetch') {
      throw new Error('No se pudo conectar con el servidor. Por favor, verifica que el servidor esté en ejecución.');
    }
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
        'Content-Type': 'application/json'
      },
      credentials: 'include',
    });

    if (!response.ok) {
      if (response.status === 401) {
        // Si el token no es válido, redirigir al login
        localStorage.removeItem('token');
        localStorage.removeItem('user');
        window.location.href = '/login';
        throw new Error('Sesión expirada');
      }
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
    const usuarioID = Number(currentUser.id);
    const actividadID = Number(activityId);

    console.log('Intentando inscribirse:', { usuarioID, actividadID });

    const response = await fetch(`${API_URL}/inscripciones/`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ 
        usuario_id: usuarioID, 
        actividad_id: actividadID 
      }),
      credentials: 'include',
    });

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({ message: 'Error al inscribirse en la actividad' }));
      throw new Error(errorData.message || 'Error al inscribirse en la actividad');
    }

    const data = await response.json().catch(() => ({ message: 'Inscripción realizada con éxito' }));
    return data;
  } catch (error) {
    console.error('Error al inscribirse en actividad:', error);
    throw error;
  }
};

// Función para obtener las inscripciones de un usuario
export const getUserInscriptions = async (userId) => {
  try {
    const token = getToken();
    if (!token) {
      throw new Error('No hay token de autenticación');
    }

    const response = await fetch(`${API_URL}/inscripciones/usuario/${userId}`, {
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json'
      },
      credentials: 'include',
    });

    if (!response.ok) {
      if (response.status === 401) {
        // Si el token no es válido, redirigir al login
        localStorage.removeItem('token');
        localStorage.removeItem('user');
        window.location.href = '/login';
        throw new Error('Sesión expirada');
      }
      throw new Error('Error al obtener las inscripciones');
    }

    const inscriptions = await response.json();
    
    // Obtener los detalles de cada actividad
    const activitiesWithDetails = await Promise.all(
      inscriptions.map(async (inscription) => {
        try {
          const activity = await getActivity(inscription.actividad_id);
          return {
            ...activity,
            fecha_inscripcion: inscription.fecha_inscripcion
          };
        } catch (error) {
          console.error(`Error al obtener detalles de actividad ${inscription.actividad_id}:`, error);
          return null;
        }
      })
    );

    return activitiesWithDetails.filter(activity => activity !== null);
  } catch (error) {
    console.error('Error al obtener inscripciones:', error);
    throw error;
  }
};

// Función para desinscribirse de una actividad
export const cancelInscription = async (userId, activityId) => {
  try {
    const token = getToken();
    if (!token) {
      throw new Error('No hay token de autenticación');
    }

    console.log('Cancelando inscripción:', { userId, activityId });

    const response = await fetch(`${API_URL}/inscripciones/`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ 
        usuario_id: parseInt(userId), 
        actividad_id: parseInt(activityId) 
      }),
      credentials: 'include',
    });

    if (!response.ok) {
      if (response.status === 401) {
        localStorage.removeItem('token');
        localStorage.removeItem('user');
        window.location.href = '/login';
        throw new Error('Sesión expirada o no autorizada');
      }
      const errorData = await response.json();
      throw new Error(errorData.message || 'Error al desinscribirse de la actividad');
    }

    return await response.json();
  } catch (error) {
    console.error('Error al desinscribirse de actividad:', error);
    throw error;
  }
};
