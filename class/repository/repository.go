package repository

import "gorm.io/gorm"

type Repository struct {
	Voucher       VoucherRepository
	RedeemVoucher RedeemVoucherRepository
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		Voucher:       *NewVoucherRepository(db),
		RedeemVoucher: *NewRedeemVoucherRepository(db),
	}
}
