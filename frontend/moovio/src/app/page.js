import MovieList from "@/components/movielist";

const home = () => {
  return (
    <div className="flex flex-col justify-center items-center gap-2">
      <div className="bg-red-800 md:w-7/12 lg:w-9/12 xl:w-5/12 w-full rounded-lg">
        <h1 className="flex justify-center items-center text-lg">Latest Movies</h1>
      </div>
      
      <div className="flex justify-center md:w-10/12 lg:w-9/12 xl:w-5/12 w-full">
        <MovieList/>
      </div>
    </div>
    
  )
}

export default home;