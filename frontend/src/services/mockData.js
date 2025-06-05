// src/services/mockData.js

export const users = [
  { id: 1, username: 'admin', password: 'admin', role: 'admin' },
  { id: 2, username: 'user', password: 'user', role: 'user' },
];

export let activities = [ /* tus actividades mock aquí */ ];

export const login = (username, password) => {
  const found = users.find(u => u.username === username && u.password === password);
  if (found) {
    localStorage.setItem('user', JSON.stringify(found)); // 🧠 persistencia
    return found;
  }
  return null;
};

export const logout = () => {
  localStorage.removeItem('user');
};

export const getCurrentUser = () => {
  const raw = localStorage.getItem('user');
  return raw ? JSON.parse(raw) : null;
};

  
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
  