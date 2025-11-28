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
    console.log('API: Iniciando login para:', email);
    const response = await fetch(`${API_URL}/usuarios/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ email, password }),
      credentials: 'include',
    });

    console.log('API: Respuesta recibida, status:', response.status);
    
    // Verificar el Content-Type antes de parsear
    const contentType = response.headers.get('content-type');
    if (!contentType || !contentType.includes('application/json')) {
      const text = await response.text();
      console.error('API: Respuesta no es JSON. Content-Type:', contentType);
      console.error('API: Respuesta como texto:', text);
      throw new Error('El servidor no devolvió una respuesta JSON válida');
    }
    
    let data;
    try {
      const text = await response.text();
      console.log('API: Respuesta como texto (antes de parsear):', text);
      if (!text || text.trim() === '') {
        throw new Error('La respuesta del servidor está vacía');
      }
      data = JSON.parse(text);
      console.log('API: Datos recibidos (parseados):', data);
    } catch (parseError) {
      console.error('API: Error al parsear JSON de la respuesta:', parseError);
      const text = await response.text();
      console.error('API: Respuesta como texto (error):', text);
      throw new Error('Error al procesar la respuesta del servidor: ' + parseError.message);
    }

    if (!response.ok) {
      // Limpiar localStorage en caso de error
      localStorage.removeItem('token');
      localStorage.removeItem('user');
      const errorMessage = data.message || data.error || 'Error al iniciar sesión';
      console.error('API: Error en respuesta:', errorMessage);
      throw new Error(errorMessage);
    }

    if (!data.token) {
      localStorage.removeItem('token');
      localStorage.removeItem('user');
      console.error('API: No se recibió token en la respuesta');
      throw new Error('No se recibió el token de autenticación');
    }

    // Validar que existe el objeto usuario
    if (!data.usuario) {
      localStorage.removeItem('token');
      localStorage.removeItem('user');
      console.error('API: No se recibió objeto usuario en la respuesta');
      throw new Error('No se recibieron los datos del usuario');
    }

    console.log('API: Objeto usuario recibido:', data.usuario);

    // Crear el objeto usuario con los datos que vienen del login
    // Manejar tanto usuario_id como UsuarioID por compatibilidad
    const usuarioId = data.usuario.usuario_id || data.usuario.UsuarioID;
    
    console.log('API: ID de usuario extraído:', usuarioId);
    
    if (!usuarioId) {
      localStorage.removeItem('token');
      localStorage.removeItem('user');
      console.error('API: El ID del usuario es undefined o null');
      throw new Error('No se recibió el ID del usuario');
    }

    const user = {
      id: usuarioId,
      email: data.usuario.email || '',
      nombre: data.usuario.nombre || '',
      rol: data.usuario.rol || ''
    };

    console.log('API: Objeto usuario creado:', user);

    // Guardar el token y el usuario en localStorage
    localStorage.setItem('token', data.token);
    localStorage.setItem('user', JSON.stringify(user));
    
    console.log('API: Login exitoso, usuario guardado en localStorage');
    return user;
  } catch (error) {
    console.error('API: Error al iniciar sesión:', error);
    // Si el error no tiene mensaje, usar uno por defecto
    if (error.message) {
      throw error;
    } else {
      throw new Error('Error de conexión. Por favor, verifica que el servidor esté en ejecución.');
    }
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

    // Verificar el Content-Type
    const contentType = response.headers.get('content-type');
    if (!contentType || !contentType.includes('application/json')) {
      const text = await response.text();
      console.error('API: Respuesta de actividades no es JSON. Content-Type:', contentType);
      console.error('API: Respuesta como texto:', text);
      throw new Error('El servidor no devolvió una respuesta JSON válida');
    }

    const text = await response.text();
    if (!text || text.trim() === '') {
      console.log('API: Respuesta vacía, devolviendo array vacío');
      return [];
    }

    let data;
    try {
      data = JSON.parse(text);
      console.log('Actividades obtenidas:', data);
      
      // Asegurar que siempre sea un array
      if (!Array.isArray(data)) {
        console.error('API: La respuesta no es un array:', data);
        throw new Error('El formato de respuesta no es válido: se esperaba un array');
      }
      
      return data;
    } catch (parseError) {
      console.error('API: Error al parsear JSON de actividades:', parseError);
      console.error('API: Respuesta como texto:', text);
      throw new Error('Error al procesar la respuesta del servidor: ' + parseError.message);
    }
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

    console.log('Intentando crear actividad:', activity);

    const response = await fetch(`${API_URL}/actividades/`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(activity),
      credentials: 'include',
    });

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({ message: 'Error al crear la actividad' }));
      throw new Error(errorData.message || 'Error al crear la actividad');
    }

    const data = await response.json().catch(() => ({ message: 'Actividad creada con éxito' }));
    return data;
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

    console.log('Intentando actualizar actividad:', { id, activity });

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
      const errorData = await response.json().catch(() => ({ message: 'Error al actualizar la actividad' }));
      throw new Error(errorData.message || 'Error al actualizar la actividad');
    }

    const data = await response.json().catch(() => ({ message: 'Actividad actualizada con éxito' }));
    return data;
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
      },
      credentials: 'include',
    });

    if (!response.ok) {
      if (response.status === 404) {
        return []; // Si no hay inscripciones, devolver array vacío
      }
      const errorData = await response.json();
      throw new Error(errorData.message || 'Error al obtener las inscripciones');
    }

    const inscriptions = await response.json();
    
    // Obtener los detalles de cada actividad
    const activitiesWithDetails = await Promise.all(
      inscriptions.map(async (inscription) => {
        try {
          const activity = await getActivity(inscription.actividad_id); // Asume que getActivity existe y funciona
          return {
            ...activity,
            fecha_inscripcion: inscription.fecha_inscripcion
          };
        } catch (error) {
          console.error(`Error al obtener detalles de actividad ${inscription.actividad_id}:`, error);
          return null; // En caso de error para una actividad, devolver null
        }
      })
    );

    // Filtrar actividades que no se pudieron cargar y devolver solo las válidas
    return activitiesWithDetails.filter(activity => activity !== null);

  } catch (error) {
    console.error('Error al obtener inscripciones:', error);
    return []; // En caso de error general, devolver array vacío
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
