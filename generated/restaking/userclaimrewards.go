// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UserClaimRewards is the `user_claim_rewards` instruction.
type UserClaimRewards struct {
	RewardPoolId *uint8
	RewardId     *uint8

	// [0] = [WRITE, SIGNER] user
	//
	// [1] = [] system_program
	//
	// [2] = [] receipt_token_mint
	//
	// [3] = [WRITE] reward_account
	//
	// [4] = [WRITE] user_reward_account
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUserClaimRewardsInstructionBuilder creates a new `UserClaimRewards` instruction builder.
func NewUserClaimRewardsInstructionBuilder() *UserClaimRewards {
	nd := &UserClaimRewards{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	nd.AccountMetaSlice[1] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	nd.AccountMetaSlice[2] = ag_solanago.Meta(Addresses["FRAGSEthVFL7fdqM8hxfxkfCZzUvmg21cqPJVvC1qdbo"])
	return nd
}

// SetRewardPoolId sets the "reward_pool_id" parameter.
func (inst *UserClaimRewards) SetRewardPoolId(reward_pool_id uint8) *UserClaimRewards {
	inst.RewardPoolId = &reward_pool_id
	return inst
}

// SetRewardId sets the "reward_id" parameter.
func (inst *UserClaimRewards) SetRewardId(reward_id uint8) *UserClaimRewards {
	inst.RewardId = &reward_id
	return inst
}

// SetUserAccount sets the "user" account.
func (inst *UserClaimRewards) SetUserAccount(user ag_solanago.PublicKey) *UserClaimRewards {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(user).WRITE().SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *UserClaimRewards) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *UserClaimRewards) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UserClaimRewards {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *UserClaimRewards) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *UserClaimRewards) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *UserClaimRewards {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(receiptTokenMint)
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *UserClaimRewards) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetRewardAccountAccount sets the "reward_account" account.
func (inst *UserClaimRewards) SetRewardAccountAccount(rewardAccount ag_solanago.PublicKey) *UserClaimRewards {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(rewardAccount).WRITE()
	return inst
}

func (inst *UserClaimRewards) findFindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: reward
	seeds = append(seeds, []byte{byte(0x72), byte(0x65), byte(0x77), byte(0x61), byte(0x72), byte(0x64)})
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindRewardAccountAddressWithBumpSeed calculates RewardAccount account address with given seeds and a known bump seed.
func (inst *UserClaimRewards) FindRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindRewardAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *UserClaimRewards) MustFindRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindRewardAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindRewardAccountAddress finds RewardAccount account address with given seeds.
func (inst *UserClaimRewards) FindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindRewardAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *UserClaimRewards) MustFindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindRewardAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetRewardAccountAccount gets the "reward_account" account.
func (inst *UserClaimRewards) GetRewardAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetUserRewardAccountAccount sets the "user_reward_account" account.
func (inst *UserClaimRewards) SetUserRewardAccountAccount(userRewardAccount ag_solanago.PublicKey) *UserClaimRewards {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(userRewardAccount).WRITE()
	return inst
}

func (inst *UserClaimRewards) findFindUserRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: user_reward
	seeds = append(seeds, []byte{byte(0x75), byte(0x73), byte(0x65), byte(0x72), byte(0x5f), byte(0x72), byte(0x65), byte(0x77), byte(0x61), byte(0x72), byte(0x64)})
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())
	// path: user
	seeds = append(seeds, user.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindUserRewardAccountAddressWithBumpSeed calculates UserRewardAccount account address with given seeds and a known bump seed.
func (inst *UserClaimRewards) FindUserRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindUserRewardAccountAddress(receiptTokenMint, user, bumpSeed)
	return
}

func (inst *UserClaimRewards) MustFindUserRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserRewardAccountAddress(receiptTokenMint, user, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindUserRewardAccountAddress finds UserRewardAccount account address with given seeds.
func (inst *UserClaimRewards) FindUserRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindUserRewardAccountAddress(receiptTokenMint, user, 0)
	return
}

func (inst *UserClaimRewards) MustFindUserRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserRewardAccountAddress(receiptTokenMint, user, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetUserRewardAccountAccount gets the "user_reward_account" account.
func (inst *UserClaimRewards) GetUserRewardAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst UserClaimRewards) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UserClaimRewards,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UserClaimRewards) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UserClaimRewards) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.RewardPoolId == nil {
			return errors.New("RewardPoolId parameter is not set")
		}
		if inst.RewardId == nil {
			return errors.New("RewardId parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.RewardAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.UserRewardAccount is not set")
		}
	}
	return nil
}

func (inst *UserClaimRewards) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UserClaimRewards")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  RewardPoolId", *inst.RewardPoolId))
						paramsBranch.Child(ag_format.Param("      RewardId", *inst.RewardId))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("              user", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    system_program", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("receipt_token_mint", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("           reward_", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("      user_reward_", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj UserClaimRewards) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `RewardPoolId` param:
	err = encoder.Encode(obj.RewardPoolId)
	if err != nil {
		return err
	}
	// Serialize `RewardId` param:
	err = encoder.Encode(obj.RewardId)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UserClaimRewards) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `RewardPoolId`:
	err = decoder.Decode(&obj.RewardPoolId)
	if err != nil {
		return err
	}
	// Deserialize `RewardId`:
	err = decoder.Decode(&obj.RewardId)
	if err != nil {
		return err
	}
	return nil
}

// NewUserClaimRewardsInstruction declares a new UserClaimRewards instruction with the provided parameters and accounts.
func NewUserClaimRewardsInstruction(
	// Parameters:
	reward_pool_id uint8,
	reward_id uint8,
	// Accounts:
	user ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	rewardAccount ag_solanago.PublicKey,
	userRewardAccount ag_solanago.PublicKey) *UserClaimRewards {
	return NewUserClaimRewardsInstructionBuilder().
		SetRewardPoolId(reward_pool_id).
		SetRewardId(reward_id).
		SetUserAccount(user).
		SetSystemProgramAccount(systemProgram).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetRewardAccountAccount(rewardAccount).
		SetUserRewardAccountAccount(userRewardAccount)
}