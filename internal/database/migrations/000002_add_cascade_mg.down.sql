ALTER TABLE progress_entry
DROP FOREIGN KEY progress_entry_user_fk,
ADD CONSTRAINT progress_entry_user_fk
FOREIGN KEY (user_id) REFERENCES user(id);

ALTER TABLE progress_report
DROP FOREIGN KEY progress_report_user_fk,
ADD CONSTRAINT progress_report_user_fk
FOREIGN KEY (user_id) REFERENCES user(id);

ALTER TABLE progress_video
DROP FOREIGN KEY progress_video_user_fk,
ADD CONSTRAINT progress_video_user_fk
FOREIGN KEY (user_id) REFERENCES user(id);

ALTER TABLE revenue_cat_subscriber
DROP FOREIGN KEY revenue_cat_subscriber_user_fk,
ADD CONSTRAINT revenue_cat_subscriber_user_fk
FOREIGN KEY (user_id) REFERENCES user(id);