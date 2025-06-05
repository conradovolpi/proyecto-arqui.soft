// src/components/CommentForm.jsx
import { useState } from 'react';
import { addComment, getCurrentUser } from '../services/mockData';

export default function CommentForm({ activityId, onCommentAdded }) {
  const [text, setText] = useState('');
  const [rating, setRating] = useState(5);
  const [success, setSuccess] = useState('');

  const user = getCurrentUser();

  const handleSubmit = (e) => {
    e.preventDefault();

    if (!text.trim()) return;

    addComment(activityId, user.id, text, rating);
    setText('');
    setRating(5);
    setSuccess('Comentario agregado');
    onCommentAdded(); // Refresca los comentarios
  };

  return (
    <form onSubmit={handleSubmit}>
      <h4>Dejar un comentario</h4>
      <textarea
        value={text}
        onChange={(e) => setText(e.target.value)}
        placeholder="Escribe tu opinión..."
        required
      />
      <br />
      <label>Calificación: </label>
      <select value={rating} onChange={(e) => setRating(Number(e.target.value))}>
        {[5, 4, 3, 2, 1].map((r) => (
          <option key={r} value={r}>{r} ⭐</option>
        ))}
      </select>
      <br />
      <button type="submit">Comentar</button>
      {success && <p style={{ color: 'green' }}>{success}</p>}
    </form>
  );
}
