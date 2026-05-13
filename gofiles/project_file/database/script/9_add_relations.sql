/*
Tengo que agregar a la tabla  para poder agregar relaciones
chat_tags.chat_created_at

reactions.chat_created_at

*/

ALTER TABLE public.chat_tags ADD CONSTRAINT fk_chats_chat_tags FOREIGN KEY (chat_id, chat_created_at) REFERENCES 
chats(id, created_at) ON UPDATE CASCADE ON DELETE cascade;

ALTER TABLE public.reactions ADD CONSTRAINT fk_chats_reactions FOREIGN KEY (chat_id, chat_created_at) REFERENCES 
chats(id, created_at) ON UPDATE CASCADE ON DELETE cascade;

-- En whatsameow
ALTER TABLE public.reactions
ADD CONSTRAINT fk_reactions_chat_id
FOREIGN KEY (chat_id)
REFERENCES chats(id)
ON UPDATE CASCADE
ON DELETE CASCADE
NOT VALID;

-- If want valid do it:
ALTER TABLE public.reactions VALIDATE CONSTRAINT fk_reactions_chat_id;
