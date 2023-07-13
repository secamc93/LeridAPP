import React from 'react';
import CrearNegocio from './components/CrearNegocio';
import ConsultarNegocio from './components/ConsultarNegocio';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';

function App() {
  return (
    <Router>
      <div className="flex flex-col min-h-screen">
        <header className="bg-gradient-to-b from-red-800 to-red-600 rounded-lg shadow-lg p-6 text-white text-center">
          <h1 className="text-3xl mx-auto">LeridAPP</h1>
        </header>
        <div className="flex-grow flex justify-center items-center">
          <nav className="w-64 bg-gradient-to-b from-red-800 to-red-600 rounded-lg shadow-lg p-6 text-white m-4">
            <ul>
              <li className="mb-2">
                <Link to="/crear-negocio">Crear Negocio</Link>
              </li>
              <li>
                <Link to="/consultar-negocio">Consultar Negocio</Link>
              </li>
            </ul>
          </nav>
          <div className="flex-grow">
            <Routes>
              <Route path="/crear-negocio" element={<CrearNegocio />} />
              <Route path="/consultar-negocio" element={<ConsultarNegocio />} />
            </Routes>
          </div>
        </div>
      </div>
    </Router>
  );
}

export default App;
