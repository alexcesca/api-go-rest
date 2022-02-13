package controllers

import (
	"net/http"
	"testing"
)

func TestHome(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Home(tt.args.w, tt.args.r)
		})
	}
}

func TestTodasPersonalidades(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TodasPersonalidades(tt.args.w, tt.args.r)
		})
	}
}

func TestRetornaUmaPersonalidade(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RetornaUmaPersonalidade(tt.args.w, tt.args.r)
		})
	}
}

func TestCriarPersonalidade(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CriarPersonalidade(tt.args.w, tt.args.r)
		})
	}
}
