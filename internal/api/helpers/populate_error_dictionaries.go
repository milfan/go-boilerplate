package api_helpers

import (
	api_error "github.com/milfan/go-boilerplate/internal/api/errors"
	pkg_errors "github.com/milfan/go-boilerplate/pkg/errors"
	"github.com/sirupsen/logrus"
)

func PopulateErrorDicts() map[string]*pkg_errors.Error {
	errs := map[string]*pkg_errors.Error{}

	for item, err := range api_error.AppErrorDicts() {
		if errs[item] == nil {
			errs[item] = err
		} else {
			errs[item] = err
			logrus.Warnf("Error dicts with code %+v was overriden on HTTP error", item)
		}
	}

	for item, err := range api_error.DataErrorDicts() {
		if errs[item] == nil {
			errs[item] = err
		} else {
			errs[item] = err
			logrus.Warnf("Error dicts with code %+v was overriden on HTTP error", item)
		}
	}

	for item, err := range api_error.InfraErrorDicts() {
		if errs[item] == nil {
			errs[item] = err
		} else {
			errs[item] = err
			logrus.Warnf("Error dicts with code %+v was overriden on HTTP error", item)
		}
	}

	for item, err := range api_error.InterfaceErrorDicts() {
		if errs[item] == nil {
			errs[item] = err
		} else {
			errs[item] = err
			logrus.Warnf("Error dicts with code %+v was overriden on HTTP error", item)
		}
	}

	return errs
}
