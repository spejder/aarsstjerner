package ms

import (
	"fmt"

	"bitbucket.org/long174/go-odoo"
)

// MemberProfile represents member.profile model.
type MemberProfile struct {
	AcceptSpecials                         *odoo.Selection `xmlrpc:"accept_specials,omptempty"`
	AccNumber                              *odoo.String    `xmlrpc:"acc_number,omptempty"`
	AccountInvoiceIds                      *odoo.Relation  `xmlrpc:"account_invoice_ids,omptempty"`
	AccountInvoiceLineIds                  *odoo.Relation  `xmlrpc:"account_invoice_line_ids,omptempty"`
	AccountInvoiceNotDraftIds              *odoo.Relation  `xmlrpc:"account_invoice_not_draft_ids,omptempty"`
	AccountInvoiceOpenIds                  *odoo.Relation  `xmlrpc:"account_invoice_open_ids,omptempty"`
	AccountTotalResiduals                  *odoo.Float     `xmlrpc:"account_total_residuals,omptempty"`
	Active                                 *odoo.Bool      `xmlrpc:"active,omptempty"`
	ActiveFunctionIds                      *odoo.Relation  `xmlrpc:"active_function_ids,omptempty"`
	ActiveFunctionsInCurrentOrganization   *odoo.Relation  `xmlrpc:"active_functions_in_current_organization,omptempty"`
	ActiveFunctionsInProfile               *odoo.Relation  `xmlrpc:"active_functions_in_profile,omptempty"`
	ActiveMembershipIds                    *odoo.Relation  `xmlrpc:"active_membership_ids,omptempty"`
	ActiveMembershipsInCurrentOrganization *odoo.Relation  `xmlrpc:"active_memberships_in_current_organization,omptempty"`
	ActiveMembershipsInProfile             *odoo.Relation  `xmlrpc:"active_memberships_in_profile,omptempty"`
	ActiveProfileForUserIds                *odoo.Relation  `xmlrpc:"active_profile_for_user_ids,omptempty"`
	ActiveProfileIds                       *odoo.Relation  `xmlrpc:"active_profile_ids,omptempty"`
	AddressCo                              *odoo.String    `xmlrpc:"address_co,omptempty"`
	Age                                    *odoo.Int       `xmlrpc:"age,omptempty"`
	AllFunctionsInCurrentOrganization      *odoo.Relation  `xmlrpc:"all_functions_in_current_organization,omptempty"`
	AllFunctionsInProfile                  *odoo.Relation  `xmlrpc:"all_functions_in_profile,omptempty"`
	AllMembershipsInCurrentOrganization    *odoo.Relation  `xmlrpc:"all_memberships_in_current_organization,omptempty"`
	AllMembershipsInProfile                *odoo.Relation  `xmlrpc:"all_memberships_in_profile,omptempty"`
	Anonymized                             *odoo.Bool      `xmlrpc:"anonymized,omptempty"`
	AnonymizeWarning                       *odoo.Bool      `xmlrpc:"anonymize_warning,omptempty"`
	BankIds                                *odoo.Relation  `xmlrpc:"bank_ids,omptempty"`
	BelongsToCompany                       *odoo.Bool      `xmlrpc:"belongs_to_company,omptempty"`
	Birthdate                              *odoo.Time      `xmlrpc:"birthdate,omptempty"`
	BirthdateShort                         *odoo.String    `xmlrpc:"birthdate_short,omptempty"`
	Bmhash                                 *odoo.String    `xmlrpc:"bmhash,omptempty"`
	BmId                                   *odoo.String    `xmlrpc:"bm_id,omptempty"`
	BmPin                                  *odoo.String    `xmlrpc:"bm_pin,omptempty"`
	BmPwdhash                              *odoo.String    `xmlrpc:"bm_pwdhash,omptempty"`
	Bmref                                  *odoo.String    `xmlrpc:"bmref,omptempty"`
	BmSsoId                                *odoo.Int       `xmlrpc:"bm_sso_id,omptempty"`
	CalendarLastNotifAck                   *odoo.Time      `xmlrpc:"calendar_last_notif_ack,omptempty"`
	CanAccessContactInfo                   *odoo.Bool      `xmlrpc:"can_access_contact_info,omptempty"`
	CanApprove                             *odoo.Bool      `xmlrpc:"can_approve,omptempty"`
	CanDelete                              *odoo.Bool      `xmlrpc:"can_delete,omptempty"`
	CanEdit                                *odoo.Bool      `xmlrpc:"can_edit,omptempty"`
	CanExpense                             *odoo.Bool      `xmlrpc:"can_expense,omptempty"`
	CanFullAll                             *odoo.Bool      `xmlrpc:"can_full_all,omptempty"`
	CanRead                                *odoo.Bool      `xmlrpc:"can_read,omptempty"`
	CategoryId                             *odoo.Relation  `xmlrpc:"category_id,omptempty"`
	ChildIds                               *odoo.Relation  `xmlrpc:"child_ids,omptempty"`
	City                                   *odoo.String    `xmlrpc:"city,omptempty"`
	CkrCheckIds                            *odoo.Relation  `xmlrpc:"ckr_check_ids,omptempty"`
	CkrPane                                *odoo.Bool      `xmlrpc:"ckr_pane,omptempty"`
	Color                                  *odoo.Int       `xmlrpc:"color,omptempty"`
	Comment                                *odoo.String    `xmlrpc:"comment,omptempty"`
	CommercialPartnerId                    *odoo.Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CompanyId                              *odoo.Many2One  `xmlrpc:"company_id,omptempty"`
	CompleteAddress                        *odoo.String    `xmlrpc:"complete_address,omptempty"`
	ContactAddress                         *odoo.String    `xmlrpc:"contact_address,omptempty"`
	ContextAge                             *odoo.Int       `xmlrpc:"context_age,omptempty"`
	ContractIds                            *odoo.Relation  `xmlrpc:"contract_ids,omptempty"`
	ContractsCount                         *odoo.Int       `xmlrpc:"contracts_count,omptempty"`
	CountryId                              *odoo.Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate                             *odoo.Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                              *odoo.Many2One  `xmlrpc:"create_uid,omptempty"`
	Credit                                 *odoo.Float     `xmlrpc:"credit,omptempty"`
	CreditLimit                            *odoo.Float     `xmlrpc:"credit_limit,omptempty"`
	CurEventId                             *odoo.Many2One  `xmlrpc:"cur_event_id,omptempty"`
	Customer                               *odoo.Bool      `xmlrpc:"customer,omptempty"`
	Date                                   *odoo.Time      `xmlrpc:"date,omptempty"`
	DateLocalization                       *odoo.Time      `xmlrpc:"date_localization,omptempty"`
	Debit                                  *odoo.Float     `xmlrpc:"debit,omptempty"`
	DebitLimit                             *odoo.Float     `xmlrpc:"debit_limit,omptempty"`
	Diseases                               *odoo.String    `xmlrpc:"diseases,omptempty"`
	DisplayName                            *odoo.String    `xmlrpc:"display_name,omptempty"`
	Ean13                                  *odoo.String    `xmlrpc:"ean13,omptempty"`
	EditBirthdate                          *odoo.Bool      `xmlrpc:"edit_birthdate,omptempty"`
	EditMemberNumber                       *odoo.Bool      `xmlrpc:"edit_member_number,omptempty"`
	EditName                               *odoo.Bool      `xmlrpc:"edit_name,omptempty"`
	EditOrganizationId                     *odoo.Bool      `xmlrpc:"edit_organization_id,omptempty"`
	EditRestrictedFields                   *odoo.Bool      `xmlrpc:"edit_restricted_fields,omptempty"`
	Email                                  *odoo.String    `xmlrpc:"email,omptempty"`
	Employee                               *odoo.Bool      `xmlrpc:"employee,omptempty"`
	EventRegistrationIds                   *odoo.Relation  `xmlrpc:"event_registration_ids,omptempty"`
	Externalid                             *odoo.String    `xmlrpc:"externalid,omptempty"`
	Fax                                    *odoo.String    `xmlrpc:"fax,omptempty"`
	Firstname                              *odoo.String    `xmlrpc:"firstname,omptempty"`
	Function                               *odoo.String    `xmlrpc:"function,omptempty"`
	FunctionIds                            *odoo.Relation  `xmlrpc:"function_ids,omptempty"`
	FunctionsText                          *odoo.String    `xmlrpc:"functions_text,omptempty"`
	Gender                                 *odoo.Selection `xmlrpc:"gender,omptempty"`
	Handicap                               *odoo.String    `xmlrpc:"handicap,omptempty"`
	HasImage                               *odoo.Bool      `xmlrpc:"has_image,omptempty"`
	HelpInfo                               *odoo.String    `xmlrpc:"help_info,omptempty"`
	Id                                     *odoo.Int       `xmlrpc:"id,omptempty"`
	Image                                  *odoo.String    `xmlrpc:"image,omptempty"`
	ImageMedium                            *odoo.String    `xmlrpc:"image_medium,omptempty"`
	ImageSmall                             *odoo.String    `xmlrpc:"image_small,omptempty"`
	ImportStatus                           *odoo.String    `xmlrpc:"import_status,omptempty"`
	InvoiceIds                             *odoo.Relation  `xmlrpc:"invoice_ids,omptempty"`
	IsActiveLeader                         *odoo.Bool      `xmlrpc:"is_active_leader,omptempty"`
	IsCompany                              *odoo.Bool      `xmlrpc:"is_company,omptempty"`
	JournalItemCount                       *odoo.Int       `xmlrpc:"journal_item_count,omptempty"`
	Lang                                   *odoo.Selection `xmlrpc:"lang,omptempty"`
	LastActiveDate                         *odoo.Time      `xmlrpc:"last_active_date,omptempty"`
	LastContactConfirm                     *odoo.Time      `xmlrpc:"last_contact_confirm,omptempty"`
	LastImport                             *odoo.Time      `xmlrpc:"last_import,omptempty"`
	LastKnownCompleteAddress               *odoo.String    `xmlrpc:"last_known_complete_address,omptempty"`
	LastKnownContactInfo                   *odoo.String    `xmlrpc:"last_known_contact_info,omptempty"`
	LastKnownEmail                         *odoo.String    `xmlrpc:"last_known_email,omptempty"`
	LastKnownPhoneCombo                    *odoo.String    `xmlrpc:"last_known_phone_combo,omptempty"`
	Lastname                               *odoo.String    `xmlrpc:"lastname,omptempty"`
	LastReconciliationDate                 *odoo.Time      `xmlrpc:"last_reconciliation_date,omptempty"`
	LastUpdate                             *odoo.Time      `xmlrpc:"__last_update,omptempty"`
	LeaderFunctionIds                      *odoo.Relation  `xmlrpc:"leader_function_ids,omptempty"`
	LocalOrgApproverIds                    *odoo.Relation  `xmlrpc:"local_org_approver_ids,omptempty"`
	LockingCkrCheckIds                     *odoo.Relation  `xmlrpc:"locking_ckr_check_ids,omptempty"`
	LocksNameIds                           *odoo.Relation  `xmlrpc:"locks_name_ids,omptempty"`
	MeetingCount                           *odoo.Int       `xmlrpc:"meeting_count,omptempty"`
	MeetingIds                             *odoo.Relation  `xmlrpc:"meeting_ids,omptempty"`
	MemberAnonymized                       *odoo.Bool      `xmlrpc:"member_anonymized,omptempty"`
	MemberId                               *odoo.Many2One  `xmlrpc:"member_id,omptempty"`
	MemberMagazineOptionId                 *odoo.Many2One  `xmlrpc:"member_magazine_option_id,omptempty"`
	MemberNumber                           *odoo.String    `xmlrpc:"member_number,omptempty"`
	MembershipIds                          *odoo.Relation  `xmlrpc:"membership_ids,omptempty"`
	MessageFollowerIds                     *odoo.Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageIds                             *odoo.Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower                      *odoo.Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageLastPost                        *odoo.Time      `xmlrpc:"message_last_post,omptempty"`
	MessageSummary                         *odoo.String    `xmlrpc:"message_summary,omptempty"`
	MessageUnread                          *odoo.Bool      `xmlrpc:"message_unread,omptempty"`
	Mobile                                 *odoo.String    `xmlrpc:"mobile,omptempty"`
	MobileClean                            *odoo.String    `xmlrpc:"mobile_clean,omptempty"`
	MunicipalityId                         *odoo.Many2One  `xmlrpc:"municipality_id,omptempty"`
	Name                                   *odoo.String    `xmlrpc:"name,omptempty"`
	NotifyEmail                            *odoo.Selection `xmlrpc:"notify_email,omptempty"`
	OpportunityCount                       *odoo.Int       `xmlrpc:"opportunity_count,omptempty"`
	OpportunityIds                         *odoo.Relation  `xmlrpc:"opportunity_ids,omptempty"`
	OptOut                                 *odoo.Bool      `xmlrpc:"opt_out,omptempty"`
	OrganizationId                         *odoo.Many2One  `xmlrpc:"organization_id,omptempty"`
	OrganizationStructureParentId          *odoo.Many2One  `xmlrpc:"organization_structure_parent_id,omptempty"`
	OrganizationTypeId                     *odoo.Many2One  `xmlrpc:"organization_type_id,omptempty"`
	OtherInfo                              *odoo.String    `xmlrpc:"other_info,omptempty"`
	OwnEventRegistrationIds                *odoo.Relation  `xmlrpc:"own_event_registration_ids,omptempty"`
	ParentId                               *odoo.Many2One  `xmlrpc:"parent_id,omptempty"`
	ParentName                             *odoo.String    `xmlrpc:"parent_name,omptempty"`
	ParishId                               *odoo.Many2One  `xmlrpc:"parish_id,omptempty"`
	PartnerId                              *odoo.Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerLatitude                        *odoo.Float     `xmlrpc:"partner_latitude,omptempty"`
	PartnerLongitude                       *odoo.Float     `xmlrpc:"partner_longitude,omptempty"`
	PartnerPayerId                         *odoo.Many2One  `xmlrpc:"partner_payer_id,omptempty"`
	PayerForProfileIds                     *odoo.Relation  `xmlrpc:"payer_for_profile_ids,omptempty"`
	PayerForProfileThisOrganizationIds     *odoo.Relation  `xmlrpc:"payer_for_profile_this_organization_ids,omptempty"`
	Pbmhash                                *odoo.String    `xmlrpc:"pbmhash,omptempty"`
	Pbmref                                 *odoo.String    `xmlrpc:"pbmref,omptempty"`
	PermissionPhoto                        *odoo.Selection `xmlrpc:"permission_photo,omptempty"`
	Phone                                  *odoo.String    `xmlrpc:"phone,omptempty"`
	PhonecallCount                         *odoo.Int       `xmlrpc:"phonecall_count,omptempty"`
	PhonecallIds                           *odoo.Relation  `xmlrpc:"phonecall_ids,omptempty"`
	PhoneCombo                             *odoo.String    `xmlrpc:"phone_combo,omptempty"`
	PlastImport                            *odoo.Time      `xmlrpc:"plast_import,omptempty"`
	PreliminaryOrganizationId              *odoo.Many2One  `xmlrpc:"preliminary_organization_id,omptempty"`
	PrimaryMembershipOrganizationId        *odoo.Many2One  `xmlrpc:"primary_membership_organization_id,omptempty"`
	ProfileAnonymized                      *odoo.Bool      `xmlrpc:"profile_anonymized,omptempty"`
	ProfileIds                             *odoo.Relation  `xmlrpc:"profile_ids,omptempty"`
	PromotionIds                           *odoo.Relation  `xmlrpc:"promotion_ids,omptempty"`
	PropertyAccountPayable                 *odoo.Many2One  `xmlrpc:"property_account_payable,omptempty"`
	PropertyAccountPosition                *odoo.Many2One  `xmlrpc:"property_account_position,omptempty"`
	PropertyAccountReceivable              *odoo.Many2One  `xmlrpc:"property_account_receivable,omptempty"`
	PropertyPaymentTerm                    *odoo.Many2One  `xmlrpc:"property_payment_term,omptempty"`
	PropertyProductPricelist               *odoo.Many2One  `xmlrpc:"property_product_pricelist,omptempty"`
	PropertyStockCustomer                  *odoo.Many2One  `xmlrpc:"property_stock_customer,omptempty"`
	PropertyStockSupplier                  *odoo.Many2One  `xmlrpc:"property_stock_supplier,omptempty"`
	PropertySupplierPaymentTerm            *odoo.Many2One  `xmlrpc:"property_supplier_payment_term,omptempty"`
	Ref                                    *odoo.String    `xmlrpc:"ref,omptempty"`
	RefCompanies                           *odoo.Relation  `xmlrpc:"ref_companies,omptempty"`
	Registered                             *odoo.Bool      `xmlrpc:"registered,omptempty"`
	RegNumber                              *odoo.String    `xmlrpc:"reg_number,omptempty"`
	RelationAllIds                         *odoo.Relation  `xmlrpc:"relation_all_ids,omptempty"`
	RelationAllMemberIds                   *odoo.Relation  `xmlrpc:"relation_all_member_ids,omptempty"`
	RelationCount                          *odoo.Int       `xmlrpc:"relation_count,omptempty"`
	RelationIds                            *odoo.Relation  `xmlrpc:"relation_ids,omptempty"`
	RelationPartnerList                    *odoo.String    `xmlrpc:"relation_partner_list,omptempty"`
	RelationPrimaryMemberIds               *odoo.Relation  `xmlrpc:"relation_primary_member_ids,omptempty"`
	RelativeForProfileId                   *odoo.Many2One  `xmlrpc:"relative_for_profile_id,omptempty"`
	RelativeMemberId                       *odoo.Many2One  `xmlrpc:"relative_member_id,omptempty"`
	RelativePartnerId                      *odoo.Many2One  `xmlrpc:"relative_partner_id,omptempty"`
	RelativeTypeId                         *odoo.Many2One  `xmlrpc:"relative_type_id,omptempty"`
	SaleOrderCount                         *odoo.Int       `xmlrpc:"sale_order_count,omptempty"`
	SaleOrderIds                           *odoo.Relation  `xmlrpc:"sale_order_ids,omptempty"`
	School                                 *odoo.String    `xmlrpc:"school,omptempty"`
	SchoolClassLetter                      *odoo.String    `xmlrpc:"school_class_letter,omptempty"`
	SchoolClassNumber                      *odoo.String    `xmlrpc:"school_class_number,omptempty"`
	SchoolStartYear                        *odoo.Int       `xmlrpc:"school_start_year,omptempty"`
	ScoutName                              *odoo.String    `xmlrpc:"scout_name,omptempty"`
	SearchRelationDate                     *odoo.Time      `xmlrpc:"search_relation_date,omptempty"`
	SearchRelationId                       *odoo.Many2One  `xmlrpc:"search_relation_id,omptempty"`
	SearchRelationPartnerCategoryId        *odoo.Many2One  `xmlrpc:"search_relation_partner_category_id,omptempty"`
	SearchRelationPartnerId                *odoo.Many2One  `xmlrpc:"search_relation_partner_id,omptempty"`
	SectionId                              *odoo.Many2One  `xmlrpc:"section_id,omptempty"`
	Self                                   *odoo.Many2One  `xmlrpc:"self,omptempty"`
	SelfRelationPartnerList                *odoo.String    `xmlrpc:"self_relation_partner_list,omptempty"`
	SignupExpiration                       *odoo.Time      `xmlrpc:"signup_expiration,omptempty"`
	SignupToken                            *odoo.String    `xmlrpc:"signup_token,omptempty"`
	SignupType                             *odoo.String    `xmlrpc:"signup_type,omptempty"`
	SignupUrl                              *odoo.String    `xmlrpc:"signup_url,omptempty"`
	SignupValid                            *odoo.Bool      `xmlrpc:"signup_valid,omptempty"`
	Speaker                                *odoo.Bool      `xmlrpc:"speaker,omptempty"`
	SpecialConsiderations                  *odoo.String    `xmlrpc:"special_considerations,omptempty"`
	State                                  *odoo.Selection `xmlrpc:"state,omptempty"`
	StateId                                *odoo.Many2One  `xmlrpc:"state_id,omptempty"`
	StoredPartnerId                        *odoo.Many2One  `xmlrpc:"stored_partner_id,omptempty"`
	Street                                 *odoo.String    `xmlrpc:"street,omptempty"`
	Street2                                *odoo.String    `xmlrpc:"street2,omptempty"`
	StreetFloor                            *odoo.String    `xmlrpc:"street_floor,omptempty"`
	StreetLetter                           *odoo.String    `xmlrpc:"street_letter,omptempty"`
	StreetName                             *odoo.String    `xmlrpc:"street_name,omptempty"`
	StreetNumber                           *odoo.String    `xmlrpc:"street_number,omptempty"`
	StreetPlacement                        *odoo.String    `xmlrpc:"street_placement,omptempty"`
	SubscriptionCard                       *odoo.String    `xmlrpc:"subscription_card,omptempty"`
	SubscriptionFeeDateEnd                 *odoo.Time      `xmlrpc:"subscription_fee_date_end,omptempty"`
	SubscriptionFeeDateStart               *odoo.Time      `xmlrpc:"subscription_fee_date_start,omptempty"`
	SubscriptionFeeHasDraftLines           *odoo.Bool      `xmlrpc:"subscription_fee_has_draft_lines,omptempty"`
	SubscriptionFeeReceivable              *odoo.Float     `xmlrpc:"subscription_fee_receivable,omptempty"`
	SubscriptionLastChargedEndDate         *odoo.Time      `xmlrpc:"subscription_last_charged_end_date,omptempty"`
	SubscriptionLastWarningDate            *odoo.Time      `xmlrpc:"subscription_last_warning_date,omptempty"`
	SubscriptionProductId                  *odoo.Many2One  `xmlrpc:"subscription_product_id,omptempty"`
	SubscriptionTransaction                *odoo.Many2One  `xmlrpc:"subscription_transaction,omptempty"`
	Supplier                               *odoo.Bool      `xmlrpc:"supplier,omptempty"`
	TaskCount                              *odoo.Int       `xmlrpc:"task_count,omptempty"`
	TaskIds                                *odoo.Relation  `xmlrpc:"task_ids,omptempty"`
	Title                                  *odoo.Many2One  `xmlrpc:"title,omptempty"`
	TotalInvoiced                          *odoo.Float     `xmlrpc:"total_invoiced,omptempty"`
	Type                                   *odoo.Selection `xmlrpc:"type,omptempty"`
	Tz                                     *odoo.Selection `xmlrpc:"tz,omptempty"`
	TzOffset                               *odoo.String    `xmlrpc:"tz_offset,omptempty"`
	UseParentAddress                       *odoo.Bool      `xmlrpc:"use_parent_address,omptempty"`
	UserFullIds                            *odoo.Relation  `xmlrpc:"user_full_ids,omptempty"`
	UserId                                 *odoo.Many2One  `xmlrpc:"user_id,omptempty"`
	UserIds                                *odoo.Relation  `xmlrpc:"user_ids,omptempty"`
	UserReadIds                            *odoo.Relation  `xmlrpc:"user_read_ids,omptempty"`
	UserReadLimitedIds                     *odoo.Relation  `xmlrpc:"user_read_limited_ids,omptempty"`
	Vat                                    *odoo.String    `xmlrpc:"vat,omptempty"`
	VatSubjected                           *odoo.Bool      `xmlrpc:"vat_subjected,omptempty"`
	WaitinglistDate                        *odoo.Time      `xmlrpc:"waitinglist_date,omptempty"`
	Website                                *odoo.String    `xmlrpc:"website,omptempty"`
	WebsiteDescription                     *odoo.String    `xmlrpc:"website_description,omptempty"`
	WebsiteMessageIds                      *odoo.Relation  `xmlrpc:"website_message_ids,omptempty"`
	WebsiteMetaDescription                 *odoo.String    `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords                    *odoo.String    `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaTitle                       *odoo.String    `xmlrpc:"website_meta_title,omptempty"`
	WebsitePublished                       *odoo.Bool      `xmlrpc:"website_published,omptempty"`
	WebsiteShortDescription                *odoo.String    `xmlrpc:"website_short_description,omptempty"`
	WriteDate                              *odoo.Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                               *odoo.Many2One  `xmlrpc:"write_uid,omptempty"`
	Zip                                    *odoo.String    `xmlrpc:"zip,omptempty"`
}

