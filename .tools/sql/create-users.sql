-- create user that should be used to connect to DB from go application --
CREATE USER golang_app WITH PASSWORD 'password_to_set';

-- grant access to needed resources for created user --
GRANT ALL PRIVILEGES ON actors, actors_episodes_associations, episode_reviews, episodes, genres, seasons, series, users TO golang_app;
