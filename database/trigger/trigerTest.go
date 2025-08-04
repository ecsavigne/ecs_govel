package trigger

func TriggerDeleteMessageByCompanyWhatsapp() string {
	return `CREATE OR REPLACE FUNCTION delete_chats_by_company_whatsapp()
RETURNS TRIGGER AS $$
BEGIN
  DELETE FROM chats WHERE company_phone = OLD.whatsapp;
  RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER delete_message_by_company_whatsapp
AFTER DELETE ON company_whatsapps
FOR EACH ROW
EXECUTE FUNCTION delete_chats_by_company_whatsapp();`
}
