// src/components/Navbar.jsx
import { Link, useNavigate } from 'react-router-dom';
import { getCurrentUser, logout } from '../services/mockData';

export default function Navbar() {
  const user = getCurrentUser();
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  return (
    <nav>
      {user ? (
        <>
          <Link to="/">Inicio</Link> |{' '}
          <Link to="/mis-actividades">Mis Actividades</Link>{' '}
          {user.role === 'admin' && <>| <Link to="/admin">Admin</Link></>}
          {' '}| <button onClick={handleLogout}>Logout</button>
        </>
      ) : (
        <>
          <Link to="/login">Login</Link>
        </>
      )}
    </nav>
  );
}
