// src/components/Navbar.jsx
import { Link } from 'react-router-dom';
import { logout } from '../services/api';

export default function Navbar({ user, setUser }) {
  const handleLogout = () => {
    console.log('Navbar: Iniciando cierre de sesión');
    logout();
    console.log('Navbar: Sesión cerrada, actualizando estado');
    setUser(null);
  };

  return (
    <nav className="navbar">
      <div className="navbar-brand">
        <Link to="/">Actividades</Link>
      </div>
      <div className="navbar-menu">
        {user ? (
          <>
            <Link to="/" className="navbar-item">Inicio</Link>
            <Link to="/mis-actividades" className="navbar-item">Mis Actividades</Link>
            {user.rol === 'admin' && (
              <Link to="/admin" className="navbar-item">Admin</Link>
            )}
            <button onClick={handleLogout} className="navbar-item">Cerrar Sesión</button>
          </>
        ) : (
          <Link to="/login" className="navbar-item">Iniciar Sesión</Link>
        )}
      </div>
    </nav>
  );
}
