import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import React from 'react';

import Usuarios from './components/Usuarios';
import Negocios from './components/Negocios';
import Categoria from './components/Categoria';

function App() {
  return (
    <Router>
      <div className="flex flex-col min-h-screen">
        <header className="bg-gradient-to-b from-red-800 to-red-600  shadow-lg p-6 text-white">
          <div className="flex justify-between">
            <h1 className="text-3xl">LeridAPP</h1>
            <nav className="mt-2">
              <ul className="flex space-x-4">
                <li className="mb-2">
                  <Link to="/usuarios">Usuarios</Link>
                </li>
                <li>
                  <Link to="/negocios">Negocios</Link>
                </li>
                <li>
                  <Link to="/categoria">Categorías</Link>
                </li>
              </ul>
            </nav>
          </div>
        </header>
        <div className="flex-grow">
          <Routes>
            <Route path="/usuarios" element={<Usuarios />} />
            <Route path="/negocios" element={<Negocios />} />
            <Route path="/categoria" element={<Categoria />} />
          </Routes>
        </div>
      </div>
    </Router>
  );
}

export default App;

