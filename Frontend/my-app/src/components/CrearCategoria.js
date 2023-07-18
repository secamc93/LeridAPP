import React, { useState } from 'react';
import axios from 'axios';

const CrearCategoria = ({ onNuevaCategoria }) => {
    const [nombre, setNombre] = useState('');
  
    const handleSubmit = async (e) => {
      e.preventDefault();
  
      try {
        const response = await axios.post('http://localhost:8000/api/categorias', { Nombre: nombre });
        // Tras el éxito del POST, llamar a onNuevaCategoria
        onNuevaCategoria(response.data);
      } catch (error) {
        console.log(error);
      }
    };
  
    // Resto del código...

  return (
    <div className="container mx-auto">
      <h1 className="text-2xl font-bold mb-4">Crear Nueva Categoría</h1>
      <form onSubmit={handleSubmit}>
        <div className="flex flex-col mb-4">
          <label htmlFor="nombre" className="mb-2">
            Nombre
          </label>
          <input
            type="text"
            id="nombre"
            className="border border-gray-300 rounded-lg p-2"
            value={nombre}
            onChange={(e) => setNombre(e.target.value)}
          />
        </div>
        <button
          type="submit"
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        >
          Crear Categoría
        </button>
      </form>
    </div>
  );
};

export default CrearCategoria;
