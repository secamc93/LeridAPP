import React, { useState } from 'react';
import axios from 'axios';

function CrearNegocio() {
  const [nombre, setNombre] = useState('');
  const [email, setEmail] = useState('');
  const [contrasena, setContrasena] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Generar fecha/hora actual
    const fechaRegistro = new Date().toISOString();

    const data = {
      Nombre: nombre,
      Email: email,
      Contrasena: contrasena,
      FechaRegistro: fechaRegistro,
    };

    console.log(data); // imprimir en consola los datos del formulario

    try {
      await axios.post('http://localhost:8000/api/usuarios', data);
      // Aquí puedes agregar lógica adicional luego de enviar el POST
    } catch (error) {
      console.log(error);
    }

    // Restablecer los valores del formulario
    setNombre('');
    setEmail('');
    setContrasena('');
  };

  return (
    <div className="bg-white rounded-lg shadow p-5">
      <h1 className="text-red-500 text-white text-center mb-4">Crear Usuario</h1>
      <form onSubmit={handleSubmit}>
        <div className="mb-4">
          <label htmlFor="nombre" className="block text-red-500 text-white">
            Nombre
          </label>
          <input
            type="text"
            id="nombre"
            value={nombre}
            onChange={(e) => setNombre(e.target.value)}
            className="border border-red-500 rounded px-2 py-1"
          />
        </div>
        <div className="mb-4">
          <label htmlFor="email" className="block text-red-500 text-white">
            Email
          </label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="border border-red-500 rounded px-2 py-1"
          />
        </div>
        <div className="mb-4">
          <label htmlFor="contrasena" className="block text-red-500 text-white">
            Contraseña
          </label>
          <input
            type="password"
            id="contrasena"
            value={contrasena}
            onChange={(e) => setContrasena(e.target.value)}
            className="border border-red-500 rounded px-2 py-1"
          />
        </div>
        <button type="submit" className="bg-red-500 text-white px-4 py-2 rounded">
          Enviar
        </button>
      </form>
    </div>
  );
}

export default CrearNegocio;
