import React, { useState, useEffect } from 'react';
import axios from 'axios';
import CrearCategoria from './CrearCategoria';
import ConsultarCategoria from './ConsultarCategoria';

const Categoria = () => {
  const [categorias, setCategorias] = useState([]);

  useEffect(() => {
    const fetchCategorias = async () => {
      try {
        const response = await axios.get('http://localhost:8000/api/categorias');
        setCategorias(response.data);
      } catch (error) {
        console.log(error);
      }
    };

    fetchCategorias();
  }, []);

  const handleNuevaCategoria = (nuevaCategoria) => {
    setCategorias([...categorias, nuevaCategoria]);
  };

  return (
    <div className="bg-red-500 rounded-lg shadow-lg flex flex-row justify-between">
      <CrearCategoria onNuevaCategoria={handleNuevaCategoria} />
      
      <ConsultarCategoria categorias={categorias} setCategorias={setCategorias} />
    </div>
  );
};

export default Categoria;
