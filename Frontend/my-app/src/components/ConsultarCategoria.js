import React from 'react';

const ConsultarCategoria = ({ categorias }) => {

  return (
    <div className= "shadow-lg rounded bg-white p-5" style={{border: '1px solid red'}}>
      <h1 className="text-2xl font-bold mb-4">Categorías</h1>
      <table className="min-w-full bg-white border border-gray-300">
        <thead>
          <tr>
            <th className="py-2 px-4 border-b">ID</th>
            <th className="py-2 px-4 border-b">Nombre</th>
          </tr>
        </thead>
        <tbody>
          {categorias.map((categoria) => (
            <tr key={categoria.ID}>
              <td className="py-2 px-4 border-b">{categoria.ID}</td>
              <td className="py-2 px-4 border-b">{categoria.Nombre}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default ConsultarCategoria;
