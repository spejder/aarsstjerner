package main

import (
	"fmt"

	"bitbucket.org/long174/go-odoo"
	"github.com/spejder/aarsstjerner/internal/ms"
)

func profiles(client *ms.Client) (*ms.MemberProfiles, error) {
	criteria := odoo.NewCriteria().Add("can_access_contact_info", "=", true)

	criteria.
		Add("state", "=", "active").
		Add("is_active_leader", "=", false)

	options := odoo.NewOptions().FetchFields(
		"id",
		"display_name",
		"scout_name",
		"membership_ids",
		"other_info",
	)

	profiles, err := client.FindMemberProfiles(criteria, options)
	if err != nil {
		return &ms.MemberProfiles{}, fmt.Errorf("finding member profiles: %w", err)
	}

	return profiles, nil
}

func memberships(client *ms.Client, ids []int64) (*ms.MemberMemberships, error) {
	criteria := odoo.NewCriteria().Add("id", "=", ids)

	options := odoo.NewOptions().FetchFields(
		"id",
		"active_flag",
		"start_date",
		"end_date",
	)

	memberships, err := client.FindMemberMemberships(criteria, options)
	if err != nil {
		return &ms.MemberMemberships{}, fmt.Errorf("finding member profiles: %w", err)
	}

	return memberships, nil
}
