package store

import (
	"log"

	giga "github.com/dogecoinfoundation/gigawallet/pkg"

	_ "github.com/mattn/go-sqlite3"
)

// interface guard ensures Mock implements giga.PaymentsStore
//var _ giga.Store = Mock{}

type Mock struct {
	invoices          map[giga.Address]giga.Invoice
	accounts          map[string]giga.Account
	accountsByAddress map[giga.Address]giga.Account
}

func (m Mock) MarkInvoiceAsPaid(id giga.Address) error {
	//TODO implement me
	panic("implement me")
}

func (m Mock) GetPendingInvoices() (<-chan giga.Invoice, error) {
	//TODO implement me
	log.Print("GetPendingInvoices: not implemented")
	return make(chan giga.Invoice), nil
}

// NewMock returns a giga.PaymentsStore implementor that stores orders in memory
func NewMock() Mock {
	return Mock{
		invoices:          make(map[giga.Address]giga.Invoice, 10),
		accounts:          make(map[string]giga.Account, 10),
		accountsByAddress: make(map[giga.Address]giga.Account, 10),
	}
}

func (m Mock) StoreInvoice(invoice giga.Invoice) error {
	m.invoices[invoice.ID] = invoice
	return nil
}

func (m Mock) GetInvoice(id giga.Address) (giga.Invoice, error) {
	v, ok := m.invoices[id]
	if !ok {
		return giga.Invoice{}, giga.NewErr(giga.NotFound, "invoice not found: %v", id)
	}
	return v, nil
}

func (m Mock) ListInvoices(account giga.Address, cursor int, limit int) (items []giga.Invoice, next_cursor int, err error) {
	return
}

func (m Mock) StoreAccount(account giga.Account) error {
	m.accounts[account.ForeignID] = account
	m.accountsByAddress[account.Address] = account
	return nil
}

func (m Mock) GetAccount(foreignID string) (giga.Account, error) {
	v, ok := m.accounts[foreignID]
	if !ok {
		return giga.Account{}, giga.NewErr(giga.NotFound, "account not found: %v", foreignID)
	}
	return v, nil
}

func (m Mock) GetAccountByAddress(id giga.Address) (giga.Account, error) {
	v, ok := m.accountsByAddress[id]
	if !ok {
		return giga.Account{}, giga.NewErr(giga.NotFound, "account not found: %v", id)
	}
	return v, nil
}

func (m Mock) GetAllUnreservedUTXOs(account giga.Address) ([]giga.UTXO, error) {
	return nil, nil
}

func (m Mock) Commit(updates []any) error {
	return nil
}
