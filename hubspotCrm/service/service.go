package crm_services

import (
	crm_repository_Companies "main/repository/Companies"
	crm_service_Companies "main/service/Companies"

	crm_repository_Deals "main/repository/Deals"
	crm_service_Deals "main/service/Deals"

	crm_repository_FeedbackSubmissions "main/repository/FeedbackSubmissions"
	crm_service_FeedbackSubmissions "main/service/FeedbackSubmissions"

	crm_repository_Contacts "main/repository/Contacts"
	crm_service_Contacts "main/service/Contacts"

	crm_repository_Tickets "main/repository/Tickets"
	crm_service_Tickets "main/service/Tickets"

	crm_repository_Emails "main/repository/Engagement/Emails"
	crm_service_Emails "main/service/Engagement/Emails"

	crm_repository_Quotes "main/repository/Quotes"
	crm_service_Quotes "main/service/Quotes"

	crm_repository_LineItems "main/repository/LineItems"
	crm_service_LineItems "main/service/LineItems"

	crm_repository_Meetings "main/repository/Engagement/Meetings"
	crm_service_Meetings "main/service/Engagement/Meetings"

	crm_repository_Products "main/repository/Products"
	crm_service_Products "main/service/Products"

	crm_repository_Tasks "main/repository/Engagement/Tasks"
	crm_service_Tasks "main/service/Engagement/Tasks"

	crm_repository_Calls "main/repository/Engagement/Calls"
	crm_service_Calls "main/service/Engagement/Calls"

	crm_repository_PostalMail "main/repository/Engagement/PostalMail"
	crm_service_PostalMail "main/service/Engagement/PostalMail"

	crm_repository_Notes "main/repository/Engagement/Notes"
	crm_service_Notes "main/service/Engagement/Notes"

	crm_repository_Communications "main/repository/Engagement/Communications"
	crm_service_Communications "main/service/Engagement/Communications"
)

func NewCompaniesService(Repo crm_repository_Companies.CompaniesRepositoryInterface) crm_service_Companies.CompaniesCRMService {
	return crm_service_Companies.CompaniesCRMService{Repo: Repo}
}

func NewDealsService(Repo crm_repository_Deals.DealsRepositoryInterface) crm_service_Deals.DealsCRMService {
	return crm_service_Deals.DealsCRMService{Repo: Repo}
}

func NewFeedbackSubmissionsService(Repo crm_repository_FeedbackSubmissions.FeedbackSubmissionsRepositoryInterface) crm_service_FeedbackSubmissions.FeedbackSubmissionsCRMService {
	return crm_service_FeedbackSubmissions.FeedbackSubmissionsCRMService{Repo: Repo}
}

func NewContactsService(Repo crm_repository_Contacts.ContactsRepositoryInterface) crm_service_Contacts.ContactsCRMService {
	return crm_service_Contacts.ContactsCRMService{Repo: Repo}
}
func NewTicketsService(Repo crm_repository_Tickets.KnowledgerArticlesRepositoryInterface) crm_service_Tickets.TicketsCRMService {
	return crm_service_Tickets.TicketsCRMService{Repo: Repo}
}
func NewEmailsService(Repo crm_repository_Emails.EmailsRepositoryInterface) crm_service_Emails.EmailsCRMService {
	return crm_service_Emails.EmailsCRMService{Repo: Repo}
}
func NewQuotesService(Repo crm_repository_Quotes.QuotesRepositoryInterface) crm_service_Quotes.QuotesCRMService {
	return crm_service_Quotes.QuotesCRMService{Repo: Repo}
}
func NewLineItemsService(Repo crm_repository_LineItems.LineItemsRepositoryInterface) crm_service_LineItems.LineItemsCRMService {
	return crm_service_LineItems.LineItemsCRMService{Repo: Repo}
}
func NewProductsService(Repo crm_repository_Products.ProductsRepositoryInterface) crm_service_Products.ProductsCRMService {
	return crm_service_Products.ProductsCRMService{Repo: Repo}
}

func NewTasksService(Repo crm_repository_Tasks.TasksRepositoryInterface) crm_service_Tasks.TasksCRMService {
	return crm_service_Tasks.TasksCRMService{Repo: Repo}
}
func NewMeetingsService(Repo crm_repository_Meetings.MeetingsRepositoryInterface) crm_service_Meetings.MeetingsCRMService {
	return crm_service_Meetings.MeetingsCRMService{Repo: Repo}
}
func NewCallsService(Repo crm_repository_Calls.CallsRepositoryInterface) crm_service_Calls.CallsCRMService {
	return crm_service_Calls.CallsCRMService{Repo: Repo}
}
func NewCommunicationsService(Repo crm_repository_Communications.CommunicationsRepositoryInterface) crm_service_Communications.CommunicationsCRMService {
	return crm_service_Communications.CommunicationsCRMService{Repo: Repo}
}
func NewNotesService(Repo crm_repository_Notes.NotesRepositoryInterface) crm_service_Notes.NotesCRMService {
	return crm_service_Notes.NotesCRMService{Repo: Repo}
}
func NewPostalMailService(Repo crm_repository_PostalMail.PostalMailRepositoryInterface) crm_service_PostalMail.PostalMailCRMService {
	return crm_service_PostalMail.PostalMailCRMService{Repo: Repo}
}
