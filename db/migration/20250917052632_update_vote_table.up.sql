ALTER TABLE question_votes
    ADD user_id UUID NOT NULL;

ALTER TABLE question_votes
    ADD CONSTRAINT fk_question_votes_users
    FOREIGN KEY (user_id) REFERENCES users(id)
    ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE answer_votes
    ADD user_id UUID NOT NULL;

ALTER TABLE answer_votes
    ADD CONSTRAINT fk_answer_votes_users
    FOREIGN KEY (user_id) REFERENCES users(id)
    ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE reply_votes
    ADD user_id UUID NOT NULL;

ALTER TABLE reply_votes
    ADD CONSTRAINT fk_reply_votes_users
    FOREIGN KEY (user_id) REFERENCES users(id)
    ON DELETE CASCADE ON UPDATE CASCADE;
