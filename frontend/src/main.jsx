import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App.jsx';
import './styles/app.css'; // puede estar vacío o personalizado

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
