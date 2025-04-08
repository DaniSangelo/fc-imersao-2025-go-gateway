package service

import (
	"github.com/devfullcycle/imersao2025/go-gateway/internal/domain"
	"github.com/devfullcycle/imersao2025/go-gateway/internal/dto"
)

type InvoiceService struct {
	invoiceRepository domain.InvoiceRepository
	accountService    AccountService
}

func NewInvoiceService(invoiceRepository domain.InvoiceRepository, accountService AccountService) *InvoiceService {
	return &InvoiceService{
		invoiceRepository: invoiceRepository,
		accountService:    accountService,
	}
}

func (s *InvoiceService) Create(input dto.CreateInvoiceInput) (*dto.CreateInvoiceOutput, error) {
	accountOutput, err := s.accountService.FindByApiKey(input.ApiKey)
	if err != nil {
		return nil, err
	}

	invoice, err := dto.ToInvoice(input, accountOutput.ID)
	if err != nil {
		return nil, err
	}

	if err := invoice.Process(); err != nil {
		return nil, err
	}

	if invoice.Status == domain.StatusApproved {
		_, err = s.accountService.UpdateBalance(input.ApiKey, invoice.Amount)
		if err != nil {
			return nil, err
		}
	}

	if err := s.invoiceRepository.Save(invoice); err != nil {
		return nil, err
	}

	return dto.FromInvoice(invoice), nil
}

func (s *InvoiceService) GetByID(id, apiKey string) (*dto.CreateInvoiceOutput, error) {
	invoice, err := s.invoiceRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	accountOutput, err := s.accountService.FindByApiKey(apiKey)
	if err != nil {
		return nil, err
	}

	if invoice.AccountID != accountOutput.ID {
		return nil, domain.ErrUnauthorizedAccess
	}

	return dto.FromInvoice(invoice), nil
}

func (s *InvoiceService) ListByAccount(accountID string) ([]*dto.CreateInvoiceOutput, error) {
	invoices, err := s.invoiceRepository.FindByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	output := make([]*dto.CreateInvoiceOutput, len(invoices))
	for i, invoice := range invoices {
		output[i] = dto.FromInvoice(invoice)
	}
	return output, nil
}

func (s *InvoiceService) ListByAccountApiKey(apiKey string) ([]*dto.CreateInvoiceOutput, error) {
	accountOutput, err := s.accountService.FindByApiKey(apiKey)
	if err != nil {
		return nil, err
	}

	return s.ListByAccount(accountOutput.ID)
}