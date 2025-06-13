// src/pages/Home.jsx
import { useEffect, useState } from 'react';
import { getActivities, enrollInActivity, getCurrentUser, getUserInscriptions } from '../services/api';
import ActivityCard from '../components/ActivityCard';

export default function Home() {
  const [activities, setActivities] = useState([]);
  const [filtered, setFiltered] = useState([]);
  const [search, setSearch] = useState('');
  const [selectedCategory, setSelectedCategory] = useState('');
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [userInscriptions, setUserInscriptions] = useState([]);
  const currentUser = getCurrentUser();

  // Obtener categorías únicas de las actividades
  const categories = [...new Set(activities.map(act => act.categoria))];

  const fetchAllActivities = async () => {
    try {
      setLoading(true);
      setError(null);
      console.log('Iniciando carga de actividades...');
      
      const allActivities = await getActivities();
      console.log('Actividades cargadas:', allActivities);
      
      if (!Array.isArray(allActivities)) {
        throw new Error('El formato de respuesta no es válido');
      }

      setActivities(allActivities);
      setFiltered(allActivities);

      if (currentUser && currentUser.id) {
        console.log('Cargando inscripciones del usuario:', currentUser.id);
        const inscriptions = await getUserInscriptions(currentUser.id);
        setUserInscriptions(inscriptions.map(insc => insc.id));
      }

    } catch (err) {
      console.error("Error al obtener actividades:", err);
      setError(err.message || "Error al cargar las actividades.");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchAllActivities();
  }, [currentUser]);

  const handleSearch = (e) => {
    const value = e.target.value.toLowerCase();
    setSearch(value);
    filterActivities(value, selectedCategory);
  };

  const handleCategoryChange = (e) => {
    const category = e.target.value;
    setSelectedCategory(category);
    filterActivities(search, category);
  };

  const filterActivities = (searchValue, category) => {
    let results = activities;

    // Filtrar por búsqueda
    if (searchValue) {
      results = results.filter(
        (act) =>
          act.titulo.toLowerCase().includes(searchValue) ||
          act.categoria.toLowerCase().includes(searchValue) ||
          (typeof act.horario_inicio === 'string' && act.horario_inicio.includes(searchValue)) ||
          (typeof act.horario_fin === 'string' && act.horario_fin.includes(searchValue))
      );
    }

    // Filtrar por categoría
    if (category) {
      results = results.filter(act => act.categoria === category);
    }

    setFiltered(results);
  };

  const handleEnroll = async (activityId) => {
    if (!currentUser || !currentUser.id) {
      setError("Debes iniciar sesión para inscribirte en una actividad.");
      return;
    }

    if (window.confirm('¿Estás seguro de que quieres inscribirte en esta actividad?')) {
      try {
        await enrollInActivity(activityId);
        alert('Inscripción realizada con éxito.');
        // Refrescar la lista de actividades y las inscripciones del usuario
        await fetchAllActivities();
      } catch (err) {
        console.error("Error al inscribirse en actividad:", err);
        setError(err.message || "Error al inscribirse en la actividad.");
      }
    }
  };

  const handleRetry = () => {
    setError(null);
    fetchAllActivities();
  };

  if (loading) {
    return (
      <div className="loading-container">
        <div className="loading-spinner"></div>
        <p>Cargando actividades...</p>
      </div>
    );
  }

  if (error) {
    return (
      <div className="error-container">
        <p className="error-message">{error}</p>
        <button onClick={handleRetry} className="retry-button">
          Reintentar
        </button>
      </div>
    );
  }

  return (
    <div className="home-container">
      <h2 className="page-title">Actividades Deportivas</h2>

      <div className="filters-container">
        <div className="search-box">
          <input
            type="text"
            placeholder="Buscar por título, categoría u horario"
            value={search}
            onChange={handleSearch}
            className="search-input"
          />
        </div>

        <div className="category-filter">
          <select 
            value={selectedCategory} 
            onChange={handleCategoryChange}
            className="category-select"
          >
            <option value="">Todas las categorías</option>
            {categories.map(category => (
              <option key={category} value={category}>{category}</option>
            ))}
          </select>
        </div>
      </div>

      {filtered.length > 0 ? (
        <div className="activities-grid">
          {filtered.map((act) => (
            <ActivityCard 
              key={act.actividad_id}
              activity={{
                id: act.actividad_id,
                title: act.titulo,
                schedule: `${new Date(act.horario_inicio).toLocaleTimeString('es-ES', { hour: '2-digit', minute: '2-digit' })} - ${new Date(act.horario_fin).toLocaleTimeString('es-ES', { hour: '2-digit', minute: '2-digit' })}`,
                instructor: act.instructor,
                category: act.categoria,
                description: act.descripcion,
                capacity: act.cupo,
                currentEnrollments: act.inscripciones ? act.inscripciones.length : 0,
              }}
              onEnroll={handleEnroll}
              isEnrolled={userInscriptions.includes(act.actividad_id)}
            />
          ))}
        </div>
      ) : (
        <div className="no-results">
          <p>No se encontraron actividades que coincidan con tu búsqueda.</p>
          <button 
            onClick={() => {
              setSearch('');
              setSelectedCategory('');
              setFiltered(activities);
            }}
            className="clear-filters-button"
          >
            Limpiar filtros
          </button>
        </div>
      )}
    </div>
  );
}
