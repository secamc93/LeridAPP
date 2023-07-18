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
    <div>
      <h2>Crear Categoría</h2>
      <CrearCategoria onNuevaCategoria={handleNuevaCategoria} />
      <h2>Consultar Categoría</h2>
      <ConsultarCategoria categorias={categorias} />
    </div>
  );
};

export default Categoria;