// MemberProfiles represents array of member.profile model.
type MemberProfiles []MemberProfile

// MemberProfileModel is the odoo model name.
const MemberProfileModel = "member.profile"

// Many2One convert MemberProfile to *Many2One.
func (mp *MemberProfile) Many2One() *odoo.Many2One {
	return odoo.NewMany2One(mp.Id.Get(), "")
}

// CreateMemberProfile creates a new member.profile model and returns its id.
func (c *Client) CreateMemberProfile(mp *MemberProfile) (int64, error) {
	return c.Create(MemberProfileModel, mp)
}

// UpdateMemberProfile updates an existing member.profile record.
func (c *Client) UpdateMemberProfile(mp *MemberProfile) error {
	return c.UpdateMemberProfiles([]int64{mp.Id.Get()}, mp)
}

// UpdateMemberProfiles updates existing member.profile records.
// All records (represented by ids) will be updated by mp values.
func (c *Client) UpdateMemberProfiles(ids []int64, mp *MemberProfile) error {
	return c.Update(MemberProfileModel, ids, mp)
}

// DeleteMemberProfile deletes an existing member.profile record.
func (c *Client) DeleteMemberProfile(id int64) error {
	return c.DeleteMemberProfiles([]int64{id})
}

// DeleteMemberProfiles deletes existing member.profile records.
func (c *Client) DeleteMemberProfiles(ids []int64) error {
	return c.Delete(MemberProfileModel, ids)
}

