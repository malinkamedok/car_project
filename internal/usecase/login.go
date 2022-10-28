package usecase

import "context"

type LoginUseCase struct {
	repoVendor  VendorRp
	repoCountry CountryRp
}

func NewLoginUseCase(rv VendorRp, rc CountryRp) *LoginUseCase {
	return &LoginUseCase{repoVendor: rv, repoCountry: rc}
}

var _ Login = (*LoginUseCase)(nil)

func (l *LoginUseCase) Login(ctx context.Context, typeUser string, name string) (int64, error) {
	if typeUser == "country" {
		return l.repoCountry.LoginCountry(ctx, name)
	}
	return l.repoVendor.LoginVendor(ctx, name)
}
