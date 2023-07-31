import React, { useState, useEffect } from 'react';
import axios from 'axios';

const AsignarCategorias = () => {
  const [todasLasCategorias, setTodasLasCategorias] = useState([]);
  const [todosLosNegocios, setTodosLosNegocios] = useState([]);
  const [categoriasSeleccionadas, setCategoriasSeleccionadas] = useState([]);
  const [negocioSeleccionado, setNegocioSeleccionado] = useState(null);

  useEffect(() => {
    const fetchCategorias = async () => {
      try {
        const response = await axios.get('http://localhost:8000/api/categorias');
        setTodasLasCategorias(response.data);
      } catch (error) {
        console.log(error);
      }
    };

    const fetchNegocios = async () => {
      try {
        const response = await axios.get('http://localhost:8000/api/negocios');
        setTodosLosNegocios(response.data);
      } catch (error) {
        console.log(error);
      }
    };

    fetchCategorias();
    fetchNegocios();
  }, []);

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const categoriaNegocios = categoriasSeleccionadas.map(categoriaId => {
        return { NegocioID: Number(negocioSeleccionado), CategoríaID: Number(categoriaId) };
      });

      console.log(categoriaNegocios);

      await Promise.all(categoriaNegocios.map(categoriaNegocio => {
        return axios.post('http://localhost:8000/api/categoriasnegocios', categoriaNegocio);
      }));

      setCategoriasSeleccionadas([]);

    } catch (error) {
      console.log(error);
    }
  };

  return (
    <div className="bg-red-500 rounded-lg shadow-lg p-6 overflow-auto ">
      <h2 className="text-2xl font-bold mb-4 text-center text-white">Asignar Categorías al Negocio</h2>
      <form onSubmit={handleSubmit} className="flex flex-col space-y-4">
        <select
          value={negocioSeleccionado}
          onChange={(e) => setNegocioSeleccionado(e.target.value)}
          className="border border-gray-300 rounded-lg p-2 focus:outline-none focus:border-white text-white bg-red-400"
        >
          {todosLosNegocios.map((negocio) => (
            <option key={negocio.ID} value={negocio.ID}>
              {negocio.Nombre}
            </option>
          ))}
        </select>

        <select 
          multiple={true}
          value={categoriasSeleccionadas}
          onChange={(e) => setCategoriasSeleccionadas(Array.from(e.target.selectedOptions, option => option.value))}
          className="border border-gray-300 rounded-lg p-2 focus:outline-none focus:border-white text-white bg-red-400"
        >
          {todasLasCategorias.map((categoria) => (
            <option key={categoria.ID} value={categoria.ID}>
              {categoria.Nombre}
            </option>
          ))}
        </select>
        <button 
          type="submit" 
          className="px-4 py-2 rounded-lg bg-white text-red-500"
        >
          Asignar Categorías
        </button>
      </form>
    </div>
  );
};

export default AsignarCategorias;

