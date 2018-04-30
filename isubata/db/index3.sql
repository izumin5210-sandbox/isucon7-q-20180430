ALTER TABLE message DROP INDEX index_message_on_id_channel_id;
ALTER TABLE message DROP INDEX index_message_on_channel_id;
CREATE INDEX index_message_on_channel_id_id ON message(channel_id, id);
