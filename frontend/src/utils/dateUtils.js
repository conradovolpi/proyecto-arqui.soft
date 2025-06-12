export const convertScheduleAndDurationToTimeRange = (schedule, duration) => {
  const now = new Date(); // Usar la fecha actual como base
  const [hours, minutes] = schedule.split(':').map(Number);

  // Establecer la hora de inicio con la fecha actual
  const startTime = new Date(now.getFullYear(), now.getMonth(), now.getDate(), hours, minutes, 0);

  let endTime = new Date(startTime.getTime()); // Copiar la hora de inicio para calcular el fin

  // Parsear la duración (ej: "1h 30m", "2h", "45m")
  const durationParts = duration.match(/(\d+h)?\s*(\d+m)?/);
  let totalMinutes = 0;

  if (durationParts && durationParts[1]) {
    totalMinutes += parseInt(durationParts[1]) * 60;
  }
  if (durationParts && durationParts[2]) {
    totalMinutes += parseInt(durationParts[2]);
  }

  endTime.setMinutes(endTime.getMinutes() + totalMinutes);

  // Devolver en formato ISO 8601 (ej: "2023-10-27T10:00:00Z") que Go puede parsear
  return {
    horario_inicio: startTime.toISOString(),
    horario_fin: endTime.toISOString(),
  };
}; 