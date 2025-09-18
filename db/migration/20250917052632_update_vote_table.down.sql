ALTER TABLE question_votes DROP CONSTRAINT fk_reply_votes_users;

ALTER TABLE question_votes DROP user_id;

ALTER TABLE answer_votes DROP CONSTRAINT fk_answer_votes_users;

ALTER TABLE answer_votes DROP user_id;

ALTER TABLE question_votes DROP CONSTRAINT fk_question_votes_users;

ALTER TABLE question_votes DROP user_id;