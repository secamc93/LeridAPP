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

  return (
    <div className="container mx-auto text-white  p-6 ">
    <h1 className="text-2xl font-bold mb-4 text-center">Categoría</h1>
      <form onSubmit={handleSubmit} className="flex flex-col items-center">
        <div className="flex flex-col mb-4 w-full">
          <input
            type="text"
            id="nombre"
            placeholder="Nombre"
            className="border border-gray-300 rounded-lg p-2 w-full"
            value={nombre}
            onChange={(e) => setNombre(e.target.value)}
            style={{ color: 'black' }}
          />
        </div>
        <button
          type="submit"
          className="bg-white hover:bg-blue-500 text-red-500 hover:text-white font-bold py-2 px-4 rounded transition-colors duration-200"
        >
          Crear
        </button>
      </form>
    </div>
  );
};

export default CrearCategoria;
