// src/pages/Login.jsx
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { login, getCurrentUser } from '../services/api';

export default function Login({ setUser }) {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleLogin = async (e) => {
    e.preventDefault();
    setError('');
    setLoading(true);

    try {
      console.log('Login: Intentando iniciar sesión...');
      const user = await login(email, password);
      console.log('Login: Login exitoso, usuario recibido:', user);
      
      // Verificar que el usuario se guardó correctamente en localStorage
      const savedUser = getCurrentUser();
      console.log('Login: Usuario leído de localStorage:', savedUser);
      
      if (!savedUser || !savedUser.id) {
        console.error('Login: El usuario no se guardó correctamente en localStorage');
        throw new Error('Error al guardar la sesión. Por favor, intenta nuevamente.');
      }
      
      // Actualizar el estado del usuario
      console.log('Login: Actualizando estado del usuario');
      setUser(savedUser);
      
      // Forzar una recarga completa para asegurar que App.jsx lea el usuario de localStorage
      // Esto garantiza que el estado se sincronice correctamente
      console.log('Login: Recargando página para sincronizar estado');
      window.location.href = '/';
    } catch (err) {
      console.error('Login: Error capturado:', err);
      setError(err.message || 'Error al iniciar sesión');
      setLoading(false);
    }
  };

  return (
    <div className="login-container">
      <h2>Iniciar Sesión</h2>
      {error && <div className="error-message">{error}</div>}
      <form onSubmit={handleLogin}>
        <div>
          <label htmlFor="email">Email:</label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div>
          <label htmlFor="password">Contraseña:</label>
          <input
            type="password"
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <button type="submit" disabled={loading}>
          {loading ? 'Iniciando sesión...' : 'Iniciar Sesión'}
        </button>
      </form>
    </div>
  );
}
