CREATE TABLE video_votes
(
    video_id    varchar(6)    NOT NULL,
    upvote      int           NOT NULL,
    downvote    int           NOT NULL,
    created_at  timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (video_id)
)
