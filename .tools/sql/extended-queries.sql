/* select average duration of actor presence in episodes */
SELECT 
  actors.fullname as actor, 
  COUNT(*) as presence_count,
  ROUND(AVG(episodes.duration) / 60) as average_duration_in_minutes 
FROM actors 
  JOIN actors_episodes_associations ON actors_episodes_associations.actor_id = actors.id 
  JOIN episodes ON actors_episodes_associations.episode_id = episodes.id 
GROUP BY actors.fullname 
ORDER BY average_duration_in_minutes DESC;

/* select number of particular actors presence in series */
SELECT 
  series.name as series_name, 
  actors.fullname as actor, 
  COUNT(*) as count 
FROM series 
  JOIN seasons ON seasons.series_id = series.id 
  JOIN episodes ON episodes.season_id = seasons.id 
  JOIN actors_episodes_associations ON actors_episodes_associations.episode_id = episodes.id 
  JOIN actors ON actors_episodes_associations.actor_id = actors.id 
GROUP BY actors.fullname, series.name 
ORDER BY count DESC;