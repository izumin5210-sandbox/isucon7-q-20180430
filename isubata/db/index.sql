CREATE INDEX index_user_on_name ON user(name);
CREATE INDEX index_message_on_id_channel_id ON message(id, channel_id);
CREATE INDEX index_message_on_channel_id ON message(channel_id);