// GetMemberProfile gets member.profile existing record.
func (c *Client) GetMemberProfile(id int64) (*MemberProfile, error) {
	mps, err := c.GetMemberProfiles([]int64{id})
	if err != nil {
		return nil, err
	}
	if mps != nil && len(*mps) > 0 {
		return &((*mps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of member.profile not found", id)
}

// GetMemberProfiles gets member.profile existing records.
func (c *Client) GetMemberProfiles(ids []int64) (*MemberProfiles, error) {
	mps := &MemberProfiles{}
	if err := c.Read(MemberProfileModel, ids, nil, mps); err != nil {
		return nil, err
	}
	return mps, nil
}

// FindMemberProfile finds member.profile record by querying it with criteria.
func (c *Client) FindMemberProfile(criteria *odoo.Criteria) (*MemberProfile, error) {
	mps := &MemberProfiles{}
	if err := c.SearchRead(MemberProfileModel, criteria, odoo.NewOptions().Limit(1), mps); err != nil {
		return nil, err
	}
	if mps != nil && len(*mps) > 0 {
		return &((*mps)[0]), nil
	}
	return nil, fmt.Errorf("member.profile was not found")
}

// FindMemberProfiles finds member.profile records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMemberProfiles(criteria *odoo.Criteria, options *odoo.Options) (*MemberProfiles, error) {
	mps := &MemberProfiles{}
	if err := c.SearchRead(MemberProfileModel, criteria, options, mps); err != nil {
		return nil, err
	}
	return mps, nil
}

// FindMemberProfileIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMemberProfileIds(criteria *odoo.Criteria, options *odoo.Options) ([]int64, error) {
	ids, err := c.Search(MemberProfileModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMemberProfileId finds record id by querying it with criteria.
func (c *Client) FindMemberProfileId(criteria *odoo.Criteria, options *odoo.Options) (int64, error) {
	ids, err := c.Search(MemberProfileModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("member.profile was not found")
}
