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
    <div className="bg-red-500 rounded-lg shadow-lg p-6  w-64 h-300 overflow-auto">
      <h1 className="text-2xl font-bold mb-4 text-center">Crear Negocio</h1>
      <form onSubmit={handleSubmit}>
        <div className="flex flex-col mb-4">
          <input
            type="text"
            id="nombre"
            placeholder="Nombre"
            className="border border-gray-300 rounded-lg p-2 focus:outline-none focus:border-red-500"
            value={nombre}
            onChange={(e) => setNombre(e.target.value)}
          />
        </div>
        <div className="flex flex-col mb-4">
          <input
            type="text"
            id="direccion"
            placeholder="Dirección"
            className="border border-gray-300 rounded-lg p-2 focus:outline-none focus:border-red-500"
            value={direccion}
            onChange={(e) => setDireccion(e.target.value)}
          />
        </div>
        <div className="flex flex-col mb-4">
          <input
            type="text"
            id="telefono"
            placeholder="Teléfono"
            className="border border-gray-300 rounded-lg p-2 focus:outline-none focus:border-red-500"
            value={telefono}
            onChange={(e) => setTelefono(e.target.value)}
          />
        </div>
        <div className="flex flex-col mb-4">
          <input
            type="text"
            id="horario"
            placeholder="Horario"
            className="border border-gray-300 rounded-lg p-2 focus:outline-none focus:border-red-500"
            value={horario}
            onChange={(e) => setHorario(e.target.value)}
          />
        </div>
        <div className="flex flex-col mb-4">
          <textarea
            id="descripcion"
            placeholder="Descripción"
            className="border border-gray-300 rounded-lg p-2 focus:outline-none focus:border-red-500"
            value={descripcion}
            onChange={(e) => setDescripcion(e.target.value)}
          />
        </div>
        <div className="flex flex-col mb-4">
          <input
            type="text"
            id="urlImagen"
            placeholder="URL de Imagen"
            className="border border-gray-300 rounded-lg p-2 focus:outline-none focus:border-red-500"
            value={urlImagen}
            onChange={(e) => setUrlImagen(e.target.value)}
          />
        </div>

        <div className="flex justify-center">
          <button type="submit" className="px-4 py-2 rounded-lg bg-white text-red-500">
            Crear
          </button>
        </div>
      </form>
    </div>
  );

};

export default CrearNegocio;
