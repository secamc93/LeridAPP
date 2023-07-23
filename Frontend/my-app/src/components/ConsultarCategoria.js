import React from 'react';
import axios from 'axios';

const ConsultarCategoria = ({ categorias, fetchCategorias, setCategorias }) => {
  const eliminarCategoria = async (id) => {

    try {
      await axios.delete(`http://localhost:8000/api/categorias/${id}`)
      .then(() => {
        setCategorias(categorias.filter((categoria) => categoria.ID !== id))
        fetchCategorias();
      });
    } catch (error) {
      console.log(error);
    }
  };

  
  



  return (
    <div className="p-5">
      
      <div className="max-h-64 overflow-auto rounded-lg">
        <table className="min-w-full bg-white border border-gray-300 rounded-lg">
          <thead>
            <tr>
              
             
            </tr>
          </thead>
          <tbody>
            {categorias.map((categoria) => (
              <tr key={categoria.ID}>
                <td className="py-2 px-4 border-b">{categoria.Nombre}</td>
                <td className="py-2 px-4 border-b">
                  <button onClick={() => eliminarCategoria(categoria.ID)} className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded">
                    X
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default ConsultarCategoria;
