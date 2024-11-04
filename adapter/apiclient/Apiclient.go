package apiclient

import (
	IApiclient "wellbe-common/domain/apiclient"
)

type api struct{
} 

func NewApiclient() IApiclient.Apiclient {
    return api{}
}