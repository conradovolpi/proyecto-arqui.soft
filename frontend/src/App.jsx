import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import Navbar from './components/Navbar';
import Login from './pages/Login';
import Home from './pages/Home';
import ActivityDetail from './pages/ActivityDetail';
import MyActivities from './pages/MyActivities';
import Admin from './pages/Admin';
import { getCurrentUser } from './services/mockData';

export default function App() {
  const user = getCurrentUser();

  return (
    <BrowserRouter>
      <Navbar />
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/" element={user ? <Home /> : <Navigate to="/login" />} />
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
          element={user?.role === 'admin' ? <Admin /> : <Navigate to="/" />}
        />
        <Route path="*" element={<Navigate to="/" />} />
      </Routes>
    </BrowserRouter>
  );
}
