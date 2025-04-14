-- name: CreateFeedFollow :one
with insert_feed_follows AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT insert_feed_follows.*, users.name AS user_name, feeds.name AS feed_name FROM insert_feed_follows
INNER JOIN users ON user_id = users.id
INNER JOIN feeds ON feed_id = feeds.id; 

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.* , users.name AS user_name, feeds.name AS feed_name FROM feed_follows 
INNER JOIN users ON users.id = user_id
INNER JOIN feeds ON feeds.id = feed_id
WHERE feed_follows.user_id =  $1;

-- name: RemoveFollowByUserURL :exec
DELETE FROM feed_follows WHERE feed_follows.user_id = $1 AND feed_follows.feed_id = (
    SELECT id FROM feeds WHERE url = $2
);
