package main

import (
	"github.com/go-playground/validator/v10"
)

type Parcel struct {
	Width  int `validate:"required,gte=1"`
	Height int `validate:"required,gte=1"`
	Length int `validate:"required,gte=1"`
	Mass   int `validate:"required,gte=1"`
}

func (p *Parcel) IsHeavy() bool {
	return p.Mass >= 20
}

func (p *Parcel) IsBulky() bool {
	if p.IsLong() {
		return true
	}
	if p.Volume() >= 1000000 {
		return true
	}
	return false
}

func (p *Parcel) IsLong() bool {
	if p.Width >= 150 {
		return true
	}
	if p.Height >= 150 {
		return true
	}
	if p.Length >= 150 {
		return true
	}
	return false
}

func (p *Parcel) Volume() int {
	return p.Width * p.Height * p.Length
}

func (p *Parcel) IsValid() (bool, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Parcel) IsRejected() bool {
	valid, _ := p.IsValid()
	if !valid {
		return true
	}
	if p.IsBulky() && p.IsHeavy() {
		return true
	}

	return false
}

func (p *Parcel) IsSpecial() bool {
	if p.IsRejected() {
		return false
	}
	if p.IsBulky() {
		return true
	}
	if p.IsHeavy() {
		return true
	}

	return false
}

func (p *Parcel) Stack() string {
	if p.IsRejected() {
		return "REJECTED"
	}
	if p.IsSpecial() {
		return "SPECIAL"
	}

	return "STANDARD"
}
