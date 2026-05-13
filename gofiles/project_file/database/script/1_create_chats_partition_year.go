package script

import (
	"fmt"
)

func CreateChatsPartitionYear() string {
	years := []int{2018, 2019, 2020, 2021, 2022, 2023, 2024, 2025, 2026, 2027, 2028, 2029, 2030, 2031, 2032, 2033, 2034, 2035, 2036, 2037, 2038, 2039, 2040, 2041, 2042, 2043, 2044, 2045, 2046, 2047, 2048, 2049, 2050}

	yearStr := "ARRAY["
	for i, v := range years {
		if i == len(years)-1 {
			yearStr += fmt.Sprintf("%d", v)
		} else {
			yearStr += fmt.Sprintf("%d, ", v)
		}
	}
	yearStr += "]"

	return fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS chats_new (
		id bigserial NOT NULL,
		company_whatsapp_id int8 NULL,
		user_id int8 NULL,
		status_id int2 NULL,
		type_id int2 NULL,
		company_id int8 REFERENCES companies(id) ON UPDATE CASCADE ON DELETE CASCADE,
		created_at timestamptz NULL,
		updated_at timestamptz NULL,
		deleted_at timestamptz NULL,
		status int2 NULL,
		"source" int2 NULL,
		seconds int4 NULL,
		response_message_source int2 NULL,
		response_message_type int2 NULL,
		response_message_seconds int4 NULL,
		message_id text NOT NULL,
		media_id text NULL,
		company_phone text NOT NULL,
		contact_phone text NOT NULL,
		message text NULL,
		"path" text NULL,
		client_original_name text NULL,
		whatsapp_date text NULL,
		service_date text NULL,
		contact_json text NULL,
		buttons_json text NULL,
		list_json text NULL,
		drive_id text NULL,
		id_parent_folder_drive text NULL,
		id_file_drive text NULL,
		url text NULL,
		response_message_id text NULL,
		response_media_id text NULL,
		response_message_text text NULL,
		response_message_path text NULL,
		response_message_original_name text NULL,
		response_message_jpeg_thumbnail text NULL,
		response_message_url text NULL,
		group_sender_data text NULL,
		"location" text NULL,
		resended bool NULL,
		played bool DEFAULT false NULL,
		data_json jsonb NULL,
	) PARTITION BY RANGE (created_at);

	DO $BlockDo$
		Declare
		year INT;
		years INT[] := %s;
		BEGIN
		    -- Create partitions
			FOREACH year IN ARRAY years LOOP
				EXECUTE format('
				CREATE TABLE IF NOT EXISTS chats_year_%%s 
				PARTITION OF chats_new
				FOR VALUES FROM (''%%s-01-01 00:00:00+00'') TO (''%%s-01-01 00:00:00+00'');', year, year, year+1);
			END LOOP;

			-- Create default partition
			CREATE TABLE IF NOT EXISTS chats_default PARTITION OF chats_new DEFAULT;
		END
	$BlockDo$;
	`, yearStr)
}
