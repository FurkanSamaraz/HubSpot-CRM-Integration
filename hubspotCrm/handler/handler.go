package crm_handlers

import (
	"fmt"
	configs "main/config"
	handler_Companies "main/handler/Companies"
	handler_Deals "main/handler/Deals"
	crm_integration "main/package"
	repository "main/repository"
	crm_repository_Companies "main/repository/Companies"
	crm_repository_Deals "main/repository/Deals"

	handler_FeedbackSubmissions "main/handler/FeedbackSubmissions"
	crm_repository_FeedbackSubmissions "main/repository/FeedbackSubmissions"

	handler_Contacts "main/handler/Contacts"
	crm_repository_Contacts "main/repository/Contacts"

	handler_Tickets "main/handler/Tickets"
	crm_repository_Tickets "main/repository/Tickets"

	handler_Emails "main/handler/Engagement/Emails"
	crm_repository_Emails "main/repository/Engagement/Emails"

	handler_Quotes "main/handler/Quotes"
	crm_repository_Quotes "main/repository/Quotes"

	handler_Meetings "main/handler/Engagement/Meetings"
	crm_repository_Meetings "main/repository/Engagement/Meetings"

	handler_Products "main/handler/Products"
	crm_repository_Products "main/repository/Products"

	handler_Tasks "main/handler/Engagement/Tasks"
	crm_repository_Tasks "main/repository/Engagement/Tasks"

	handler_Calls "main/handler/Engagement/Calls"
	crm_repository_Calls "main/repository/Engagement/Calls"

	handler_LineItems "main/handler/LineItems"
	crm_repository_LineItems "main/repository/LineItems"

	handler_Communications "main/handler/Engagement/Communications"
	crm_repository_Communications "main/repository/Engagement/Communications"

	handler_Notes "main/handler/Engagement/Notes"
	crm_repository_Notes "main/repository/Engagement/Notes"

	handler_PostalMail "main/handler/Engagement/PostalMail"
	crm_repository_PostalMail "main/repository/Engagement/PostalMail"

	service "main/service"
)

type CRM_CONTROLLER struct {
	Base *crm_integration.CRM_BASE
}

func App() {
	dbClient := configs.GetCollection()

	CompaniesRepositoryDb := repository.NewCompaniesRepositoryDB(dbClient)
	Companies := handler_Companies.CompaniesHandler{Service: service.NewCompaniesService(crm_repository_Companies.CompaniesDB{CrmCollection: CompaniesRepositoryDb.CrmCollection})}

	DealsRepositoryDb := repository.NewDealsRepositoryDB(dbClient)
	Deals := handler_Deals.DealsHandler{Service: service.NewDealsService(crm_repository_Deals.DealsDB{CrmCollection: DealsRepositoryDb.CrmCollection})}

	FeedbackSubmissionsRepositoryDb := repository.NewFeedbackSubmissionsRepositoryDB(dbClient)
	FeedbackSubmissions := handler_FeedbackSubmissions.FeedbackSubmissionsHandler{Service: service.NewFeedbackSubmissionsService(crm_repository_FeedbackSubmissions.FeedbackSubmissionsDB{CrmCollection: FeedbackSubmissionsRepositoryDb.CrmCollection})}

	ContactsRepositoryDb := repository.NewContactsRepositoryDB(dbClient)
	Contacts := handler_Contacts.ContactsHandler{Service: service.NewContactsService(crm_repository_Contacts.ContactsDB{CrmCollection: ContactsRepositoryDb.CrmCollection})}

	TicketsRepositoryDb := repository.NewTicketsRepositoryDB(dbClient)
	Tickets := handler_Tickets.TicketsHandler{Service: service.NewTicketsService(crm_repository_Tickets.KnowledgerArticlesDB{CrmCollection: TicketsRepositoryDb.CrmCollection})}

	EmailsRepositoryDb := repository.NewEmailsRepositoryDB(dbClient)
	Emails := handler_Emails.EmailsHandler{Service: service.NewEmailsService(crm_repository_Emails.EmailsDB{CrmCollection: EmailsRepositoryDb.CrmCollection})}

	QuotesRepositorysDb := repository.NewQuotesRepositoryDB(dbClient)
	Quotes := handler_Quotes.QuotesHandler{Service: service.NewQuotesService(crm_repository_Quotes.QuotesDB{CrmCollection: QuotesRepositorysDb.CrmCollection})}

	MeetingsRepositoryDb := repository.NewMeetingsRepositoryDB(dbClient)
	Meetings := handler_Meetings.MeetingsHandler{Service: service.NewMeetingsService(crm_repository_Meetings.MeetingsDB{CrmCollection: MeetingsRepositoryDb.CrmCollection})}

	ProductsRepositoryDb := repository.NewProductsRepositoryDB(dbClient)
	Products := handler_Products.ProductsHandler{Service: service.NewProductsService(crm_repository_Products.ProductsDB{CrmCollection: ProductsRepositoryDb.CrmCollection})}

	CallsRepositoryDb := repository.NewCallsRepositoryDB(dbClient)
	Calls := handler_Calls.CallsHandler{Service: service.NewCallsService(crm_repository_Calls.CallsDB{CrmCollection: CallsRepositoryDb.CrmCollection})}

	TasksRepositoryDb := repository.NewTasksRepositoryDB(dbClient)
	Tasks := handler_Tasks.TasksHandler{Service: service.NewTasksService(crm_repository_Tasks.TasksDB{CrmCollection: TasksRepositoryDb.CrmCollection})}

	LineItemsRepositoryDb := repository.NewLineItemsRepositoryDB(dbClient)
	LineItems := handler_LineItems.LineItemsHandler{Service: service.NewLineItemsService(crm_repository_LineItems.LineItemsDB{CrmCollection: LineItemsRepositoryDb.CrmCollection})}

	PostalMailRepositoryDb := repository.NewPostalMailRepositoryDB(dbClient)
	PostalMail := handler_PostalMail.PostalMailHandler{Service: service.NewPostalMailService(crm_repository_PostalMail.PostalMailDB{CrmCollection: PostalMailRepositoryDb.CrmCollection})}

	NotesRepositoryDb := repository.NewNotesRepositoryDB(dbClient)
	Notes := handler_Notes.NotesHandler{Service: service.NewNotesService(crm_repository_Notes.NotesDB{CrmCollection: NotesRepositoryDb.CrmCollection})}

	CommunicationsRepositoryDb := repository.NewCommunicationsRepositoryDB(dbClient)
	Communications := handler_Communications.CommunicationsHandler{Service: service.NewCommunicationsService(crm_repository_Communications.CommunicationsDB{CrmCollection: CommunicationsRepositoryDb.CrmCollection})}

	fmt.Println(
		Companies,
		Deals,
		FeedbackSubmissions,
		Contacts,
		Tickets,
		Emails,
		Quotes,
		Meetings,
		Products,
		Calls,
		Tasks,
		LineItems,
		PostalMail,
		Notes,
		Communications,
	)

}
