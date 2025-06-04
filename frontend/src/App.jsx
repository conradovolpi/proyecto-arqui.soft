import { Routes, Route } from 'react-router-dom';
import Home from './Home.jsx';

export default function App() {
  return (
    <>
      <h1>🏋️ Actividades Deportivas</h1>
      <Routes>
        <Route path="/" element={<Home />} />
      </Routes>
    </>
  );
}
