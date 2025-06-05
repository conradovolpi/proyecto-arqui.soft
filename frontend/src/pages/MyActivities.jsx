// src/pages/MyActivities.jsx
import { useEffect, useState } from 'react';
import { getCurrentUser, getUserActivities } from '../services/mockData';
import ActivityCard from '../components/ActivityCard';

export default function MyActivities() {
  const [myActivities, setMyActivities] = useState([]);
  const user = getCurrentUser();

  useEffect(() => {
    if (user) {
      const data = getUserActivities(user.id);
      setMyActivities(data);
    }
  }, [user]);

  return (
    <div>
      <h2>Mis Actividades</h2>
      {myActivities.length > 0 ? (
        myActivities.map((act) => (
          <ActivityCard key={act.id} activity={act} />
        ))
      ) : (
        <p>No estás inscrito en ninguna actividad aún.</p>
      )}
    </div>
  );
}
