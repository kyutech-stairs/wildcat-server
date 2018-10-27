package services

type Service interface {
	ToOffsetTypeRequest() map[string]Service
}
