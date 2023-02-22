package repository

import "gorm.io/gorm"

type CrmRepositoryDB struct {
	CrmCollection *gorm.DB
}

func NewCompaniesRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}

func NewDealsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewFeedbackSubmissionsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewContactsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewTicketsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewEmailsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewQuotesRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewLineItemsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewMeetingsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewProductsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewTasksRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewCallsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewCommunicationsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewNotesRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewPostalMailRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
