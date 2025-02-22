'use client'
import { useState,useEffect } from "react";

const MovieList = () => {
    const [data,setData] = useState(null)
    const [error,setError] = useState(null)
    const [isLoading,setIsLoading] = useState(true)

    const apiURL = `${process.env.NEXT_PUBLIC_API_URL}/getmovielist`

    useEffect(() => {
        const fetchApi = async () => {
            try{
                const response = await fetch(apiURL);
                const jsonData = await response.json();
                setData(jsonData)
                setIsLoading(false)
            }catch (error){
                setError(error.message)
            }
            
        }
        fetchApi();
    },[]);
   
    if(isLoading) {
        return (
          <div className="flex justify-center items-center h-screen">
            <div className="loading"></div>
          </div>
        );
    }

    return (
        <div>
            <div className="grid lg:grid-cols-8 md:grid-cols-5 sm:grid-cols-4 grid-cols-3 gap-3 p-5">
                {data.data.map(movie => {
                    return (
                        <div 
                            key={movie.title} 
                            className="p-3 flex flex-col justify-end rounded-sm shadow-lg bg-cover bg-no-repeat bg-center md:h-[150px] md:w-[100px] lg:w-[95px] h-[300] transition-all hover:scale-110"
                            style={{ backgroundImage: `url(${movie.cover})` }}
                        >
                            <h3 className="font-bold text-xs pt-4 text-wrap">{movie.title} ({movie.year})</h3>
                        </div>
                    );
                })}
             </div>   
        </div>
    );
}

export default MovieList;