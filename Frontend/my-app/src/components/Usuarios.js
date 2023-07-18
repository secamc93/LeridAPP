import React from 'react';
import CrearUsuario from './CrearUsuario';
import ConsultarUsuario from './ConsultarUsuario';

const Usuarios = () => {
  return (
    <div className="flex flex-row items-center justify-center bg-white rounded-lg shadow p-4 ">
      <CrearUsuario />
      <ConsultarUsuario />
    </div>
  );
};

export default Usuarios;
