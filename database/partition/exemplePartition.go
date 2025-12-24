package partition

func OpenConversationPartition() string {
	return `
	CREATE TABLE IF NOT EXISTS open_conversations (
  id SERIAL,
  company_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now(),
  deleted_at TIMESTAMP,
  company_phone TEXT NOT NULL,
  contact_phone TEXT NOT NULL,
  date_open_conversation TIMESTAMP NOT NULL,
  date_close_conversation TIMESTAMP NOT NULL,
  is_close_conversation BOOLEAN DEFAULT false,
  UNIQUE (company_phone, contact_phone,company_id),
  PRIMARY KEY (company_id, id)
) PARTITION BY HASH (company_id);

-- Crear 10 particiones
DO $$
BEGIN
  FOR i IN 0..9 LOOP
    EXECUTE format('
      CREATE TABLE IF NOT EXISTS open_conversations_p%s
      PARTITION OF open_conversations
      FOR VALUES WITH (MODULUS 10, REMAINDER %s);', i, i);
  END LOOP;
END$$;
 `
}
