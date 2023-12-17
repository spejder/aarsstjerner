package ms

import (
	"fmt"

	"bitbucket.org/long174/go-odoo"
)

// MemberMembership represents member.membership model.
type MemberMembership struct {
	ActiveFlag *odoo.Bool `xmlrpc:"active_flag,omptempty"`
	EndDate    *odoo.Time `xmlrpc:"end_date,omptempty"`
	Id         *odoo.Int  `xmlrpc:"id,omptempty"`
	StartDate  *odoo.Time `xmlrpc:"start_date,omptempty"`
}

// MemberMemberships represents array of member.membership model.
type MemberMemberships []MemberMembership

// MemberMembershipModel is the odoo model name.
const MemberMembershipModel = "member.membership"

// Many2One convert MemberMembership to *Many2One.
func (mm *MemberMembership) Many2One() *odoo.Many2One {
	return odoo.NewMany2One(mm.Id.Get(), "")
}

// CreateMemberMembership creates a new member.membership model and returns its id.
func (c *Client) CreateMemberMembership(mm *MemberMembership) (int64, error) {
	return c.Create(MemberMembershipModel, mm)
}

// UpdateMemberMembership updates an existing member.membership record.
func (c *Client) UpdateMemberMembership(mm *MemberMembership) error {
	return c.UpdateMemberMemberships([]int64{mm.Id.Get()}, mm)
}

// UpdateMemberMemberships updates existing member.membership records.
// All records (represented by ids) will be updated by mp values.
func (c *Client) UpdateMemberMemberships(ids []int64, mm *MemberMembership) error {
	return c.Update(MemberMembershipModel, ids, mm)
}

// DeleteMemberMembership deletes an existing member.membership record.
func (c *Client) DeleteMemberMembership(id int64) error {
	return c.DeleteMemberMemberships([]int64{id})
}

// DeleteMemberMemberships deletes existing member.membership records.
func (c *Client) DeleteMemberMemberships(ids []int64) error {
	return c.Delete(MemberMembershipModel, ids)
}

// GetMemberMembership gets member.membership existing record.
func (c *Client) GetMemberMembership(id int64) (*MemberMembership, error) {
	mms, err := c.GetMemberMemberships([]int64{id})
	if err != nil {
		return nil, err
	}
	if mms != nil && len(*mms) > 0 {
		return &((*mms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of member.membership not found", id)
}

// GetMemberMemberships gets member.membership existing records.
func (c *Client) GetMemberMemberships(ids []int64) (*MemberMemberships, error) {
	mms := &MemberMemberships{}
	if err := c.Read(MemberMembershipModel, ids, nil, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMemberMembership finds member.membership record by querying it with criteria.
func (c *Client) FindMemberMembership(criteria *odoo.Criteria) (*MemberMembership, error) {
	mms := &MemberMemberships{}
	if err := c.SearchRead(MemberMembershipModel, criteria, odoo.NewOptions().Limit(1), mms); err != nil {
		return nil, err
	}
	if mms != nil && len(*mms) > 0 {
		return &((*mms)[0]), nil
	}
	return nil, fmt.Errorf("member.membership was not found")
}

// FindMemberMemberships finds member.membership records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMemberMemberships(criteria *odoo.Criteria, options *odoo.Options) (*MemberMemberships, error) {
	mms := &MemberMemberships{}
	if err := c.SearchRead(MemberMembershipModel, criteria, options, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMemberMembershipIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMemberMembershipIds(criteria *odoo.Criteria, options *odoo.Options) ([]int64, error) {
	ids, err := c.Search(MemberMembershipModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMemberMembershipId finds record id by querying it with criteria.
func (c *Client) FindMemberMembershipId(criteria *odoo.Criteria, options *odoo.Options) (int64, error) {
	ids, err := c.Search(MemberMembershipModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("member.membership was not found")
}
