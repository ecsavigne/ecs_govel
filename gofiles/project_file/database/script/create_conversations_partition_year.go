package script

import (
	"fmt"
)

func CreateConversationsPartitionYear() string {
	years := []int{2026, 2027, 2028, 2029, 2030, 2031, 2032, 2033, 2034, 2035, 2036, 2037, 2038, 2039, 2040, 2041, 2042, 2043, 2044, 2045, 2046, 2047, 2048, 2049, 2050}

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
	CREATE TABLE IF NOT EXISTS conversations (
		id bigserial,
		ig_account_id bigint NOT NULL,
		user_id bigint NOT NULL,
		company_id bigint NOT NULL,
		status int2 NOT NULL,
		created_at timestamptz DEFAULT now(),
		updated_at timestamptz NULL,
		deleted_at timestamptz NULL,
		external_id text NULL,
		recipient_id text not null,
		recipient_account text NULL,
		payload jsonb NULL,
		PRIMARY KEY (id, created_at)
	) PARTITION BY RANGE (created_at);

	-- Indexes for the main table can be created in logical of golang struct only index unique
	CREATE UNIQUE INDEX IF NOT EXISTS idx_conversations_unique ON conversations (external_id, ig_account_id);
	

	DO $BlockDo$
		Declare
		year INT;
		years INT[] := %s;
		BEGIN
		    -- Create partitions
			FOREACH year IN ARRAY years LOOP
				EXECUTE format('
				CREATE TABLE IF NOT EXISTS conversations_year_%%s 
				PARTITION OF conversations
				FOR VALUES FROM (''%%s-01-01 00:00:00+00'') TO (''%%s-01-01 00:00:00+00'');', year, year, year+1);
			END LOOP;

			-- Create default partition
			CREATE TABLE IF NOT EXISTS conversations_default PARTITION OF conversations DEFAULT;
		END
	$BlockDo$;
	`, yearStr)
}
