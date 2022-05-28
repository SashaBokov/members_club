package service

import (
	"errors"
	"fmt"
	mc "members_club"
	"regexp"
	"time"
)

func GetListOfMembers() *[]mc.Member {
	return mc.Members
}

// TODO: Refactored
//func AddMember(name, email string) error {
//	if err := checkEmail(email); err != nil {
//		return err
//	}
//
//	e, m, d := time.Now().Date()
//	date := fmt.Sprintf("%v.%v.%v", d, m, e)
//	members := append(*mc.Members,
//		mc.Member{
//			Name:  name,
//			Email: email,
//			Date:  date})
//	mc.Members = &members
//	fmt.Println(name, email, date)
//	return nil
//}

// TODO: Refactored
//func checkEmail(email string) error {
//	for _, m := range *mc.Members {
//		if m.Email == email {
//			return errors.New(fmt.Sprintf("error Email: %s alredi used", email))
//		}
//	}
//	return nil
//}

func AddMember(name, email string) error {
	if err := MatchEmail(email); err != nil {
		return err
	}
	if doseExistEmail(email) {
		return errors.New(fmt.Sprintf("error Email: %s alredi exists", email))
	}

	e, m, d := time.Now().Date()
	date := fmt.Sprintf("%v.%v.%v", d, m, e)
	members := append(*mc.Members,
		mc.Member{
			Name:  name,
			Email: email,
			Date:  date})
	mc.Members = &members
	fmt.Println(name, email, date)
	return nil
}

func doseExistEmail(email string) bool {
	for _, m := range *mc.Members {
		if m.Email == email {
			return true
		}
	}

	return false
}

// TODO: New code
func RemoveMember(email string) error {
	if err := MatchEmail(email); err != nil {
		return err
	}
	if !doseExistEmail(email) {
		return errors.New(fmt.Sprintf("error Email: %s not exists", email))
	}

	var index int
	for i, m := range *mc.Members {
		if m.Email == email {
			index = i
			break
		}
	}

	members := append((*mc.Members)[:index], (*mc.Members)[index+1:]...)
	mc.Members = &members
	return nil
}

// TODO: Was H/W
func MatchEmail(email string) error {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if emailRegex.MatchString(email) {
		return errors.New("invalid Email")
	}
	return nil
}
