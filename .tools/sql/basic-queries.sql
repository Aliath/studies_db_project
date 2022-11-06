/* insert a new genre */
INSERT INTO genres(name, description) VALUES(
  'horror',
  'Literature, film, and television that is meant to scare, startle, shock, and even repulse audiences.'
);

/* query that fails - inserting a series without proper genre_id */
INSERT INTO series(name, description, genre_id) 
VALUES (
  'Stranger Things', 'In 1980s Indiana, a group of young friends witness supernatural forces and secret government exploits. As they search for answers, the children unravel a series of extraordinary mysteries.', 
  404
);

/* query that doesn't fail - inserting a series with proper genre_id */
INSERT INTO series(name, description, genre_id) VALUES (
  'Stranger Things', 'In 1980s Indiana, a group of young friends witness supernatural forces and secret government exploits. As they search for answers, the children unravel a series of extraordinary mysteries.', 
  1
);

/* select number of genres in all series */
SELECT genres.name, COUNT(*) FROM series LEFT JOIN genres ON series.genre_id = genres.id GROUP BY genres.name;

UPDATE episode_reviews SET rate = 2 WHERE rate < 2;