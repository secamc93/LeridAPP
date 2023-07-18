
const ConsultarNegocios = ({ negocios }) => {
  

  return (
    <div className="container mx-auto">
      <h1 className="text-2xl font-bold mb-4">Negocios</h1>
      <div className="overflow-x-auto p-5 m-10">
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
              <th className="py-2 px-4 border-b">Fecha de Eliminación</th>
              <th className="py-2 px-4 border-b">ID de Usuario</th>
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
                <td className="py-2 px-4 border-b">{negocio.CreatedAt}</td>
                <td className="py-2 px-4 border-b">{negocio.UpdatedAt}</td>
                <td className="py-2 px-4 border-b">{negocio.DeletedAt}</td>
                <td className="py-2 px-4 border-b">{negocio.UsuarioID}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default ConsultarNegocios;
