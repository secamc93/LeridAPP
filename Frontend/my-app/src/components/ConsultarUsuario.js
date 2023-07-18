import React, { useState, useEffect } from 'react';
import axios from 'axios';

function ConsultarUsuario() {
  const [datos, setDatos] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get('http://localhost:8000/api/usuarios');
        setDatos(response.data);
      } catch (error) {
        console.error('Error fetching data: ', error);
      }
    };

    fetchData();
  }, []);

  return (
    <div className="shadow-lg rounded bg-white p-5 m-10" >
        <h1 className="text-red-500 text-white text-center mb-4">Usuarios</h1>
      <table className="table-auto w-full">
        <thead className="bg-red-500 text-white">
          <tr>
            <th className="px-4 py-2">Nombre</th>
            <th className="px-4 py-2">Email</th>
            <th className="px-4 py-2">Fecha de Registro</th>
          </tr>
        </thead>
        <tbody>
          {datos.map((dato, index) => (
            <tr key={index} className={`text-red-500 ${index % 2 === 0 ? 'bg-red-100' : ''}`}>
              <td className="border px-4 py-2">{dato.Nombre}</td>
              <td className="border px-4 py-2">{dato.Email}</td>
              <td className="border px-4 py-2">{new Date(dato.FechaRegistro).toLocaleString()}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default ConsultarUsuario;
