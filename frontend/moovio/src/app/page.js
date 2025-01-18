import MovieList from "@/components/movielist";

const home = () => {
  return (
    <div>
      <h1 className="flex justify-center items-center text-4xl font-bold pb-10">Latest Movies</h1>
      <MovieList/>
    </div>
    
  )
}

export default home;