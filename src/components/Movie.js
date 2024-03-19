import { useEffect, useState } from "react";

import { useParams } from "react-router-dom";

const Movie = () => {
  // Define a state of a movie
  const [movie, setMovie] = useState({});
  // Get the id from the URL
  let { id } = useParams();

  // Create a fake API for testing purpose
  useEffect(() => {
    let myMovie = {
      id: 1,
      title: "Highlander",
      release_date: "1986-03-07",
      runtime: 116,
      mpaa_rating: "R",
      description: "Some long description",
    };
    setMovie(myMovie);
  }, [id]);

  return (
    <div>
      <h2>Movie: {movie.title}</h2>
      <small>
        <em>
          {movie.release_date}, {movie.runtime} minutes, Rated{" "}
          {movie.mpaa_rating}
        </em>
      </small>
      <hr />
      <p>{movie.description}</p>
    </div>
  );
};

export default Movie;
