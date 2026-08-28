package migration

func (self *coreMigration) createIndexMessage() error {
	if !self.db.Migrator().HasIndex(&Message{}, "ConversationID") {
		return self.db.Migrator().CreateIndex(&Message{}, "ConversationID")
	}

	if !self.db.Migrator().HasIndex(&Message{}, "TypeID") {
		return self.db.Migrator().CreateIndex(&Message{}, "TypeID")
	}

	if !self.db.Migrator().HasIndex(&Message{}, "StatusID") {
		return self.db.Migrator().CreateIndex(&Message{}, "StatusID")
	}

	if !self.db.Migrator().HasIndex(&Message{}, "UserID") {
		return self.db.Migrator().CreateIndex(&Message{}, "UserID")
	}

	if !self.db.Migrator().HasIndex(&Message{}, "CreatedAt") {
		return self.db.Migrator().CreateIndex(&Message{}, "CreatedAt")
	}

	if !self.db.Migrator().HasIndex(&Message{}, "UpdatedAt") {
		return self.db.Migrator().CreateIndex(&Message{}, "UpdatedAt")
	}

	if !self.db.Migrator().HasIndex(&Message{}, "DeletedAt") {
		return self.db.Migrator().CreateIndex(&Message{}, "DeletedAt")
	}

	if !self.db.Migrator().HasIndex(&Message{}, "SenderAt") {
		return self.db.Migrator().CreateIndex(&Message{}, "SenderAt")
	}

	if !self.db.Migrator().HasIndex(&Message{}, "ConversationCreatedAt") {
		return self.db.Migrator().CreateIndex(&Message{}, "ConversationCreatedAt")
	}

	if !self.db.Migrator().HasIndex(&Message{}, "SenderID") {
		return self.db.Migrator().CreateIndex(&Message{}, "SenderID")
	}

	if !self.db.Migrator().HasIndex(&Message{}, "ExternalID") {
		return self.db.Migrator().CreateIndex(&Message{}, "ExternalID")
	}

	return nil
}

func (self *coreMigration) AddConstraintMessage() error {
	var err error
	if !self.db.Migrator().HasConstraint(&Message{}, "ReplyToMessage") {
		err = self.db.Migrator().CreateConstraint(&Message{}, "ReplyToMessage")
		if err != nil {
			return err
		}
	}

	// other constraints only some here

	return self.createIndexMessage()
}
