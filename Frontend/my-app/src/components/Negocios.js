import React, { useState, useEffect } from 'react';
import axios from 'axios';
import ConsultarNegocios from './ConsultarNegocio';
import CrearNegocio from './CrearNegocio';

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
        <div className="flex flex-col items-center">
            <div className="flex flex-row">
            </div>
            <div className="flex flex-row mt-8">
                <div className="mx-4">
                    <CrearNegocio setNegocios={setNegocios} />
                </div>
                <div className="mx-4">
                    <ConsultarNegocios negocios={negocios} />
                </div>
            </div>
        </div>
    );
};

export default Negocios;
