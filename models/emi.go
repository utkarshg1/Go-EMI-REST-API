package models

import (
	"errors"
	"fmt"
	"math"
)

type LoanDetails struct {
	Principal      float64 `json:"principal" binding:"required"`
	Period         uint32  `json:"period" binding:"required"`
	RateOfInterest float64 `json:"rateOfInterest" binding:"required"`
}

type LoanCalcs struct {
	EMI      float64
	Interest float64
	Amount   float64
}

func New(Principal float64, Period uint32, RateOfInterest float64) (*LoanDetails, error) {
	if Principal <= 0 {
		return nil, errors.New("principal amount must be greater than zero")
	}
	if Period == 0 {
		return nil, errors.New("loan Period cannot be zero")
	}
	if RateOfInterest <= 0 || RateOfInterest > 100 {
		return nil, errors.New("rate of interest must be between 0 and 100")
	}
	return &LoanDetails{
		Principal:      Principal,
		Period:         Period,
		RateOfInterest: RateOfInterest,
	}, nil
}

func (l *LoanDetails) OutputLoanDetails() {
	fmt.Printf("Principal : %.2f INR\n", l.Principal)
	fmt.Printf("Period : %d years\n", l.Period)
	fmt.Printf("Rate of Intrest : %.2f percent per annum\n", l.RateOfInterest)
}

func (l *LoanDetails) GetEMICalcs() *LoanCalcs {
	n := float64(l.Period) * 12.0
	r := l.RateOfInterest / 1200.0
	x := math.Pow(1+r, n)
	emi := l.Principal * r * x / (x - 1)
	amt := emi * n
	interest := amt - l.Principal
	return &LoanCalcs{
		math.Round(emi),
		math.Round(interest),
		math.Round(amt),
	}
}
