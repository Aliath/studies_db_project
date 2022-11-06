/* select series by their episode rates */

CREATE VIEW series_rates AS
    SELECT series.name as series_name, ROUND(AVG(episode_reviews.rate), 1) as avg_rate
    FROM series
        JOIN seasons ON series.id = seasons.series_id
        JOIN episodes ON seasons.id = episodes.season_id
        JOIN episode_reviews ON episodes.id = episode_reviews.episode_id
    GROUP BY series.name
    ORDER BY avg_rate DESC;

/* get number of reviews & avg rate per user */
CREATE VIEW user_ratings AS
    SELECT users.full_name, COUNT(*) as count, ROUND(AVG(episode_reviews.rate), 1) as avg_rate FROM episode_reviews
        JOIN users ON episode_reviews.reviewer_id = users.id
    GROUP BY users.full_name
    ORDER BY count DESC;