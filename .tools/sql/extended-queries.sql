/* select average duration of actor presence in episodes */
SELECT 
  actors.id as actor_id,
  actors.fullname as actor, 
  COUNT(*) as presence_count,
  ROUND(AVG(episodes.duration) / 60) as average_duration_in_minutes 
FROM actors 
  JOIN actors_episodes_associations ON actors_episodes_associations.actor_id = actors.id 
  JOIN episodes ON actors_episodes_associations.episode_id = episodes.id 
GROUP BY actors.fullname, actors.id
ORDER BY average_duration_in_minutes DESC;

/* select number of particular actors presence in series */
SELECT 
  series.name as series_name, 
  actors.fullname as actor,
  actors.id as actor_id, 
  COUNT(*) as count 
FROM series 
  JOIN seasons ON seasons.series_id = series.id 
  JOIN episodes ON episodes.season_id = seasons.id 
  JOIN actors_episodes_associations ON actors_episodes_associations.episode_id = episodes.id 
  JOIN actors ON actors_episodes_associations.actor_id = actors.id 
GROUP BY actors.fullname, series.name, actors.id
ORDER BY count DESC;

/* select reviews from given user */
SELECT
  CONCAT(series.name, ' > ', seasons.name, ' > ', episodes.name),
  episode_reviews.review,
  episode_reviews.rate,
  episodes.name,
  episode_reviews.created_at
FROM episode_reviews
  JOIN episodes ON episode_reviews.episode_id = episodes.id
  JOIN seasons ON episodes.season_id = season_id
  JOIN series ON seasons.series_id = series.id
WHERE episode_reviews.reviewer_id = 1;