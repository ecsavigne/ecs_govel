package migration

func (self *coreMigration) createIndexMessage() error {
	if !self.db.Migrator().HasIndex(&Message{}, "ConversationID") {
		if e := self.db.Migrator().CreateIndex(&Message{}, "ConversationID"); e != nil {
			return e
		}
	}

	if !self.db.Migrator().HasIndex(&Message{}, "TypeID") {
		if e := self.db.Migrator().CreateIndex(&Message{}, "TypeID"); e != nil {
			return e
		}
	}

	if !self.db.Migrator().HasIndex(&Message{}, "StatusID") {
		if e := self.db.Migrator().CreateIndex(&Message{}, "StatusID"); e != nil {
			return e
		}
	}

	if !self.db.Migrator().HasIndex(&Message{}, "UserID") {
		if e := self.db.Migrator().CreateIndex(&Message{}, "UserID"); e != nil {
			return e
		}
	}

	if !self.db.Migrator().HasIndex(&Message{}, "CreatedAt") {
		if e := self.db.Migrator().CreateIndex(&Message{}, "CreatedAt"); e != nil {
			return e
		}
	}

	if !self.db.Migrator().HasIndex(&Message{}, "UpdatedAt") {
		if e := self.db.Migrator().CreateIndex(&Message{}, "UpdatedAt"); e != nil {
			return e
		}
	}

	if !self.db.Migrator().HasIndex(&Message{}, "DeletedAt") {
		if e := self.db.Migrator().CreateIndex(&Message{}, "DeletedAt"); e != nil {
			return e
		}
	}

	if !self.db.Migrator().HasIndex(&Message{}, "SenderAt") {
		if e := self.db.Migrator().CreateIndex(&Message{}, "SenderAt"); e != nil {
			return e
		}
	}

	if !self.db.Migrator().HasIndex(&Message{}, "ConversationCreatedAt") {
		if e := self.db.Migrator().CreateIndex(&Message{}, "ConversationCreatedAt"); e != nil {
			return e
		}
	}

	if !self.db.Migrator().HasIndex(&Message{}, "SenderID") {
		if e := self.db.Migrator().CreateIndex(&Message{}, "SenderID"); e != nil {
			return e
		}
	}

	if !self.db.Migrator().HasIndex(&Message{}, "ExternalID") {
		if e := self.db.Migrator().CreateIndex(&Message{}, "ExternalID"); e != nil {
			return e
		}
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
