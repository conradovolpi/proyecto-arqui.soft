import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { useState, useEffect } from 'react';
import Navbar from './components/Navbar';
import Login from './pages/Login';
import Home from './pages/Home';
import ActivityDetail from './pages/ActivityDetail';
import MyActivities from './pages/MyActivities';
import Admin from './pages/Admin';
import { getCurrentUser } from './services/api';

export default function App() {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    console.log('App: Iniciando carga de usuario');
    const currentUser = getCurrentUser();
    console.log('App: Usuario cargado:', currentUser);
    setUser(currentUser);
    setLoading(false);
  }, []);

  const handleUserChange = (newUser) => {
    console.log('App: Actualizando usuario:', newUser);
    setUser(newUser);
  };

  if (loading) {
    return <div>Cargando...</div>;
  }

  console.log('App: Renderizando con usuario:', user);

  return (
    <BrowserRouter>
      <div className="app">
        <Navbar user={user} setUser={handleUserChange} />
        <main className="container">
          <Routes>
            <Route 
              path="/login" 
              element={!user ? <Login setUser={handleUserChange} /> : <Navigate to="/" />} 
            />
            <Route 
              path="/" 
              element={user ? <Home /> : <Navigate to="/login" />} 
            />
            <Route
              path="/actividad/:id"
              element={user ? <ActivityDetail /> : <Navigate to="/login" />}
            />
            <Route
              path="/mis-actividades"
              element={user ? <MyActivities /> : <Navigate to="/login" />}
            />
            <Route
              path="/admin"
              element={user?.rol === 'admin' ? <Admin /> : <Navigate to="/" />}
            />
            <Route path="*" element={<Navigate to="/" />} />
          </Routes>
        </main>
      </div>
    </BrowserRouter>
  );
}
