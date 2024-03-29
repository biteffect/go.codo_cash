package codo

type TransferStatus string

func (s TransferStatus) String() string {
	return string(s)
}

const (
	TrStatusDraft     TransferStatus = "draft"
	TrStatusPrepared  TransferStatus = "prepared"
	TrStatusAccepted  TransferStatus = "accepted"
	TrStatusFail      TransferStatus = "fail"
	TrStatusPayable   TransferStatus = "payable"
	TrStatusHold      TransferStatus = "hold"
	TrStatusReady     TransferStatus = "ready"
	TrStatusPromised  TransferStatus = "promised"
	TrStatusDelivered TransferStatus = "delivered"
	TrStatusPaid      TransferStatus = "paid"
	TrStatusCanceled  TransferStatus = "canceled"
)

func TransferStatuses() []TransferStatus {
	return []TransferStatus{
		TrStatusDraft,
		TrStatusPrepared,
		TrStatusAccepted,
		TrStatusFail,
		TrStatusPayable,
		TrStatusHold,
		TrStatusReady,
		TrStatusPromised,
		TrStatusDelivered,
		TrStatusPaid,
		TrStatusCanceled,
	}
}
