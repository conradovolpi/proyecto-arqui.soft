// src/services/mockData.js
export const users = [
    { id: 1, username: 'admin', password: 'admin', role: 'admin' },
    { id: 2, username: 'user', password: 'user', role: 'user' },
  ];
  
  export let activities = [
    {
      id: 1,
      title: 'Yoga',
      description: 'Clase de yoga relajante',
      instructor: 'María',
      duration: '1h',
      schedule: '08:00',
      category: 'Bienestar',
      capacity: 5,
      enrolledUsers: [2],
      comments: [
        { userId: 2, text: 'Muy buena clase', rating: 5 },
      ],
    },
    {
      id: 2,
      title: 'Crossfit',
      description: 'Entrenamiento de alta intensidad',
      instructor: 'Juan',
      duration: '45m',
      schedule: '10:00',
      category: 'Fuerza',
      capacity: 10,
      enrolledUsers: [],
      comments: [],
    },
  ];
  
  export let sessions = {
    user: null,
  };
  
  // login simulado
  export const login = (username, password) => {
    const found = users.find(u => u.username === username && u.password === password);
    if (found) {
      sessions.user = found;
      return found;
    }
    return null;
  };
  
  export const logout = () => {
    sessions.user = null;
  };
  
  export const getCurrentUser = () => sessions.user;
  
  export const getActivities = () => activities;
  
  export const getActivityById = (id) => activities.find(a => a.id === Number(id));
  
  export const enrollUser = (activityId, userId) => {
    const activity = getActivityById(activityId);
    if (!activity) return { success: false, msg: 'Actividad no encontrada' };
  
    if (activity.enrolledUsers.includes(userId)) {
      return { success: false, msg: 'Ya estás inscrito' };
    }
  
    if (activity.enrolledUsers.length >= activity.capacity) {
      return { success: false, msg: 'Cupo lleno' };
    }
  
    activity.enrolledUsers.push(userId);
    return { success: true, msg: 'Inscripción exitosa' };
  };
  
  export const getUserActivities = (userId) =>
    activities.filter(a => a.enrolledUsers.includes(userId));
  
  export const addComment = (activityId, userId, text, rating) => {
    const activity = getActivityById(activityId);
    if (activity) {
      activity.comments.push({ userId, text, rating });
    }
  };
  
  export const addActivity = (data) => {
    const newActivity = {
      ...data,
      id: activities.length + 1,
      enrolledUsers: [],
      comments: [],
    };
    activities.push(newActivity);
  };
  
  export const editActivity = (id, data) => {
    const index = activities.findIndex(a => a.id === Number(id));
    if (index !== -1) {
      activities[index] = { ...activities[index], ...data };
    }
  };
  
  export const deleteActivity = (id) => {
    activities = activities.filter(a => a.id !== Number(id));
  };
  