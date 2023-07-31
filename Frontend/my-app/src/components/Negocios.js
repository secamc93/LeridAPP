import React, { useState, useEffect } from 'react';
import axios from 'axios';
import ConsultarNegocios from './ConsultarNegocio';
import CrearNegocio from './CrearNegocio';
import Categoria from './Categoria';
import AsignarCategorias from './AsignarCategorias';


const Negocios = () => {
    const [negocios, setNegocios] = useState([]);

    useEffect(() => {
        const fetchNegocios = async () => {
            try {
                const response = await axios.get('http://localhost:8000/api/negocios');
                setNegocios(response.data);
            } catch (error) {
                console.log(error);
            }
        };

        fetchNegocios();
    }, []);

    return (
        <div className="flex flex-row justify-between p-8">
            <div className="flex flex-col items-start space-y-4">
                <div className="mx-4">
                    <Categoria />
                </div>
                <div className="mx-4">
                    <CrearNegocio setNegocios={setNegocios} />
                </div>
            </div>
            <div className="flex flex-row mt-8">
                <div className="mx-4">
                    <ConsultarNegocios negocios={negocios}  setNegocios={setNegocios}/>
                    <AsignarCategorias negocios={negocios}  setNegocios={setNegocios} />
                </div>
            </div>
        </div>
    );
    
};

export default Negocios;
