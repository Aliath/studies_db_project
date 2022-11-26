CREATE PROCEDURE delete_critical_reviews(rate_threshold integer)
LANGUAGE SQL
AS $$
    DELETE FROM episode_reviews
    WHERE
        reviewer_id IN (SELECT users.id FROM episode_reviews LEFT JOIN users ON episode_reviews.reviewer_id = users.id GROUP BY users.id HAVING AVG(rate) < rate_threshold);
$$;
