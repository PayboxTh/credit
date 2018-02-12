package member_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/beevik/ntp"
	"github.com/mrtomyum/mesy/member"
)

func TestSetMember(t *testing.T) {
	p := member.NewPeople("Tom", "Arnontavilas", "23523453254324")
	c1 := member.NewCompany()
	c2 := member.NewCompany()
	m := member.NewMember("0819983003", p, c1, c2)
	fmt.Println("Member:", *m)
	t.Log("ok")

}


