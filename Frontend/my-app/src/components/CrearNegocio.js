import React, { useState, useEffect } from 'react';
import axios from 'axios';

const CrearNegocio = ({ setNegocios }) => {


  const [nombre, setNombre] = useState('');
  const [direccion, setDireccion] = useState('');
  const [telefono, setTelefono] = useState('');
  const [horario, setHorario] = useState('');
  const [descripcion, setDescripcion] = useState('');
  const [urlImagen, setUrlImagen] = useState('');
  const [negociosCategorias, setNegociosCategorias] = useState([]);
  const [todasLasCategorias, setTodasLasCategorias] = useState([]);

  useEffect(() => {
    const fetchCategorias = async () => {
      try {
        const res = await axios.get('http://localhost:8000/api/categorias');
        setTodasLasCategorias(res.data);
      } catch (error) {
        console.log(error);
      }
    };

    fetchCategorias();
  }, []);

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const nuevoNegocio = {
        Nombre: nombre,
        Direccion: direccion,
        Telefono: telefono,
        Horario: horario,
        Descripcion: descripcion,
        UrlImagen: urlImagen,
        NegociosCategorias: negociosCategorias.map(catNombre => {
          const categoria = todasLasCategorias.find(cat => cat.nombre === catNombre);
          return categoria ? categoria.id : null;
        }),
        UsuarioID: 1,
        Comentarios: []
      };

      const response = await axios.post('http://localhost:8000/api/negocios', nuevoNegocio);

    // Actualiza la lista de negocios en el estado.
      setNegocios(prevNegocios => [...prevNegocios, response.data]);

      setNombre('');
      setDireccion('');
      setTelefono('');
      setHorario('');
      setDescripcion('');
      setUrlImagen('');
      setNegociosCategorias([]);
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <div className="container mx-auto p-10 shadow rounded-lg">
      <h1 className="text-2xl font-bold mb-4 text-center">Crear Negocio</h1>
      <form onSubmit={handleSubmit}>
        <div className="flex flex-col mb-4">
          <div className="flex justify-center">
            <label htmlFor="nombre" className="mb-2">
              Nombre
            </label>
          </div>
          <input
            type="text"
            id="nombre"
            className="border border-gray-300 rounded-lg p-2 focus:outline-none focus:border-red-500"
            value={nombre}
            onChange={(e) => setNombre(e.target.value)}
          />
        </div>
        <div className="flex flex-col mb-4">
          <div className="flex justify-center">
            <label htmlFor="direccion" className="mb-2">
              Dirección
            </label>
          </div>
          <input
            type="text"
            id="direccion"
            className="border border-gray-300 rounded-lg p-2 focus:outline-none focus:border-red-500"
            value={direccion}
            onChange={(e) => setDireccion(e.target.value)}
          />
        </div>
        <div className="flex flex-col mb-4">
          <div className="flex justify-center">
            <label htmlFor="telefono" className="mb-2">
              Teléfono
            </label>
          </div>
          <input
            type="text"
            id="telefono"
            className="border border-gray-300 rounded-lg p-2 focus:outline-none focus:border-red-500"
            value={telefono}
            onChange={(e) => setTelefono(e.target.value)}
          />
        </div>
        <div className="flex flex-col mb-4">
          <div className="flex justify-center">
            <label htmlFor="horario" className="mb-2">
              Horario
            </label>
          </div>
          <input
            type="text"
            id="horario"
            className="border border-gray-300 rounded-lg p-2 focus:outline-none focus:border-red-500"
            value={horario}
            onChange={(e) => setHorario(e.target.value)}
          />
        </div>
        <div className="flex flex-col mb-4">
          <div className="flex justify-center">
            <label htmlFor="descripcion" className="mb-2">
              Descripción
            </label>
          </div>
          <textarea
            id="descripcion"
            className="border border-gray-300 rounded-lg p-2 focus:outline-none focus:border-red-500"
            value={descripcion}
            onChange={(e) => setDescripcion(e.target.value)}
          />
        </div>
        <div className="flex flex-col mb-4">
          <div className="flex justify-center">
            <label htmlFor="urlImagen" className="mb-2">
              URL de Imagen
            </label>
          </div>
          <input
            type="text"
            id="urlImagen"
            className="border border-gray-300 rounded-lg p-2 focus:outline-none focus:border-red-500"
            value={urlImagen}
            onChange={(e) => setUrlImagen(e.target.value)}
          />
        </div>
 
        <div className="flex justify-center">
          <button type="submit" className="px-4 py-2 rounded-lg bg-blue-500 text-white">
            Crear
          </button>
        </div>
      </form>
    </div>
  );
};

export default CrearNegocio;
