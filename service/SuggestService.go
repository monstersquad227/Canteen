package service

import (
	"Canteen/model"
)

func SuggestList() model.Res {
	response := model.Res{
		Code:    0,
		Message: "Successful",
		Data:    nil,
	}
	s := model.Suggest{}
	result, err := s.List()
	if err != nil {
		response.Message = err.Error()
		return response
	}
	response.Data = result
	return response
}

func SuggestSave(category, location, details, user, contract, date string) model.Res {
	response := model.Res{
		Code:    0,
		Message: "Successful",
		Data:    nil,
	}
	s := model.Suggest{
		Category: category,
		Location: location,
		Details:  details,
		User:     user,
		Contract: contract,
		Date:     date,
	}
	result, err := s.Save()
	if err != nil {
		response.Message = err.Error()
		return response
	}
	response.Data = result
	return response
}
