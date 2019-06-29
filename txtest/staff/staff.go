package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Staff struct {
	NameToMap map[string]*Employee
	DepSet    map[int32][]string
	CenterSet map[int32][]string
}

type Employee struct {
	Name   string
	Dep    int32
	Center int32
}

func (s *Staff) Init(filename string) error {
	s.NameToMap = make(map[string]*Employee)
	s.DepSet = make(map[int32][]string)
	s.CenterSet = make(map[int32][]string)
	f, _ := os.Open(filename)
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		line = strings.ReplaceAll(line, "\r", "")
		line = strings.ReplaceAll(line, "\n", "")
		st := strings.Split(line, " ")

		e := &Employee{Name: st[0]}
		if dep, err := strconv.Atoi(st[1]); err != nil {
			return err
		} else {
			e.Dep = int32(dep)
		}
		if center, err := strconv.Atoi(st[2]); err != nil {
			return err
		} else {
			e.Center = int32(center)
		}

		s.NameToMap[e.Name] = e

		depSet := s.DepSet[e.Dep]
		depSet = append(depSet, e.Name)
		s.DepSet[e.Dep] = depSet

		centerSet := s.CenterSet[e.Center]
		centerSet = append(centerSet, e.Name)
		s.CenterSet[e.Center] = centerSet
	}
	return nil
}
func (s *Staff) GetMemberInfo(name string) (int32, int32) {
	var dep, center int32
	dep = s.NameToMap[name].Dep
	center = s.NameToMap[name].Center
	return dep, center
}

func (s *Staff) GetMembersOfDepart(depId int32) []string {
	return s.DepSet[depId]
}

func (s *Staff) GetMembersOfCenter(centerId int32) []string {
	return s.CenterSet[centerId]
}

func main() {
	staff := &Staff{}
	if err := staff.Init("data.txt"); err != nil {
		fmt.Printf("err 68 %v\n", err)
	}
	dep, cen := staff.GetMemberInfo("moon")
	fmt.Printf("depId and centerId of moon is : %#v %v \n", dep, cen)

	fmt.Printf("members of department 10 : %v\n", staff.GetMembersOfDepart(10))
	fmt.Printf("members of center 108 : %v\n", staff.GetMembersOfCenter(108))
}
