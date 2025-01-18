import MovieList from "@/components/movielist";

const TopRated = () => {
    return (
        <div>
            <h1 className="flex justify-center items-center text-4xl font-bold pb-10">Top Rated Movies</h1>
            <MovieList/>
        </div>
    )
}

export default TopRated;