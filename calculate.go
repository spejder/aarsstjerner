package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/spejder/aarsstjerner/internal/ms"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

//nolint:cyclop
func calculate(ctx *cli.Context, profile ms.MemberProfile, membershipsMap map[int64]ms.MemberMembership) (int, string) {
	if len(profile.MembershipIds.Get()) == 0 {
		return 0, ""
	}

	active := false
	seconds := 0

	for _, msid := range profile.MembershipIds.Get() {
		memberships := membershipsMap[msid]

		active = active || memberships.ActiveFlag.Get()
		start := memberships.StartDate.Get().Unix()
		end := memberships.EndDate.Get().Unix()

		if end < 0 {
			end = time.Now().Unix()
		}

		seconds += int(end - start)
	}

	//nolint:gomnd
	years := ((seconds / 86400) + ctx.Int("slack")) / 365

	if !active {
		return 0, ""
	}

	if years < 1 || years > 10 {
		return 0, ""
	}

	aarstjerne := ""
	info := otherInfo{}

	err := yaml.Unmarshal([]byte(profile.OtherInfo.Get()), &info)
	if err == nil && !ctx.Bool("all") && info.Aarstjerne == years {
		return 0, ""
	}

	if err == nil && info.Aarstjerne > 0 {
		aarstjerne = fmt.Sprintf(" (har %d-årsstjerne)", info.Aarstjerne)
	}

	name := profile.DisplayName.Get()
	if scoutName := profile.ScoutName.Get(); scoutName != "" {
		name += fmt.Sprintf(", \"%s\"", scoutName)
	}

	if ctx.Bool("fake-names") {
		faker := gofakeit.New(profile.Id.Get())
		name = fmt.Sprintf("%s, \"%s\"", faker.Name(), faker.PetName())
	}

	return years, fmt.Sprintf("%s%s", name, aarstjerne)
}
