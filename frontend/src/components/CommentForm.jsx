// src/components/CommentForm.jsx
import { useState } from 'react';
import { getCurrentUser } from '../services/api';

export default function CommentForm({ activityId, onCommentAdded }) {
  const [rating, setRating] = useState(5);
  const [comment, setComment] = useState('');
  const [error, setError] = useState('');
  const user = getCurrentUser();

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!user) {
      setError('Debes iniciar sesión para comentar');
      return;
    }

    try {
      // TODO: Implementar la funcionalidad de comentarios cuando esté disponible en el backend
      setError('La funcionalidad de comentarios estará disponible próximamente');
      // Limpiar el formulario
      setRating(5);
      setComment('');
      if (onCommentAdded) {
        onCommentAdded();
      }
    } catch (err) {
      setError(err.message || 'Error al enviar el comentario');
    }
  };

  return (
    <form onSubmit={handleSubmit} className="comment-form">
      <h4>Deja tu comentario</h4>
      {error && <p className="error-message">{error}</p>}
      
      <div className="rating-input">
        <label>Calificación:</label>
        <select value={rating} onChange={(e) => setRating(Number(e.target.value))}>
          {[5, 4, 3, 2, 1].map(num => (
            <option key={num} value={num}>
              {'⭐'.repeat(num)}
            </option>
          ))}
        </select>
      </div>

      <div className="comment-input">
        <label>Comentario:</label>
        <textarea
          value={comment}
          onChange={(e) => setComment(e.target.value)}
          placeholder="Escribe tu comentario aquí..."
          required
        />
      </div>

      <button type="submit" className="submit-button">
        Enviar Comentario
      </button>
    </form>
  );
}
