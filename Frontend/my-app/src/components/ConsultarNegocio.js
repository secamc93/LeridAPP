import React, { useState, useEffect } from 'react';
import axios from 'axios';

const ConsultarNegocios = ({negocios, setNegocios}) => {
 

  useEffect(() => {
    const fetchNegocios = async () => {
      const res = await axios.get('http://localhost:8000/api/negocios');
      setNegocios(res.data);
    };

    fetchNegocios();
  }, []);

  const eliminarNegocio = async (id) => {
    try {
      const res = await axios.delete(`http://localhost:8000/api/negocios/${id}`);
      console.log(res.data);
      setNegocios(prevNegocios => prevNegocios.filter(negocio => negocio.ID !== id));

    } catch(err) {
      console.error(err);
    }
  };

  return (
    <div className="container mx-auto max-w-screen-xl">
      <h1 className="text-2xl font-bold mb-4">Negocios</h1>
      <div className="overflow-x-auto p-5 m-10" style={{ maxHeight: '500px', overflowY: 'auto' }}>
        <table className="min-w-full bg-white border border-gray-300 ">
          <thead>
            <tr>
              <th className="py-2 px-4 border-b">ID</th>
              <th className="py-2 px-4 border-b">Nombre</th>
              <th className="py-2 px-4 border-b">Dirección</th>
              <th className="py-2 px-4 border-b">Teléfono</th>
              <th className="py-2 px-4 border-b">Horario</th>
              <th className="py-2 px-4 border-b">Descripción</th>
              <th className="py-2 px-4 border-b">URL de Imagen</th>
              <th className="py-2 px-4 border-b">Fecha de Creación</th>
              <th className="py-2 px-4 border-b">Fecha de Actualización</th>
              <th className="py-2 px-4 border-b">ID de Usuario</th>
              <th className="py-2 px-4 border-b">Acciones</th>
            </tr>
          </thead>
          <tbody>
            {negocios.map((negocio) => (
              <tr key={negocio.ID}>
                <td className="py-2 px-4 border-b">{negocio.ID}</td>
                <td className="py-2 px-4 border-b">{negocio.Nombre}</td>
                <td className="py-2 px-4 border-b">{negocio.Direccion}</td>
                <td className="py-2 px-4 border-b">{negocio.Telefono}</td>
                <td className="py-2 px-4 border-b">{negocio.Horario}</td>
                <td className="py-2 px-4 border-b">{negocio.Descripcion}</td>
                <td className="py-2 px-4 border-b">{negocio.UrlImagen}</td>
                <td className="py-2 px-4 border-b">{new Date(negocio.CreatedAt).toLocaleDateString()}</td>
                <td className="py-2 px-4 border-b">{new Date(negocio.UpdatedAt).toLocaleDateString()}</td>
                <td className="py-2 px-4 border-b">{negocio.UsuarioID}</td>
                <td className="py-2 px-4 border-b">
                  <button onClick={() => eliminarNegocio(negocio.ID)} className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded">
                    Eliminar
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

export default ConsultarNegocios;

