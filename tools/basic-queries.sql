/* insert a new genre */
INSERT INTO genres(name, description) VALUES(
    'horror',
    'Literature, film, and television that is meant to scare, startle, shock, and even repulse audiences.'
)

/* query that fails - inserting a series without proper genre_id */
INSERT INTO series(name, description, genre_id) VALUES(
    'Stranger Things',
    'In 1980s Indiana, a group of young friends witness supernatural forces and secret government exploits. As they search for answers, the children unravel a series of extraordinary mysteries.',
    404
)

/* query that doesn't fail - inserting a series with proper genre_id */
INSERT INTO series(name, description, genre_id) VALUES(
    'Stranger Things',
    'In 1980s Indiana, a group of young friends witness supernatural forces and secret government exploits. As they search for answers, the children unravel a series of extraordinary mysteries.',
    1
)

/* select number of genres in all series */

SELECT genres.name, COUNT(*) FROM series LEFT JOIN genres ON series.genre_id = genres.id GROUP BY genres.name

/* select average duration of actors presence in episodes */ 
SElECT actors.fullname, COUNT(*), ROUND(AVG(episodes.duration) / 60) as average_duration_in_minutes FROM actors
	JOIN actors_episodes_associations ON actors_episodes_associations.actor_id = actors.id
	JOIN episodes ON actors_episodes_associations.episode_id = episodes.id
GROUP BY actors.fullname
ORDER BY average_duration_in_minutes DESC

/* select number of particular actor presence in serie episodes */
SELECT series.name, actors.fullname, COUNT(*) as count FROM series
	JOIN seasons ON seasons.series_id = series.id
	JOIN episodes ON episodes.season_id = seasons.id
	JOIN actors_episodes_associations ON actors_episodes_associations.episode_id = episodes.id
	JOIN actors ON actors_episodes_associations.actor_id = actors.id
GROUP BY
	actors.fullname, series.name
ORDER BY
	count DESC

/* select min/max duration of episode by series */
SELECT
	series.name,
	(MIN(episodes.duration) / 60) as min_duration_in_minutes,
	(MAX(episodes.duration) / 60) as max_duration_in_minutes
FROM series
	JOIN seasons ON seasons.series_id = series.id
	JOIN episodes ON episodes.season_id = seasons.id
GROUP BY
	series.name