package group

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (ms Members) ValidateBasic() error {
	index := make(map[string]struct{}, len(ms.Members))
	for i := range ms.Members {
		member := ms.Members[i]
		if err := member.ValidateBasic(); err != nil {
			return err
		}
		addr := member.Address
		if _, exists := index[addr]; exists {
			return sdkerrors.Wrapf(ErrDuplicate, "address: %s", addr)
		}
		index[addr] = struct{}{}
	}
	return nil
}

type AccAddresses []sdk.AccAddress

// ValidateBasic verifies that there's no duplicate address.
// Individual account address validation has to be done separately.
func (a AccAddresses) ValidateBasic() error {
	index := make(map[string]struct{}, len(a))
	for i := range a {
		accAddr := a[i]
		addr := string(accAddr)
		if _, exists := index[addr]; exists {
			return sdkerrors.Wrapf(ErrDuplicate, "address: %s", accAddr.String())
		}
		index[addr] = struct{}{}
	}
	return nil
}

// ValidateBasic verifies that check the length of options and if there is any duplication.
func (ops Options) ValidateBasic() error {
	if len(ops.Titles) == 0 {
		return sdkerrors.Wrap(ErrEmpty, "poll options")
	}

	for i, x := range ops.Titles {
		if len(x) == 0 {
			return sdkerrors.Wrapf(ErrEmpty, "option %d", i)
		}
	}

	if err := assertOptionsLength(ops, "options"); err != nil {
		return err
	}

	index := make(map[string]struct{}, len(ops.Titles))
	for _, x := range ops.Titles {
		if _, exists := index[x]; exists {
			return sdkerrors.Wrapf(ErrDuplicate, "option %s", x)
		}
		index[x] = struct{}{}
	}
	return nil
}