package main

import(
	"testing"
	"github.com/golang/mock/gomock"
	"log"
)

func TestGetBMI(t *testing.T){
	tCases := []struct{
		user user
		expected string
	}{
		{user{"user1",1.75,70},"user1's BMI is 22.9"},
		{user{"user2",1.66,52},"user2's BMI is 18.9"},
	}
	for _,tc := range tCases{
		acutual := tc.user.getBMI()
		if acutual != tc.expected{
			t.Errorf("expected:%v ,acutual:%v",tc.expected,acutual)
		}
	}
}

func TestGetBMI2(t *testing.T){
	tCases := []struct{
		username string
		userBmi float64
		expected string
	}{
		{"user1",18.477,"user1's BMI is 18.5"},
		{"user2",18.922,"user2's BMI is 18.9"},
	}

	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

	for _ , tc := range tCases{
		calclator := NewMockICalculator(ctrl)
    	calclator.EXPECT().calcBMI(gomock.Any(),gomock.Any()).Return(tc.userBmi)
		user := user{
			name:tc.username,
			height:0,
			weight:0,
		}
		acutual := user.getBMI2(calclator)
		if acutual != tc.expected{
			t.Errorf("expected:%v ,acutual:%v",tc.expected,acutual)
		}
		log.Println(acutual)
	}
}