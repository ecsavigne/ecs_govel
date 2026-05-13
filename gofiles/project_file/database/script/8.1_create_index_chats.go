// Execute if index not exist
package script

func CreateIndexChats() string {
	return `
		DO $BlockDo$
			BEGIN
				-- Create index
				CREATE INDEX IF NOT EXISTS chats_company_whatsapp_id_idx ON chats (company_whatsapp_id);
				CREATE INDEX IF NOT EXISTS chats_status_id_idx ON chats (status_id);
				CREATE INDEX IF NOT EXISTS chats_type_id_idx ON chats (type_id);
				CREATE INDEX IF NOT EXISTS chats_company_id_idx ON chats (company_id);
				CREATE INDEX IF NOT EXISTS chats_created_at_idx ON chats (created_at);
				CREATE INDEX IF NOT EXISTS chats_updated_at_idx ON chats (updated_at);
				CREATE INDEX IF NOT EXISTS chats_deleted_at_idx ON chats (deleted_at);
				CREATE INDEX IF NOT EXISTS chats_status_idx ON chats (status);
				CREATE INDEX IF NOT EXISTS chats_media_id_idx ON chats (media_id);
				CREATE INDEX IF NOT EXISTS chats_company_phone_contact_phone_message_id_idx ON chats (company_phone, contact_phone, message_id);
			END
		$BlockDo$
	`
}
