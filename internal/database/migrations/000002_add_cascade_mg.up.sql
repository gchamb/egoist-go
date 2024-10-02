-- Modify progress_entry table
ALTER TABLE progress_entry
DROP FOREIGN KEY progress_entry_ibfk_1,
ADD FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE;

-- Modify progress_report table
ALTER TABLE progress_report
DROP FOREIGN KEY progress_report_ibfk_1,
ADD FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE;

-- Modify progress_video table
ALTER TABLE progress_video
DROP FOREIGN KEY progress_video_ibfk_1,
ADD FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE;

-- Modify revenue_cat_subscriber table
ALTER TABLE revenue_cat_subscriber
DROP FOREIGN KEY revenue_cat_subscriber_ibfk_1,
ADD FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE;
